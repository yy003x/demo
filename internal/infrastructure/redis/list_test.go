package redis

import (
	"context"
	"fmt"
	"testing"
	"time"
)

type redisTester struct {
	ctx   context.Context
	redis *LoggingConn
}

func newRedisTester() *redisTester {
	c, err := newRedisConn()
	if err != nil {
		fmt.Println("newRedisConn failed with err:", err)
		return nil
	}
	return &redisTester{
		context.Background(),
		c,
	}
}

var (
	tester = newRedisTester()
)

func TestLoggingConn_BLPop(t *testing.T) {
	value, err := tester.redis.BLPop(tester.ctx, 3*time.Second, "TestLoggingConn_BLPop_missing")
	t.Log("BRPop return ", value, err)
	if err == nil {
		t.Error("BLPop should timeout err")
	}

	k := "TestLoggingConn_BLPop"
	v := "12345"
	len, err := tester.redis.LPush(tester.ctx, k, v)
	t.Log("LPUSH return ", len, err)
	if err != nil {
		t.Error(err)
	}
	value, err = tester.redis.BLPop(tester.ctx, 3*time.Second, k)
	t.Log("BLPop return ", value, err)
	if err != nil {
		t.Error(err)
	}
	if value != v {
		t.Errorf("BRPop return %s != %s", value, v)
	}
}

func TestLoggingConn_BRPop(t *testing.T) {
	value, err := tester.redis.BRPop(tester.ctx, 3*time.Second, "TestLoggingConn_BRPop_missing")
	t.Log("BRPop return ", value, err)
	if err == nil {
		t.Error("BLPop should timeout err")
	}

	k := "TestLoggingConn_BRPop"
	v := "12345"
	len, err := tester.redis.LPush(tester.ctx, k, v)
	t.Log("LPUSH return ", len, err)
	if err != nil {
		t.Error(err)
	}
	value, err = tester.redis.BRPop(tester.ctx, 3*time.Second, k)
	t.Log("BRPop return ", value, err)
	if err != nil {
		t.Error(err)
	}
	if value != v {
		t.Errorf("BRPop return %s != %s", value, v)
	}
}

func TestLoggingConn_BRPopLPush(t *testing.T) {
	src := "TestLoggingConn_BRPopLPush_Src"
	srcV := "112233"
	l, err := tester.redis.LPush(tester.ctx, src, srcV, srcV)
	t.Log("LPUSH return ", l, err)
	if err != nil {
		t.Error(err)
	}

	vs, err := tester.redis.LRangeAll(tester.ctx, src)
	t.Log("LRangeAll return ", vs, err)
	if err != nil {
		t.Error(err)
	}

	dest := "TestLoggingConn_BRPopLPush_Desc"
	destV := "334455"
	len2, err := tester.redis.LPush(tester.ctx, dest, destV)
	t.Log("LPUSH return ", len2, err)
	if err != nil {
		t.Error(err)
	}

	value, err := tester.redis.BRPopLPush(tester.ctx, src, dest, 3*time.Second)
	t.Log("BRPopLPush return ", value, err)
	if err != nil {
		t.Error(err)
		return
	}
	if value != srcV {
		t.Errorf("BRPopLPush return %s != %s", value, srcV)
	}

	value2, err := tester.redis.LRangeAll(tester.ctx, dest)
	t.Log("LRangeAll return ", value2, err)
	if err != nil {
		t.Error(err)
	}
	if len(value2) != 2 {
		t.Errorf("BRPop return %s != %s", value2, srcV)
	}

	err = tester.redis.MDel(tester.ctx, []string{src, dest})
	if err != nil {
		t.Error(err)
	}
}

