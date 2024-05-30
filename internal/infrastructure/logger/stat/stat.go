package stat

import (
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/cast"
)

const (
	MAX_LOG_CHANNEL_BUFFER = 2000
)

type TimePredicate func(uint64) bool

type MetricEvent int

const (

	/************** 服务端 ***************/
	//服务请求总数
	MetricServerReqTotal = iota
	//慢响应数
	MetricServerSlowReplyTotal
	//状态码错误数
	MetricServerErrorTotal

	/************** 客户端 ***************/
	//请求总数
	MetricClientReqTotal
	//慢响应数
	MetricClientSlowReplyTotal
	//状态码错误数
	MetricClientErrorTotal

	/************** redis ***************/
	//请求总数
	MetricRedisReqTotal
	//慢响应数
	MetricRedisSlowReplyTotal
	//错误数
	MetriRedisErrorTotal

	/************** mysql ***************/
	//请求总数
	MetricMysqlReqTotal
	//慢响应数
	MetricMysqlSlowReplyTotal
	//错误数
	MetriMysqlErrorTotal

	MetricEventTotal
)

type Config struct {
	SlowMsThreshold       time.Duration
	ClientSlowMsThreshold time.Duration
	DbSlowMsThreshold     time.Duration
	RedisSlowMsThreshold  time.Duration
	WindowDuration        time.Duration
	WindowSize            uint32
}

type Stat struct {
	logChan               chan map[string]interface{}
	metrics               *Metrics
	bucketArray           *BucketLeapArray
	slowMsThreshold       time.Duration
	clientSlowMsThreshold time.Duration
	dbSlowMsThreshold     time.Duration
	redisSlowMsThreshold  time.Duration
	windowDuration        time.Duration
}

func NewStat(config *Config) *Stat {
	return &Stat{
		logChan:               make(chan map[string]interface{}, MAX_LOG_CHANNEL_BUFFER),
		metrics:               NewMetrics(),
		bucketArray:           NewBucketLeapArray(config.WindowSize, uint32(config.WindowDuration.Milliseconds())),
		slowMsThreshold:       config.SlowMsThreshold,
		clientSlowMsThreshold: config.ClientSlowMsThreshold,
		dbSlowMsThreshold:     config.DbSlowMsThreshold,
		redisSlowMsThreshold:  config.RedisSlowMsThreshold,
		windowDuration:        config.WindowDuration,
	}
}

func (s *Stat) Push(msgs map[string]interface{}) {
	if len(s.logChan) >= MAX_LOG_CHANNEL_BUFFER {
		return
	}
	s.logChan <- msgs
}

func (s *Stat) Consumer() *Stat {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Errorf("stat err: %+v", err)
			}
		}()

		for v := range s.logChan {
			s.statMain(v)
		}
	}()

	return s
}

func (s *Stat) statMain(logAttr map[string]interface{}) {
	//解析服务端
	if logAttr["x_module"] == "access/server" {
		s.statAccessServer(logAttr)
		return
	}

	//解析客户端
	if logAttr["x_module"] == "access/client" {
		s.statAccessClient(logAttr)
		return
	}

	//解析redis
	if logAttr["x_module"] == "pkg/redis" {
		s.statRedis(logAttr)
		return
	}

	//解析mysql
	if logAttr["x_module"] == "pkg/gorm" {
		s.statMysql(logAttr)
		return
	}
}

func (s *Stat) statAccessServer(logAttr map[string]interface{}) {
	action, kind := s.getActionAndComponent(logAttr)

	statData := make(map[MetricEvent]int64)
	statData[MetricServerReqTotal] = 1
	httpCode, ok := logAttr["x_code"].(int32)
	if ok && (httpCode != 0 && httpCode != 200) {
		// request error inc
		s.metrics.ServerRequestErrorCounter.With(kind, action).Inc()
		statData[MetricServerErrorTotal] = 1
	}

	duration, ok := logAttr["x_duration"].(float64)
	if ok {
		// request duration set
		s.metrics.UrlDurationMsHistogram.With(action).Observe(duration * 1000)
		if duration > s.slowMsThreshold.Seconds() {
			statData[MetricServerSlowReplyTotal] = 1
		}
	}
	s.bucketArray.BulkAddCount(statData)
}

