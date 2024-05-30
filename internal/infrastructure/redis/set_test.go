package redis

import (
	"testing"
)

func TestLoggingConn_PipelineSAdd(t *testing.T) {
	k := "TestLoggingConn_PipelineSAdd"
	values := []any{
		"12345",
		"56789",
	}

	err := tester.redis.PipelineSAdd(tester.ctx, map[string]interface{}{
		k: values,
	})
	t.Log("PipelineSAdd return ", err)
	if err != nil {
		t.Error(err)
	}

	members, err := tester.redis.SMembers(tester.ctx, k)
	t.Log("SMembers return ", members, err)
	if err != nil {
		t.Error(err)
	}

	if len(members) != len(values) {
		t.Error("SMembers return != values")
	}

	tester.redis.Del(tester.ctx, k)
}

func TestLoggingConn_PipelineSMembers(t *testing.T) {
	keys := []string{
		"TestLoggingConn_PipelineSMembers_1",
		"TestLoggingConn_PipelineSMembers_2",
		"TestLoggingConn_PipelineSMembers_3",
		"TestLoggingConn_PipelineSMembers_4",
	}

	values := []any{
		"12345",
		"12345",
		"12345",
		"56789",
	}

	err := tester.redis.PipelineSAdd(tester.ctx, map[string]interface{}{
		keys[0]: values,
		keys[1]: values,
		keys[2]: values,
	})
	result, err := tester.redis.PipelineSMembers(tester.ctx, keys)
	t.Log("PipelineSAdd return ", result, err)
	if err != nil {
		t.Error(err)
	}

	err = tester.redis.MDel(tester.ctx, keys)
	if err != nil {
		t.Error(err)
	}

	tester.redis.MDel(tester.ctx, keys)
}

func TestLoggingConn_SAdd(t *testing.T) {
	k := "TestLoggingConn_SAdd"
	values := []any{
		"12345",
		"56789",
	}

	cards, err := tester.redis.SAdd(tester.ctx, k, values...)
	t.Log("PipelineSAdd return ", cards, err)
	if err != nil {
		t.Error(err)
	}
	if cards != len(values) {
		t.Error("SAdd return cards != len(values)")
	}

	tester.redis.Del(tester.ctx, k)
}

func TestLoggingConn_SCard(t *testing.T) {
	k := "TestLoggingConn_SCard"
	values := []any{
		"12345",
		"12345",
		"12345",
		"56789",
	}

	cards1, err := tester.redis.SAdd(tester.ctx, k, values...)
	t.Log("PipelineSAdd return ", cards1, err)
	if err != nil {
		t.Error(err)
	}

	cards2, err := tester.redis.SCard(tester.ctx, k)
	t.Log("SCard return ", cards2, err)
	if err != nil {
		t.Error(err)
	}
	if cards1 != cards2 {
		t.Error("SCard return cards1 != cards2")
	}

	mk := "TestLoggingConn_SAdd_missing"
	cards3, err := tester.redis.SCard(tester.ctx, mk)
	t.Log("SCard return ", cards3, err)
	if err != nil {
		t.Error(err)
	}
	if cards3 != 0 {
		t.Error("SCard return cards3 != 0")
	}

	tester.redis.Del(tester.ctx, k)
	tester.redis.Del(tester.ctx, mk)
}

func TestLoggingConn_SIsMember(t *testing.T) {

}

func TestLoggingConn_SMembers(t *testing.T) {
	k := "TestLoggingConn_SMembers"
	values := []any{
		"12345",
		"12345",
		"12345",
		"56789",
	}

	cards1, err := tester.redis.SAdd(tester.ctx, k, values...)
	t.Log("PipelineSAdd return ", cards1, err)
	if err != nil {
		t.Error(err)
	}

	members, err := tester.redis.SMembers(tester.ctx, k)
	t.Log("SCard return ", members, err)
	if err != nil {
		t.Error(err)
	}
	if len(members) != cards1 {
		t.Error("SMembers return len(members) != cards1")
	}

	exist, err := tester.redis.SIsMember(tester.ctx, k, "56789")
	t.Log("SIsMember return ", exist, err)
	if err != nil {
		t.Error(err)
	}

	if exist != true {
		t.Error("SIsMember 56789 should return true")
	}

	exist, e2 := tester.redis.SIsMember(tester.ctx, k, "asdafasd")
	t.Log("SCard return ", exist, e2)
	if e2 != nil {
		t.Error(e2)
	}

	if exist == true {
		t.Error("SIsMember asdafasd should return false")
	}

	exist, e3 := tester.redis.SIsMember(tester.ctx, "111111", "asdafasd")
	t.Log("SIsMember return ", exist, e3)

	if exist != false {
		t.Error("SIsMember Key 111111 should return false")
	}

	tester.redis.Del(tester.ctx, k)
}

func TestLoggingConn_SRem(t *testing.T) {
	k := "TestLoggingConn_SRem"
	values := []any{
		"12345",
		"56789",
	}

	cards1, err := tester.redis.SAdd(tester.ctx, k, values...)
	t.Log("PipelineSAdd return ", cards1, err)
	if err != nil {
		t.Error(err)
	}

	removed, err := tester.redis.SRem(tester.ctx, k, values...)
	t.Log("SRem return ", cards1, err)
	if err != nil {
		t.Error(err)
	}

	if removed != len(values) {
		t.Error("SRem should return removed count ", len(values))
	}
	tester.redis.Del(tester.ctx, k)

	mk := "TestLoggingConn_SRem_missing"
	removed2, err := tester.redis.SRem(tester.ctx, mk, values...)
	t.Log("TestLoggingConn_SRem_missing return ", removed2, err)
	if err != nil {
		t.Error(err)
	}

	if removed2 != 0 {
		t.Error("SRem should return removed count ", 0)
	}
}