func TestLoggingConn_LIndex(t *testing.T) {
	k := "TestLoggingConn_LIndex"
	values := []any{
		"12345",
		"56789",
	}
	len, err := tester.redis.RPush(tester.ctx, k, values...)
	t.Log("LPUSH return ", len, err)
	if err != nil {
		t.Error(err)
	}

	for index, v := range values {
		value, err := tester.redis.LIndex(tester.ctx, k, index)
		t.Logf("LIndex(%s, %d) return %s, %v", k, index, value, err)
		if err != nil {
			t.Error(err)
		}
		if value != v {
			t.Errorf("LIndex return %s != %s", value, v)
		}
	}

	err = tester.redis.Del(tester.ctx, k)
	if err != nil {
		t.Error(err)
	}
}

func TestLoggingConn_LInsertAfter(t *testing.T) {
	k := "TestLoggingConn_LInsertAfter"
	values := []any{
		"12345",
		"56789",
		"112233",
	}
	insertValue := "5324412"
	len, err := tester.redis.RPush(tester.ctx, k, values...)
	t.Log("LPUSH return ", len, err)
	if err != nil {
		t.Error(err)
	}

	len, err = tester.redis.LInsertAfter(tester.ctx, k, values[1], insertValue)
	t.Logf("LInsertAfter(%v, %v, %v) return %v, %v", k, values[1], insertValue, len, err)
	if err != nil {
		t.Error(err)
	}

	value, err := tester.redis.LIndex(tester.ctx, k, 2)
	t.Logf("LIndex(%s, %d) return %s, %v", k, 1, value, err)
	if err != nil {
		t.Error(err)
	}
	if value != insertValue {
		t.Errorf("LIndex return %s != %s", value, insertValue)
	}

	err = tester.redis.Del(tester.ctx, k)
	if err != nil {
		t.Error(err)
	}
}

func TestLoggingConn_LInsertBefore(t *testing.T) {
	k := "TestLoggingConn_LInsertBefore"
	values := []any{
		"12345",
		"56789",
		"112233",
	}
	insertValue := "5324412"
	len, err := tester.redis.RPush(tester.ctx, k, values...)
	t.Log("LPUSH return ", len, err)
	if err != nil {
		t.Error(err)
	}

	len, err = tester.redis.LInsertBefore(tester.ctx, k, values[1], insertValue)
	t.Logf("LInsertBefore(%v, %v, %v) return %v, %v", k, values[1], insertValue, len, err)
	if err != nil {
		t.Error(err)
	}

	value, err := tester.redis.LIndex(tester.ctx, k, 1)
	t.Logf("LIndex(%s, %d) return %s, %v", k, 1, value, err)
	if err != nil {
		t.Error(err)
	}
	if value != insertValue {
		t.Errorf("LIndex return %s != %s", value, insertValue)
	}

	err = tester.redis.Del(tester.ctx, k)
	if err != nil {
		t.Error(err)
	}
}

func TestLoggingConn_LLen(t *testing.T) {
	k := "TestLoggingConn_LLen"
	values := []any{
		"12345",
		"56789",
	}

	len, err := tester.redis.LPush(tester.ctx, k, values...)
	t.Log("LPUSH return ", len, err)
	if err != nil {
		t.Error(err)
	}

	getLen, err := tester.redis.LLen(tester.ctx, k)
	t.Log("LLen return ", len, err)
	if err != nil {
		t.Error(err)
	}

	if getLen != len {
		t.Errorf("LLen return %d != len:%d", getLen, len)
	}
}

func TestLoggingConn_LPop(t *testing.T) {
	k := "TestLoggingConn_LPop"
	v := "12345"
	len, err := tester.redis.LPush(tester.ctx, k, v)
	t.Log("LPUSH return ", len, err)
	if err != nil {
		t.Error(err)
	}
	value, err := tester.redis.LPop(tester.ctx, k)
	t.Log("LPop return ", value, err)
	if err != nil {
		t.Error(err)
	}
	if value != v {
		t.Errorf("LPop return %s != %s", value, v)
	}
}

