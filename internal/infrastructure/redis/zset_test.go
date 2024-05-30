package redis

import (
	"testing"
)

func TestLoggingConn_PipelineZAddWithDuration(t *testing.T) {}

func TestLoggingConn_PipelineZRange(t *testing.T) {}

func beforeTest(t *testing.T) {
	k := "TestLoggingConn_ZSET"
	score := 1000
	member := "member"

	tester.redis.ZAddMulti(tester.ctx, k, map[any]any{
		"a": 10.1,
		"b": 20.2,
		"c": 30.3,
		"d": 40.4,
		"e": 50.5,
	})
	err := tester.redis.ZAdd(tester.ctx, k, score, member)
	t.Log("ZAdd return ", err)
	if err != nil {
		t.Error(err)
	}
}

func afterTest(t *testing.T) {
	k := "TestLoggingConn_ZSET"
	tester.redis.Del(tester.ctx, k)
}

func TestLoggingConn_ZAdd(t *testing.T) {
	beforeTest(t)
	afterTest(t)
}

func TestLoggingConn_ZAddMulti(t *testing.T) {
	beforeTest(t)
	k := "TestLoggingConn_ZSET"
	score2 := int(2000)
	member2 := "member2"
	score3 := int(3000)
	member4 := "member3"
	err := tester.redis.ZAddMulti(tester.ctx, k, map[any]any{
		member2: score2,
		member4: score3,
	})
	t.Log("ZAddMulti return ", err)
	if err != nil {
		t.Error(err)
	}
	afterTest(t)
}

func TestLoggingConn_ZCard(t *testing.T) {
	beforeTest(t)
	cards, err := tester.redis.ZCard(tester.ctx, "TestLoggingConn_ZSET")
	t.Log("ZAddMulti return ", err)
	if err != nil {
		t.Error(err)
	}
	if cards < 1 {
		t.Error("ZCard return error")
	}
	afterTest(t)
}

func TestLoggingConn_ZCount(t *testing.T) {
	beforeTest(t)
	cards, err := tester.redis.ZCount(tester.ctx, "TestLoggingConn_ZSET", 15, 25)
	t.Log("ZCount return ", cards, err)
	if err != nil {
		t.Error(err)
	}
	if cards != 1 {
		t.Error("ZCard return error")
	}
	afterTest(t)
}

func TestLoggingConn_ZRange(t *testing.T) {
	beforeTest(t)
	err := tester.redis.ZAddMulti(tester.ctx, "TestLoggingConn_ZSET", map[any]any{
		"qweqwe":        1000.2,
		"zxczxac":       2000.3,
		"cvbcvbass":     3000.45,
		"asadafafafafa": 4000.56,
	})
	cards, err := tester.redis.ZRange(tester.ctx, "TestLoggingConn_ZSET", 0, -1)
	t.Log("ZRange return ", cards, err)
	if err != nil {
		t.Error(err)
	}
	if len(cards) != 10 {
		t.Error("ZCard return error")
	}
	afterTest(t)
}

func TestLoggingConn_ZRangeWithScores(t *testing.T) {
	beforeTest(t)
	err := tester.redis.ZAddMulti(tester.ctx, "TestLoggingConn_ZSET", map[any]any{
		"qweqwe":        1000.2,
		"zxczxac":       2000.3,
		"cvbcvbass":     3000.45,
		"asadafafafafa": 4000.56,
	})
	cards, err := tester.redis.ZRangeWithScores(tester.ctx, "TestLoggingConn_ZSET", 0, -1)
	t.Log("ZRange return ", cards, err)
	if err != nil {
		t.Error(err)
	}

	if cards["qweqwe"] != 1000.2 {
		t.Error("ZRangeWithScores error")
	}
	afterTest(t)
}