func (s *Stat) statAccessClient(logAttr map[string]interface{}) {
	statData := make(map[MetricEvent]int64)
	statData[MetricClientReqTotal] = 1
	httpCode, ok := logAttr["x_code"].(int32)
	if ok && httpCode != 0 {
		statData[MetricClientErrorTotal] = 1
	}

	duration, ok := logAttr["x_duration"].(float64)
	if ok && duration > s.clientSlowMsThreshold.Seconds() {
		statData[MetricClientSlowReplyTotal] = 1
	}
	s.bucketArray.BulkAddCount(statData)
}

func (s *Stat) statRedis(logAttr map[string]interface{}) {
	statData := make(map[MetricEvent]int64)
	statData[MetricRedisReqTotal] = 1

	if logAttr["x_level"] == "log.error" {
		statData[MetriRedisErrorTotal] = 1
	}

	duration, ok := logAttr["x_duration"].(float64)
	if ok && duration > s.redisSlowMsThreshold.Seconds() {
		statData[MetricRedisSlowReplyTotal] = 1
	}
	s.bucketArray.BulkAddCount(statData)
}

func (s *Stat) statMysql(logAttr map[string]interface{}) {
	statData := make(map[MetricEvent]int64)
	statData[MetricMysqlReqTotal] = 1

	if logAttr["x_level"] == "log.error" {
		statData[MetriMysqlErrorTotal] = 1
	}

	duration, ok := logAttr["x_duration"].(float64)
	if ok {
		s.metrics.SqlDurationMsHistogram.Observe(duration * 1000)
		if duration > s.dbSlowMsThreshold.Seconds() {
			statData[MetricMysqlSlowReplyTotal] = 1
		}
	}
	s.bucketArray.BulkAddCount(statData)
}

func (s *Stat) getActionAndComponent(logAttr map[string]interface{}) (string, string) {
	action, ok := logAttr["x_action"]
	if !ok {
		action = "unknown"
	}
	component, ok := logAttr["x_component"]
	if !ok {
		component = "unknown"
	}
	return cast.ToString(action), cast.ToString(component)
}

func (s *Stat) GetNoticeSummary() string {
	summary := s.bucketArray.BulkCount(
		MetricServerReqTotal,
		MetricServerSlowReplyTotal,
		MetricServerErrorTotal,

		MetricClientReqTotal,
		MetricClientSlowReplyTotal,
		MetricClientErrorTotal,

		MetricRedisReqTotal,
		MetricRedisSlowReplyTotal,
		MetriRedisErrorTotal,

		MetricMysqlReqTotal,
		MetricMysqlSlowReplyTotal,
		MetriMysqlErrorTotal,
	)

	//没有错误返回空
	if summary[MetricServerErrorTotal] <= 2 &&
		summary[MetricClientErrorTotal] <= 0 &&
		summary[MetriRedisErrorTotal] <= 0 &&
		summary[MetriMysqlErrorTotal] <= 0 {
		return ""
	}

	//有错误出现返回报警
	msg := "最近 %d 秒运行情况: \n"
	msg += "-----自身服务-----\n"
	msg += "请求总量: %d \n"
	msg += "慢请求 >%d ms: %d \n"
	msg += "状态码异常数量: %d \n\n"
	msg += "-----对外客户端-----\n"
	msg += "请求总量: %d \n"
	msg += "慢请求 >%d ms: %d \n"
	msg += "异常数量: %d \n\n"
	msg += "-----redis请求-----\n"
	msg += "请求总量: %d \n"
	msg += "慢请求 >%d ms: %d \n"
	msg += "异常数量: %d \n\n"
	msg += "-----mysql请求-----\n"
	msg += "请求总量: %d \n"
	msg += "慢请求 >%d ms: %d \n"
	msg += "异常数量: %d \n"

	msg = fmt.Sprintf(msg,
		int(s.windowDuration.Seconds()),

		summary[MetricServerReqTotal],
		s.slowMsThreshold.Milliseconds(),
		summary[MetricServerSlowReplyTotal],
		summary[MetricServerErrorTotal],

		summary[MetricClientReqTotal],
		s.clientSlowMsThreshold.Milliseconds(),
		summary[MetricClientSlowReplyTotal],
		summary[MetricClientErrorTotal],

		summary[MetricRedisReqTotal],
		s.redisSlowMsThreshold.Milliseconds(),
		summary[MetricRedisSlowReplyTotal],
		summary[MetriRedisErrorTotal],

		summary[MetricMysqlReqTotal],
		s.dbSlowMsThreshold.Milliseconds(),
		summary[MetricMysqlSlowReplyTotal],
		summary[MetriMysqlErrorTotal],
	)

	return msg
}
