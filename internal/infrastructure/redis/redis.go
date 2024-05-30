package redis

import (
	"context"
	"encoding/json"
	"errors"
	"strconv"
	"strings"
	"time"

	"git.100tal.com/kratos-lib/common"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/cast"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

const (
	MaxReplay       = 200
	MaxTimeDuration = time.Second
)

type LoggingConn struct {
	Pool        *redis.Pool
	logger      *log.Helper
	recordTrace bool
}

type LoggingConnOption func(o *LoggingConn)

type logAttr struct {
	method    string
	command   string
	script    string
	args      []interface{}
	reply     interface{}
	startTime time.Time

	err    error
	shrink time.Duration
}

type Config struct {
	Addr               string
	Auth               string
	SelectDb           int
	MaxIdleConns       int
	IdleTimeout        time.Duration
	DialConnectTimeout time.Duration
	DialReadTimeout    time.Duration
	DialWriteTimeout   time.Duration
}

// NewLoggingPool returns a logging wrapper around a connection.
func NewLoggingPool(config *Config, logger log.Logger, ops ...LoggingConnOption) (*LoggingConn, error) {
	redisPool := &redis.Pool{
		MaxIdle:     config.MaxIdleConns,
		IdleTimeout: config.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.Addr,
				redis.DialConnectTimeout(config.DialConnectTimeout),
				redis.DialReadTimeout(config.DialReadTimeout),
				redis.DialWriteTimeout(config.DialWriteTimeout),
				redis.DialPassword(config.Auth),
				redis.DialDatabase(config.SelectDb),
			)
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if t.Add(config.IdleTimeout).After(time.Now()) {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}

	if _, err := redisPool.Get().Do("PING"); err != nil {
		return nil, err
	}

	lc := &LoggingConn{
		Pool:        redisPool,
		logger:      log.NewHelper(log.With(logger, "x_module", "pkg/redis")),
		recordTrace: false,
	}

	for _, o := range ops {
		o(lc)
	}

	return lc, nil
}

func WithLogger(logger *log.Helper) LoggingConnOption {
	return func(o *LoggingConn) {
		o.logger = logger
	}
}

func WithTrace() LoggingConnOption {
	return func(o *LoggingConn) {
		o.recordTrace = true
	}
}

func stringifyReply(reply any) string {
	str, _ := json.Marshal(reply)
	return string(str)
}

func (c *LoggingConn) log(ctx context.Context, attr logAttr) {
	var (
		x_action   string
		x_params   string
		x_shrink   float64
		x_response string
		x_error    string
	)

	x_action = strings.Join([]string{attr.method, attr.command}, ".")
	x_shrink = attr.shrink.Seconds()

	replyStr := cast.ToString(attr.reply)
	if len(replyStr) > MaxReplay {
		x_response = replyStr[:MaxReplay] + "..."
	} else {
		x_response = replyStr
	}
	if attr.err != nil {
		x_error = attr.err.Error()
	}
	if len(attr.args) > 0 {
		strList := append([]string{attr.command}, cast.ToStringSlice(attr.args)...)
		x_params = strings.Join(strList, " ")
	}

	c.logger.WithContext(ctx).Infow(
		"x_action", x_action,
		"x_param", x_params,
		"x_response", x_response,
		"x_shrink", x_shrink,
		"x_error", x_error,
		"x_duration", time.Since(attr.startTime).Seconds(),
	)

	if c.recordTrace {
		tracer := otel.GetTracerProvider().Tracer("redis")
		ops := []trace.SpanStartOption{
			trace.WithSpanKind(trace.SpanKindClient),
			trace.WithAttributes(
				semconv.DBSystemRedis,
				semconv.DBOperationKey.String(x_action),
				semconv.DBStatementKey.String(x_params),
				attribute.String("db.redis.response", x_response),
			),
			trace.WithTimestamp(attr.startTime),
		}
		_, span := tracer.Start(ctx, "redis", ops...)
		if attr.err != nil {
			span.RecordError(attr.err)
		}

		defer span.End()
	}
}

func (c *LoggingConn) do(ctx context.Context, commandName string, args ...interface{}) (interface{}, error) {
	conn, err := c.Pool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	startTime := time.Now()
	reply, err := conn.Do(commandName, args...)
	c.log(ctx, logAttr{
		method:    "do",
		startTime: startTime,
		command:   commandName,
		args:      args,
		reply:     reply,
		err:       err,
	})
	return reply, err
}

func (c *LoggingConn) doContext(ctx context.Context, commandName string, args ...interface{}) (interface{}, error) {
	conn, err := c.Pool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	startTime := time.Now()
	reply, err := redis.DoContext(conn, ctx, commandName, args...)
	c.log(ctx, logAttr{
		method:    "doContext",
		startTime: startTime,
		command:   commandName,
		args:      args,
		reply:     reply,
		err:       err,
	})
	return reply, err
}

func (c *LoggingConn) DoWithTimeout(ctx context.Context, commandName string, args ...interface{}) (interface{}, error) {
	conn, err := c.Pool.GetContext(ctx)
	if err != nil {
		if err == redis.ErrPoolExhausted {
			c.logger.Errorf("DoWithTimeout.%s failed, %s", commandName, err.Error())
		}
		return nil, err
	}
	defer conn.Close()

	startTime := time.Now()
	timeout := common.ShrinkDuration(ctx, MaxTimeDuration)
	reply, err := redis.DoWithTimeout(conn, timeout, commandName, args...)
	c.log(ctx, logAttr{
		method:    "DoWithTimeout",
		startTime: startTime,
		command:   commandName,
		args:      args,
		reply:     reply,
		err:       err,
		shrink:    timeout,
	})
	return reply, err
}

func (c *LoggingConn) EvalWithContext(ctx context.Context, script string, keyCount int, keyAndArgs ...interface{}) (interface{}, error) {
	conn, err := c.Pool.GetContext(ctx)
	if err != nil {
		if err == redis.ErrPoolExhausted {
			c.logger.Errorf("EvalWithTimeout script:%s failed, %s", script, err.Error())
		}
		return nil, err
	}
	defer conn.Close()

	startTime := time.Now()
	timeout := common.ShrinkDuration(ctx, MaxTimeDuration)

	lua := redis.NewScript(keyCount, script)
	reply, err := lua.DoContext(ctx, conn, keyAndArgs...)
	c.log(ctx, logAttr{
		method:    "EvalWithContext",
		startTime: startTime,
		script:    script,
		args:      keyAndArgs,

		reply:  reply,
		err:    err,
		shrink: timeout,
	})
	return reply, err
}

func (c *LoggingConn) IsErrNil(err error) bool {
	return err == redis.ErrNil
}

func (c *LoggingConn) TTl(ctx context.Context, key string) int {
	times, _ := redis.Int(c.DoWithTimeout(ctx, "TTL", key))
	return times
}

func (c *LoggingConn) Expire(ctx context.Context, key string, time time.Duration) error {
	seconds := int(time.Seconds())
	var err error
	if seconds > 0 {
		_, err = c.DoWithTimeout(ctx, "EXPIRE", key, int(time.Seconds()))
	}
	return err
}

func (c *LoggingConn) ExpireAt(ctx context.Context, key string, expireAt int64) error {
	var err error
	if expireAt <= time.Now().Unix() || len(strconv.Itoa(int(expireAt))) != 10 {
		expireAt = time.Now().Unix()
	}

	_, err = redis.Int64(c.DoWithTimeout(ctx, "EXPIREAT", key, int(expireAt)))
	if err != nil {
		return errors.New("set redis timestamp failed, " + err.Error())
	}
	return nil
}

func (c *LoggingConn) PipelineExpire(ctx context.Context, keys []string, duration time.Duration) error {
	seconds := int(duration.Seconds())
	if seconds <= 0 {
		return nil
	}
	startTime := time.Now()
	conn, err := c.Pool.GetContext(ctx)
	if err != nil {
		return err
	}

	var args []any
	conn.Send("MULTI")
	for _, key := range keys {
		if len(key) == 0 {
			continue
		}
		//args = append(args, key)
		conn.Send("EXPIRE", key, seconds)
	}
	_, err = conn.Do("EXEC")
	conn.Close()

	c.log(ctx, logAttr{
		method:    "PipelineExpire",
		startTime: startTime,
		command:   "MULTI EXPIRE",
		args:      args,
		reply:     "",
		err:       err,
	})
	return err
}

func (c *LoggingConn) Del(ctx context.Context, key string) error {
	_, err := c.DoWithTimeout(ctx, "DEL", key)
	return err
}

func (c *LoggingConn) MDel(ctx context.Context, keys []string) (err error) {
	startTime := time.Now()
	conn, err := c.Pool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	conn.Send("MULTI")
	var args []any
	for _, key := range keys {
		if len(key) == 0 {
			continue
		}
		args = append(args, key)
		conn.Send("DEL", strings.Trim(key, " "))
	}
	_, err = conn.Do("EXEC")

	c.log(ctx, logAttr{
		method:    "MDel",
		startTime: startTime,
		command:   "MULTI DEL",
		args:      args,
		reply:     "",
		err:       err,
	})
	return
}

//Exists 判断key是否存在
func (c *LoggingConn) Exists(ctx context.Context, key string) (bool, error) {
	res, err := redis.Bool(c.DoWithTimeout(ctx, "EXISTS", key))
	return res, err
}