func TestLoggingConn_ZRangeByLex(t *testing.T) {
	beforeTest(t)
	err := tester.redis.ZAddMulti(tester.ctx, "TestLoggingConn_ZSET", map[any]any{
		"a": 10,
		"b": 10,
		"c": 10,
		"d": 10,
	})
	cards, err := tester.redis.ZRangeByLex(tester.ctx, "TestLoggingConn_ZSET", "-", "[c")
	t.Log("ZRange return ", cards, err)
	if err != nil {
		t.Error(err)
	}
	if len(cards) != 3 {
		t.Error("ZCard return error")
	}
	afterTest(t)
}

func TestLoggingConn_ZRangeByScore(t *testing.T) {
	beforeTest(t)
	err := tester.redis.ZAddMulti(tester.ctx, "TestLoggingConn_ZSET", map[any]any{
		"a": 10,
		"b": 20,
		"c": 30,
		"d": 40,
	})
	cards, err := tester.redis.ZRangeByScore(tester.ctx, "TestLoggingConn_ZSET", 15, 25)
	t.Log("ZRange return ", cards, err)
	if err != nil {
		t.Error(err)
	}
	if len(cards) != 1 {
		t.Error("ZCard return error")
	}
	afterTest(t)
}

func TestLoggingConn_ZRangeByScoreWithScores(t *testing.T) {
	beforeTest(t)
	err := tester.redis.ZAddMulti(tester.ctx, "TestLoggingConn_ZSET", map[any]any{
		"a": 10.1,
		"b": 20.3,
		"c": 30.5,
		"d": 40.7,
	})
	cards, err := tester.redis.ZRangeByScoreWithScores(tester.ctx, "TestLoggingConn_ZSET", 15, 25)
	t.Log("ZRange return ", cards, err)
	if err != nil {
		t.Error(err)
	}
	if len(cards) != 1 {
		t.Error("ZCard return error")
	}
	afterTest(t)
}

func TestLoggingConn_ZRank(t *testing.T) {
	beforeTest(t)
	err := tester.redis.ZAddMulti(tester.ctx, "TestLoggingConn_ZSET", map[any]any{
		"0": 0,
		"a": 10,
		"b": 20,
		"c": 30,
		"d": 40,
		"e": 30,
	})
	rank, err := tester.redis.ZRank(tester.ctx, "TestLoggingConn_ZSET", "b")
	t.Log("ZRank return ", rank, err)
	if err != nil {
		t.Error(err)
	}
	if rank != 2 {
		t.Error("ZCard return error")
	}
	afterTest(t)
}

func TestLoggingConn_ZRem(t *testing.T) {
	beforeTest(t)
	_, err := tester.redis.ZRem(tester.ctx, "TestLoggingConn_ZSET", "member")
	if err != nil {
		t.Error(err)
	}
	afterTest(t)
}

func TestLoggingConn_ZRemRangeByLex(t *testing.T) {
	beforeTest(t)
	err := tester.redis.ZAddMulti(tester.ctx, "TestLoggingConn_ZSET", map[any]any{
		"a": 10,
		"b": 20,
		"c": 30,
		"d": 40,
	})
	_, err = tester.redis.ZRemRangeByLex(tester.ctx, "TestLoggingConn_ZSET", "-", "+")
	if err != nil {
		t.Error(err)
	}
	afterTest(t)
}

func TestLoggingConn_ZRemRangeByRank(t *testing.T) {
	beforeTest(t)
	err := tester.redis.ZAddMulti(tester.ctx, "TestLoggingConn_ZSET", map[any]any{
		"a": 10,
		"b": 20,
		"c": 30,
		"d": 40,
	})
	cards, err := tester.redis.ZRemRangeByRank(tester.ctx, "TestLoggingConn_ZSET", 1, 2)
	if err != nil {
		t.Error(err)
	}
	if cards != 2 {
		t.Error("ZCard return error")
	}
	afterTest(t)
}

