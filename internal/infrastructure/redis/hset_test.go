package redis

import (
	"math"
	"testing"
	"time"
)

func TestLoggingConn_HDel(t *testing.T) {
	//
	k := "TestLoggingConn_HDel"
	f := "field"
	v := "12345"

	err := tester.redis.HSet(tester.ctx, k, f, v)
	t.Log("HSet return ", err)
	if err != nil {
		t.Error(err)
	}

	err = tester.redis.HDel(tester.ctx, k, f)
	t.Log("HDel return ", err)
	if err != nil {
		t.Error(err)
	}

	tester.redis.Del(tester.ctx, k)
}

func TestLoggingConn_HExists(t *testing.T) {
	k := "TestLoggingConn_HExists"
	f := "field"
	v := "12345"

	err := tester.redis.HSet(tester.ctx, k, f, v)
	t.Log("HSet return ", err)
	if err != nil {
		t.Error(err)
	}

	exists, err := tester.redis.HExists(tester.ctx, k, f)
	t.Log("HExists return ", err)
	if err != nil {
		t.Error(err)
	}
	if exists != true {
		t.Error("HExists return != true")
	}

	tester.redis.Del(tester.ctx, k)
}

func TestLoggingConn_HGetAll(t *testing.T) {}

func TestLoggingConn_HGetAllMap(t *testing.T) {}

func TestLoggingConn_HGetString(t *testing.T) {}

func TestLoggingConn_HIncrBy(t *testing.T) {
	k := "TestLoggingConn_HExists"
	f := "field"
	v := 12345

	err := tester.redis.HSet(tester.ctx, k, f, v)
	t.Log("HSet return ", err)
	if err != nil {
		t.Error(err)
	}

	delta := int64(10)
	iv, err := tester.redis.HIncrBy(tester.ctx, k, f, delta)
	t.Log("HIncrBy return ", iv, err)
	if err != nil {
		t.Error(err)
	}

	if iv != int64(v)+delta {
		t.Error("HIncrBy afterIncrInt wrong")
	}

	tester.redis.Del(tester.ctx, k)
}

func TestLoggingConn_HIncrByFloat(t *testing.T) {
	k := "TestLoggingConn_HExistsTestLoggingConn_HIncrByFloat"
	f1 := "field_int"
	v1 := 12345.12
	f2 := "field_str"
	v2 := "12345.12"
	f3 := "field_error_str"
	v3 := "1ui1u"

	err := tester.redis.HMSet(tester.ctx, k, map[string]any{
		f1: v1,
		f2: v2,
		f3: v3,
	})
	t.Log("HMSet return ", err)
	if err != nil {
		t.Error(err)
	}

	//测试正常浮点数
	delta := 10.1
	iv, err := tester.redis.HIncrByFloat(tester.ctx, k, f1, 10.1)
	t.Log("HIncrByFloat return ", iv, err)
	if err != nil {
		t.Error(err)
	}

	if math.Abs(iv-v1-delta) > 1e-4 {
		t.Error("HIncrBy afterIncrInt wrong", iv, v1+delta)
	}

	//测试字符串类型写入的浮点数
	iv2, err := tester.redis.HIncrByFloat(tester.ctx, k, f2, 10.1)
	t.Log("HIncrByFloat return ", iv2, err)
	if err != nil {
		t.Error(err)
	}

	if math.Abs(iv2-v1-delta) > 1e-6 {
		t.Error("HIncrByFloat afterIncrInt wrong", iv2, v1+delta)
	}

	//测试不是浮点数类型的情况
	iv3, err := tester.redis.HIncrByFloat(tester.ctx, k, f3, 10.1)
	t.Log("HIncrByFloat return ", iv3, err)
	if err == nil {
		t.Error("HIncrByFloat should return ERR hash value is not a float")
	}

	tester.redis.Del(tester.ctx, k)
}

func TestLoggingConn_HKeys(t *testing.T) {
	k := "TestLoggingConn_HKeys"
	f := "field"
	v := "12345.12"

	err := tester.redis.HSet(tester.ctx, k, f, v)
	t.Log("HSet return ", err)
	if err != nil {
		t.Error(err)
	}

	keys, err := tester.redis.HKeys(tester.ctx, k)
	t.Log("HDel return ", err)
	if err != nil {
		t.Error(err)
	}

	if len(keys) != 1 {
		t.Error("HKeys return wrong")
	}

	keys, err = tester.redis.HKeys(tester.ctx, "TestLoggingConn_HKeys_missing")
	t.Log("HKeys TestLoggingConn_HKeys_missing return ", keys, err)
	if err != nil {
		t.Error(err)
	}

	if len(keys) != 0 {
		t.Error("HKeys return wrong")
	}

	tester.redis.Del(tester.ctx, k)
}