func TestLoggingConn_LPush(t *testing.T) {
	k := "TestLoggingConn_LPop"
	v := "12345"
	len, err := tester.redis.LPush(tester.ctx, k, v)
	t.Log("LPUSH return ", len, err)
	if err != nil {
		t.Error(err)
	}
}

func TestLoggingConn_LPushX(t *testing.T) {
	k := "TestLoggingConn_LPop"
	v := "12345"
	len, err := tester.redis.LPushX(tester.ctx, k, v)
	t.Log("LPushX return ", len, err)
	if err != nil {
		t.Error(err)
	}

	l1, err := tester.redis.LPush(tester.ctx, k, v)
	t.Log("LPushX return ", l1, err)
	if err != nil {
		t.Error(err)
	}

	l2, err := tester.redis.LPushX(tester.ctx, k, v)
	t.Log("LPushX return ", l2, err)
	if err != nil {
		t.Error(err)
	}
}

func TestLoggingConn_LRange(t *testing.T) {
	k := "TestLoggingConn_LLen"
	values := []any{
		"12345",
		"56789",
	}

	length, err := tester.redis.RPush(tester.ctx, k, values...)
	t.Log("LPUSH return ", length, err)
	if err != nil {
		t.Error(err)
	}

	getValues, err := tester.redis.LRange(tester.ctx, k, 0, -1)
	t.Log("LRange return ", getValues, err)
	if err != nil {
		t.Error(err)
	}

	if len(getValues) != length {
		t.Errorf("LRange return %d != len:%d", len(getValues), length)
	}
}

func TestLoggingConn_LRangeAll(t *testing.T) {
	k := "TestLoggingConn_LRangeAll"
	values := []any{
		"12345",
		"56789",
	}

	length, err := tester.redis.RPush(tester.ctx, k, values...)
	t.Log("LPUSH return ", length, err)
	if err != nil {
		t.Error(err)
	}

	getValues, err := tester.redis.LRangeAll(tester.ctx, k)
	t.Log("LRange return ", getValues, err)
	if err != nil {
		t.Error(err)
	}

	if len(getValues) != length {
		t.Errorf("LRange return %d != len:%d", len(getValues), length)
	}

	nv, err := tester.redis.LRangeAll(tester.ctx, "not_exists_set")
	t.Log("LRangeAll not_exists_set return ", nv, err)
	if err != nil {
		t.Error(err)
	}
}

func TestLoggingConn_LRem(t *testing.T) {
	k := "TestLoggingConn_LRem"
	values := []any{
		"12345",
		"12345",
		"12345",
		"56789",
	}

	length, err := tester.redis.RPush(tester.ctx, k, values...)
	t.Log("LPUSH return ", length, err)
	if err != nil {
		t.Error(err)
	}

	remLen, err := tester.redis.LRem(tester.ctx, k, 3, "12345")
	t.Log("LRem return ", remLen, err)
	if err != nil {
		t.Error(err)
	}

	if remLen != 3 {
		t.Error("LRem should be ", remLen)
	}
}

func TestLoggingConn_LSet(t *testing.T) {
	k := "TestLoggingConn_LRem"
	values := []any{
		"12345",
		"12345",
		"12345",
		"56789",
	}

	length, err := tester.redis.LPush(tester.ctx, k, values...)
	t.Log("LPUSH return ", length, err)
	if err != nil {
		t.Error(err)
	}

	err = tester.redis.LSet(tester.ctx, k, 1, "X")
	t.Log("LSet return ", err)
	if err != nil {
		t.Error(err)
	}

	vs, err := tester.redis.LRangeAll(tester.ctx, k)
	t.Log("LRangeAll return ", vs, err)
	if err != nil {
		t.Error(err)
	}
	if len(vs) != len(values)+1 {
		t.Error("TestLoggingConn_LSet,LRangeAll should return 5 element")
	}

	err = tester.redis.Del(tester.ctx, k)
	if err != nil {
		t.Error(err)
	}
}