func TestLoggingConn_ZRemRangeByScore(t *testing.T) {
	beforeTest(t)
	tester.redis.ZAddMulti(tester.ctx, "TestLoggingConn_ZSET", map[any]any{
		"a": 10,
		"b": 20,
		"c": 30,
		"d": 40,
		"e": 50,
	})
	cards, err := tester.redis.ZRemRangeByScore(tester.ctx, "TestLoggingConn_ZSET", 10, 20)
	t.Log("ZRemRangeByScore return ", cards, err)
	if err != nil {
		t.Error(err)
	}
	if cards != 2 {
		t.Error("ZCard return error")
	}
	afterTest(t)
}

func TestLoggingConn_ZRevRange(t *testing.T) {
	beforeTest(t)
	tester.redis.ZAddMulti(tester.ctx, "TestLoggingConn_ZSET", map[any]any{
		"a": 10.1,
		"b": 20.2,
		"c": 30.5,
		"d": 40,
		"e": 50,
	})
	cards, err := tester.redis.ZRevRange(tester.ctx, "TestLoggingConn_ZSET", 0, -1)
	t.Log("ZRevRange return ", cards, err)
	if err != nil {
		t.Error(err)
	}
	if len(cards) < 2 {
		t.Error("ZRevRangeByScore return error")
	}
	afterTest(t)
}

func TestLoggingConn_ZRevRangeByScore(t *testing.T) {
	beforeTest(t)
	cards, err := tester.redis.ZRevRangeByScore(tester.ctx, "TestLoggingConn_ZSET", 0, 10000)
	t.Log("ZRevRangeByScore return ", cards, err)
	if err != nil {
		t.Error(err)
	}
	if len(cards) < 2 {
		t.Error("ZRevRangeByScore return error")
	}
	afterTest(t)
}

func TestLoggingConn_ZRevRangeByScoreWithScores(t *testing.T) {
	beforeTest(t)
	cards, err := tester.redis.ZRevRangeByScoreWithScores(tester.ctx, "TestLoggingConn_ZSET", 0, 10000)
	t.Log("ZRevRangeByScoreWithScores return ", cards, err)
	if err != nil {
		t.Error(err)
	}
	if len(cards) < 2 {
		t.Error("ZCard return error")
	}
	afterTest(t)
}

func TestLoggingConn_ZRevRank(t *testing.T) {
	beforeTest(t)
	rank, err := tester.redis.ZRevRank(tester.ctx, "TestLoggingConn_ZSET", "b")
	t.Log("ZRevRank return ", rank, err)
	if err != nil {
		t.Error(err)
	}
	if rank < 1 {
		t.Error("ZRevRank return error")
	}
	afterTest(t)
}

func TestLoggingConn_ZScan(t *testing.T) {
	beforeTest(t)
	tester.redis.ZAddMulti(tester.ctx, "TestLoggingConn_ZSET", map[any]any{
		"TestLoggingConn_ZSET_a": 10,
		"TestLoggingConn_ZSET_b": 20,
		"TestLoggingConn_ZSET_c": 30,
		"d":                      40,
		"e":                      50,
	})
	cards, c, err := tester.redis.ZScan(tester.ctx, "TestLoggingConn_ZSET", 0, "TestLoggingConn_ZSET_*", 10)
	t.Log("ZScan return ", cards, c, err)
	if err != nil {
		t.Error(err)
	}

	if len(cards) < 1 {
		t.Error("ZScan return error")
	}
	afterTest(t)
}

func TestLoggingConn_ZScore(t *testing.T) {
	beforeTest(t)
	v := 20.2
	k := "TestLoggingConn_ZSET"
	tester.redis.ZAddMulti(tester.ctx, k, map[any]any{
		"b": v,
	})
	score, err := tester.redis.ZScore(tester.ctx, k, "b")
	t.Log("ZScore return ", score, err)
	if err != nil {
		t.Error(err)
	}
	if score != v {
		t.Error("ZScore != 20")
	}
	afterTest(t)
}
