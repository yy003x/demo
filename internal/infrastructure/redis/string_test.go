package redis

import (
	"testing"
	"time"
)

func TestLoggingConn_MSetNx(t *testing.T) {
	k := "TestLoggingConn_MSetNx"
	v := "12345"

	tester.redis.Del(tester.ctx, k)
	ok, err := tester.redis.MSetNx(tester.ctx, k, v)
	t.Log("MSetNx return ", ok, err)
	if err != nil || !ok {
		t.Error(err)
	}

	ok, err = tester.redis.MSetNx(tester.ctx, k, v)
	t.Log("MSetNx return ", ok, err)
	if err != nil {
		t.Error(err)
	}
	if ok {
		t.Error("MSetNx should false")
	}
	tester.redis.Del(tester.ctx, k)
}

func TestLoggingConn_GetSet(t *testing.T) {
	k := "TestLoggingConn_GetSet"
	v := "12345"

	old, err := tester.redis.GetSet(tester.ctx, k, v)
	t.Log("GetSet return ", old, err)
	if err != nil {
		t.Error(err)
	}

	old, err = tester.redis.GetSet(tester.ctx, k, v)
	t.Log("GetSet return ", old, err)
	if err != nil {
		t.Error(err)
	}
	tester.redis.Del(tester.ctx, k)
}

func TestLoggingConn_SetBit(t *testing.T) {
	k := "TestLoggingConn_SetBit"
	offset := 100

	old, err := tester.redis.SetBit(tester.ctx, k, offset, 1)
	t.Log("GetSet return ", old, err)
	if err != nil {
		t.Error(err)
	}

	bit, err := tester.redis.GetBit(tester.ctx, k, offset)
	t.Log("GetBit return ", bit, err)
	if err != nil {
		t.Error(err)
	}

	if bit != 1 {
		t.Error("GetBit != 1")
	}
	tester.redis.Del(tester.ctx, k)
}

func TestLoggingConn_Append(t *testing.T) {
	k := "TestLoggingConn_GetSet"
	v := "12345"

	err := tester.redis.Set(tester.ctx, k, v)
	t.Log("GetSet return ", err)
	if err != nil {
		t.Error(err)
	}

	apd := "111"
	l, err := tester.redis.Append(tester.ctx, k, apd)
	t.Log("Append return ", l, err)
	if err != nil {
		t.Error(err)
	}

	if l != len(v)+len(apd) {
		t.Error("l != len(v)+len(apd)")
	}

	s, err := tester.redis.GetString(tester.ctx, k)
	if s != v+apd {
		t.Error("s != v+apd")
	}
	tester.redis.Del(tester.ctx, k)
}

func TestLoggingConn_MSetEx(t *testing.T) {
	tester.redis.Del(tester.ctx, "a")
	tester.redis.Del(tester.ctx, "b")
	tester.redis.Del(tester.ctx, "c")
	tester.redis.Del(tester.ctx, "d")
	tester.redis.Del(tester.ctx, "e")
	values := make(map[string]any, 5)
	values["a"] = "1"
	values["b"] = "2"
	values["c"] = "3"
	values["d"] = "4"
	values["e"] = "5"
	err := tester.redis.MSetEx(tester.ctx, time.Second*10, values)
	t.Log("MSetEx return ", err)
	if err != nil {
		t.Error(err)
	}

}