func TestLoggingConn_LTrim(t *testing.T) {
	k := "TestLoggingConn_LTrim"
	values := []any{
		"1",
		"1",
		"3",
		"4",
	}

	length, err := tester.redis.RPush(tester.ctx, k, values...)
	t.Log("LPUSH return ", length, err)
	if err != nil {
		t.Error(err)
	}
	vs, err := tester.redis.LRangeAll(tester.ctx, k)
	t.Log("LRangeAll return ", vs, err)
	if err != nil {
		t.Error(err)
	}

	err = tester.redis.LTrim(tester.ctx, k, 1, 2)
	t.Log("LTrim return ", err)
	if err != nil {
		t.Error(err)
	}

	vs, err = tester.redis.LRangeAll(tester.ctx, k)
	t.Log("LRangeAll return ", vs, err)
	if err != nil {
		t.Error(err)
	}
	if len(vs) != 2 {
		t.Error("TestLoggingConn_LSet,LRangeAll should return 2 element")
	}

	err = tester.redis.Del(tester.ctx, k)
	if err != nil {
		t.Error(err)
	}
}

func TestLoggingConn_RPop(t *testing.T) {
	k := "TestLoggingConn_RPop"
	values := []any{
		"12345",
		"12345",
		"12345",
		"56789",
	}

	length, err := tester.redis.RPush(tester.ctx, k, values...)
	t.Log("LPUSH return ", length, err)
	if err != nil {
		t.Error(err)
	}

	pop, err := tester.redis.RPop(tester.ctx, k)
	t.Log("RPop return ", pop, err)
	if err != nil {
		t.Error(err)
	}

	if pop != values[len(values)-1] {
		t.Error("RPop should return ", values[len(values)-1])
	}
}

func TestLoggingConn_RPopLPush(t *testing.T) {
	k1 := "TestLoggingConn_RPopLPush_1"
	k2 := "TestLoggingConn_RPopLPush_2"
	values := []any{
		"12345",
		"12345",
		"12345",
		"56789",
	}

	length, err := tester.redis.RPush(tester.ctx, k1, values...)
	t.Log("LPUSH return ", length, err)
	if err != nil {
		t.Error(err)
	}

	pop1, err := tester.redis.RPopLPush(tester.ctx, k1, k2)
	t.Log("RPopLPush return ", pop1, err)
	if err != nil {
		t.Error(err)
	}

	if pop1 != values[len(values)-1] {
		t.Error("RPop should return ", values[len(values)-1])
	}

	pop2, err := tester.redis.RPop(tester.ctx, k2)
	t.Log("RPop return ", pop2, err)
	if err != nil {
		t.Error(err)
	}

	if pop2 != pop1 || pop2 != values[len(values)-1] {
		t.Error("RPop should return ", values[len(values)-1])
	}
}

func TestLoggingConn_RPush(t *testing.T) {
	k1 := "TestLoggingConn_RPush"
	values := []any{
		"12345",
		"12345",
		"12345",
		"56789",
	}

	length, err := tester.redis.RPush(tester.ctx, k1, values...)
	t.Log("LPUSH return ", length, err)
	if err != nil {
		t.Error(err)
	}
}

func TestLoggingConn_RPushX(t *testing.T) {
	k := "TestLoggingConn_RPushX"
	v := "12345"
	if tester == nil {
		t.Error("tester is nil")
		return
	}
	len, err := tester.redis.RPushX(tester.ctx, k, v)
	t.Log("RPushX return ", len, err)
	if err != nil {
		t.Error(err)
	}

	l1, err := tester.redis.RPush(tester.ctx, k, v)
	t.Log("RPush return ", l1, err)
	if err != nil {
		t.Error(err)
	}

	l2, err := tester.redis.RPushX(tester.ctx, k, v)
	t.Log("RPushX return ", l2, err)
	if err != nil {
		t.Error(err)
	}
}
