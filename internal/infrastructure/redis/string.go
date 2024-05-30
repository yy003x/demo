package redis

import (
	"context"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"
)

func (c *LoggingConn) GetString(ctx context.Context, key string) (string, error) {
	res, err := redis.String(c.DoWithTimeout(ctx, "GET", key))
	return res, err
}

func (c *LoggingConn) GetInt(ctx context.Context, key string) (int, error) {
	res, err := redis.Int(c.DoWithTimeout(ctx, "GET", key))
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (c *LoggingConn) MGetInt(ctx context.Context, keys ...string) ([]int, error) {
	newKeys := make([]interface{}, len(keys))
	for k := range keys {
		newKeys[k] = keys[k]
	}
	return redis.Ints(c.DoWithTimeout(ctx, "MGET", newKeys...))
}

func (c *LoggingConn) MGetString(ctx context.Context, keys ...string) ([]string, error) {
	newKeys := make([]interface{}, len(keys))
	for k := range keys {
		newKeys[k] = keys[k]
	}
	return redis.Strings(c.DoWithTimeout(ctx, "MGET", newKeys...))
}

func (c *LoggingConn) Set(ctx context.Context, key string, val interface{}) error {
	_, err := c.DoWithTimeout(ctx, "SET", key, val)
	return err
}

func (c *LoggingConn) SetEX(ctx context.Context, key string, val interface{}, duration time.Duration) error {
	var err error
	seconds := int(duration.Seconds())
	if seconds > 0 {
		_, err = c.DoWithTimeout(ctx, "SET", key, val, "EX", seconds)
	} else {
		_, err = c.DoWithTimeout(ctx, "SET", key, val)
	}
	return err
}

func (c *LoggingConn) MSet(ctx context.Context, args ...interface{}) error {
	_, err := c.DoWithTimeout(ctx, "MSET", args...)
	return err
}

// MSetNx 同时设置一个或多个 key-value 对，当且仅当所有给定 key 都不存在。
// 当所有 key 都成功设置，返回 true 。 如果所有给定 key 都设置失败(至少有一个 key 已经存在)，那么返回 false 。
func (c *LoggingConn) MSetNx(ctx context.Context, args ...interface{}) (bool, error) {
	return redis.Bool(c.DoWithTimeout(ctx, "MSETNX", args...))
}

func (c *LoggingConn) MSetEx(ctx context.Context, duration time.Duration, values map[string]any) error {
	seconds := int(duration.Seconds())
	var args []any
	var keys []string
	for k, v := range values {
		keys = append(keys, k)
		args = append(args, k, v)
	}

	err := c.MSet(ctx, args...)
	if err == nil {
		if seconds > 0 {
			err = c.PipelineExpire(ctx, keys, duration)
		}
	}
	return err
}

func (c *LoggingConn) SetNXWithExpire(ctx context.Context, key, value string, duration time.Duration) (bool, error) {
	seconds := int(duration.Seconds())
	var err error
	var result interface{}
	if seconds > 0 {
		result, err = c.DoWithTimeout(ctx, "SET", key, value, "EX", seconds, "NX")
	} else {
		result, err = c.DoWithTimeout(ctx, "SET", key, value, "NX")
	}
	str, err := redis.String(result, err)
	if err == redis.ErrNil {
		return false, nil
	}
	return strings.ToUpper(str) == "OK", err
}

func (c *LoggingConn) IncrBy(ctx context.Context, key string, val int) error {
	_, err := c.DoWithTimeout(ctx, "INCRBY", key, val)
	return err
}

func (c *LoggingConn) Incr(ctx context.Context, key string) (int, error) {
	return redis.Int(c.DoWithTimeout(ctx, "INCR", key))
}

func (c *LoggingConn) IncrByWithDuration(ctx context.Context, key string, val int, duration time.Duration) error {
	var err error
	seconds := int(duration.Seconds())
	_, err = c.DoWithTimeout(ctx, "INCRBY", key, val)
	if seconds > 0 {
		_, err = c.DoWithTimeout(ctx, "EXPIRE", key, seconds)
	}
	return err
}

func (c *LoggingConn) DecrBy(ctx context.Context, key string, val int) error {
	_, err := c.DoWithTimeout(ctx, "DECRBY", key, val)
	return err
}

func (c *LoggingConn) DecrByWithDuration(ctx context.Context, key string, val int, duration time.Duration) error {
	var err error
	seconds := int(duration.Seconds())
	_, err = c.DoWithTimeout(ctx, "DECRBY", key, val)
	if seconds > 0 {
		_, err = c.DoWithTimeout(ctx, "EXPIRE", key, seconds)
	}
	return err
}

// GetSet 将给定 key 的值设为 value ，并返回 key 的旧值(old value)。
func (c *LoggingConn) GetSet(ctx context.Context, key string, val interface{}) (string, error) {
	return redis.String(c.DoWithTimeout(ctx, "GETSET", key, val))
}

// GetBit 命令用于对 key 所储存的字符串值，获取指定偏移量上的位(bit)。
// 返回字符串值指定偏移量上的位(bit 0/1)。 当偏移量 OFFSET 比字符串值的长度大，或者 key 不存在时，返回 0 。
func (c *LoggingConn) GetBit(ctx context.Context, key string, offset int) (int, error) {
	return redis.Int(c.DoWithTimeout(ctx, "GETBIT", key, offset))
}

// SetBit 命令用于对 key 所储存的字符串值，设置或清除指定偏移量上的位(bit)。
// 返回指定偏移量原来储存的位。
func (c *LoggingConn) SetBit(ctx context.Context, key string, offset int, bit int) (int, error) {
	return redis.Int(c.DoWithTimeout(ctx, "SETBIT", key, offset, bit))
}

func (c *LoggingConn) BitCount(ctx context.Context, key string) (int, error) {
	return redis.Int(c.DoWithTimeout(ctx, "BITCOUNT", key))
}

// EXISTS 命令用于检查给定 key 是否存在。
func (c *LoggingConn) EXISTS(ctx context.Context, key string) (int, error) {
	reply, err := redis.Int(c.DoWithTimeout(ctx, "EXISTS", key))
	return reply, err
}

// Append 如果 key 已经存在并且是一个字符串，APPEND 命令将指定的 value 追加到该 key 原来值（value）的末尾。
// 返回追加指定值之后，key 中字符串的长度。
func (c *LoggingConn) Append(ctx context.Context, key string, value interface{}) (int, error) {
	return redis.Int(c.DoWithTimeout(ctx, "APPEND", key, value))
}

func (c *LoggingConn) PipelineBitCount(ctx context.Context, keyList []string) (map[string]int, error) {
	conn, err := c.Pool.GetContext(ctx)
	if err != nil {
		if err == redis.ErrPoolExhausted {
			c.logger.Errorf("PipelineBitCount failed, %s", err.Error())
		}
		return nil, err
	}
	defer conn.Close()

	startTime := time.Now()
	for _, key := range keyList {
		_ = conn.Send("BITCOUNT", key)
	}
	if err := conn.Flush(); err != nil {
		return nil, err
	}

	var args []any
	containers := make(map[string]int, len(keyList))
	for _, key := range keyList {
		args = append(args, key)
		reply, err := redis.Int(redis.ReceiveContext(conn, ctx))
		if err != nil {
			c.logger.WithContext(ctx).Errorf("Redis PipelineBitCount error %v", err)
			continue
		}
		containers[key] = reply
	}

	c.log(ctx, logAttr{
		method:    "PipelineBitCount",
		startTime: startTime,
		command:   "BITCOUNT",
		args:      args,
		reply:     stringifyReply(containers),
		err:       err,
	})
	return containers, nil
}