func TestLoggingConn_HLen(t *testing.T) {
	k := "TestLoggingConn_HKeys"
	f := "field"
	v := "12345.12"

	err := tester.redis.HSet(tester.ctx, k, f, v)
	t.Log("HSet return ", err)
	if err != nil {
		t.Error(err)
	}

	keys, err := tester.redis.HLen(tester.ctx, k)
	t.Log("HLen return ", keys, err)
	if err != nil {
		t.Error(err)
	}
	if keys != 1 {
		t.Error("HSet keys != 1")
	}
	tester.redis.HSet(tester.ctx, k, "field2", v)
	keys, err = tester.redis.HLen(tester.ctx, k)
	t.Log("HLen return ", keys, err)
	if err != nil {
		t.Error(err)
	}
	if keys != 2 {
		t.Error("HLen keys != 2")
	}
	tester.redis.Del(tester.ctx, k)
}

func TestLoggingConn_HMGET(t *testing.T) {}

func TestLoggingConn_HMGETString(t *testing.T) {}

func TestLoggingConn_HMSet(t *testing.T) {}

func TestLoggingConn_HMSetWithDuration(t *testing.T) {}

func TestLoggingConn_HScan(t *testing.T) {
	k := "TestLoggingConn_HScan"
	v := map[string]string{
		"TestLoggingConn_XHScan_f1": "zxc",
		"TestLoggingConn_XHScan_f2": "cvb",
		"TestLoggingConn_HScan_f1":  "qwe",
		"TestLoggingConn_HScan_f2":  "asd",
		"TestLoggingConn_HScan_f3":  "dfg",
		"TestLoggingConn_HScan_f4":  "ghj",
		"TestLoggingConn_YHScan_f1": "jkl",
		"TestLoggingConn_YHScan_f2": "bnm",
	}

	err := tester.redis.HMSet(tester.ctx, k, v)
	t.Log("HMSet return ", err)
	if err != nil {
		t.Error(err)
	}

	fv, c, err := tester.redis.HScan(tester.ctx, k, 0, "TestLoggingConn_HScan_*", 1)
	t.Log("HScan return ", fv, c, err)
	if err != nil {
		t.Error(err)
	}
}

func TestLoggingConn_HSet(t *testing.T) {}

func TestLoggingConn_HSetNx(t *testing.T) {
	k := "TestLoggingConn_HSetNx"
	f := "field"
	v := "12345.12"

	ok, err := tester.redis.HSetNx(tester.ctx, k, f, v)
	t.Log("HSetNx return ", ok, err)
	if err != nil {
		t.Error(err)
	}

	ok2, err := tester.redis.HSetNx(tester.ctx, k, f, v)
	t.Log("HSetNx2 return ", ok2, err)
	if ok2 == true {
		t.Error("HSetNx exists should false")
	}

	tester.redis.Del(tester.ctx, k)
}

func TestLoggingConn_HSetNxWithDuration(t *testing.T) {
	k := "TestLoggingConn_HSetNx"
	f := "field"
	v := "12345.12"

	ok, err := tester.redis.HSetNxWithDuration(tester.ctx, k, f, v, 10*time.Second)
	t.Log("HSetNxWithDuration return ", ok, err)
	if err != nil {
		t.Error(err)
	}

	ok2, err := tester.redis.HSetNxWithDuration(tester.ctx, k, f, v, 10*time.Second)
	t.Log("HSetNxWithDuration return ", ok2, err)
	if ok2 == true {
		t.Error("HSetNxWithDuration exists should false")
	}

	tester.redis.Del(tester.ctx, k)
}

func TestLoggingConn_HValues(t *testing.T) {
	k := "TestLoggingConn_HValues"
	f := "field"
	v := "12345.12"

	err := tester.redis.HSet(tester.ctx, k, f, v)
	t.Log("HSet return ", err)
	if err != nil {
		t.Error(err)
	}

	keys, err := tester.redis.HValues(tester.ctx, k)
	t.Log("HDel return ", err)
	if err != nil {
		t.Error(err)
	}

	if len(keys) != 1 {
		t.Error("HKeys return wrong")
	}

	tester.redis.Del(tester.ctx, k)
}

func TestLoggingConn_PipelineHGetAll(t *testing.T) {}

func TestLoggingConn_PipelineHGetAllMap(t *testing.T) {}

func TestLoggingConn_PipelineHLenMap(t *testing.T) {}

func TestLoggingConn_PipelineHMSet(t *testing.T) {}

func TestLoggingConn_PipelineHMSetWithDuration(t *testing.T) {}

func TestLoggingConn_PipelineHMSetWithDurations(t *testing.T) {}
