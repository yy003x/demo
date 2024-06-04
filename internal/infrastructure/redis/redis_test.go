package redis

import (
	"be_demo/internal/infrastructure/logger"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

func newRedisConn() (*LoggingConn, error) {
	l := log.With(
		logger.NewLogger(logger.Config{
			Path:  "./log/api-%Y%m%d.log",
			Level: "info",
		}),
		"x_source", logger.Caller(3),
		"x_trace_id", logger.TraceID("x-traceid"),
		"x_rpc_id", logger.IncrRpcId("x-rpcid"),
	)
	redisConn, err := NewLoggingPool(&Config{
		Addr:         "127.0.0.1:6379",
		Auth:         "123456",
		SelectDb:     0,
		MaxIdleConns: 1,
	}, l)
	return redisConn, err
}

func TestSetNxWithExpire(t *testing.T) {
	r, err := newRedisConn()
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	r.SetNXWithExpire(ctx, "abcd", "1", 10*time.Second)
}

func TestExpireAt(t *testing.T) {
	r, err := newRedisConn()
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	err = r.Set(ctx, "test-expireat", 1)
	if err != nil {
		t.Error(err)
	}
	r.ExpireAt(ctx, "test-expireat", 1666699553)
}

func TestIncrbyEX(t *testing.T) {
	r, err := newRedisConn()
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	_ = r.IncrByWithDuration(ctx, "te", 1, time.Second*10)
}

func TestGet(t *testing.T) {
	r, err := newRedisConn()
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	a, err := r.GetInt(ctx, "te")
	fmt.Println(a, err)
}

func TestExists(t *testing.T) {
	r, err := newRedisConn()
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	a, err := r.Exists(ctx, "te")
	fmt.Println(a, err)
}

func TestPipelineZAddWithDuration(t *testing.T) {
	r, err := newRedisConn()
	if err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	setData := make(map[string]map[interface{}]int)

	memberScores := make(map[interface{}]int)

	memberScores["relation.SkuUnionid1"] = 1
	memberScores["relation.SkuUnionid2"] = 2

	setData["te1"] = memberScores

	memberScores2 := make(map[interface{}]int)

	memberScores2["relation.SkuUnionid3"] = 3
	memberScores2["relation.SkuUnionid4"] = 5
	setData["te2"] = memberScores2
	var du time.Duration
	err = r.PipelineZAddWithDuration(ctx, []string{"te1", "te2"}, setData, du)
	fmt.Println(err)
}
