package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"time"

	"github.com/gomodule/redigo/redis"
)

// ZAdd 向有序集合添加一个或多个成员，或者更新已存在成员的分数,
// Redis ZADD 命令用于将一个或多个成员元素及其分数值加入到有序集当中。
// 如果某个成员已经是有序集的成员，那么更新这个成员的分数值，并通过重新插入这个成员元素，来保证该成员在正确的位置上。
// 分数值可以是整数值或双精度浮点数。
// 如果有序集合 key 不存在，则创建一个空的有序集并执行 ZADD 操作。
// 当 key 存在但不是有序集类型时，返回一个错误。
// 注意： 在 Redis 2.4 版本以前， ZADD 每次只能添加一个元素。
func (c *LoggingConn) ZAdd(ctx context.Context, key string, score interface{}, member interface{}) error {
	_, err := c.DoWithTimeout(ctx, "ZADD", key, score, member)
	return err
}

// ZAddMulti 增加多个member和score对到指定有序集合中，memberScores的key代表member成员，value代表member的score分数
func (c *LoggingConn) ZAddMulti(ctx context.Context, key string, memberScores map[interface{}]interface{}) error {
	var args []interface{}
	args = append(args, key)
	for member, score := range memberScores {
		args = append(args, score, member)
	}
	_, err := c.DoWithTimeout(ctx, "ZADD", args...)
	return err
}

// ZRank 返回有序集中指定成员的排名。其中有序集成员按分数值递增(从小到大)顺序排列。
func (c *LoggingConn) ZRank(ctx context.Context, key string, member interface{}) (int, error) {
	return redis.Int(c.DoWithTimeout(ctx, "ZRANK", key, member))
}

// ZRevRank 返回有序集合中指定成员的排名，有序集成员按分数值递减(从大到小)排序
func (c *LoggingConn) ZRevRank(ctx context.Context, key string, member string) (int, error) {
	return redis.Int(c.DoWithTimeout(ctx, "ZREVRANK", key, member))
}

// ZScore 返回有序集中，成员的分数
func (c *LoggingConn) ZScore(ctx context.Context, key string, member interface{}) (float64, error) {
	return redis.Float64(c.DoWithTimeout(ctx, "ZSCORE", key, member))
}

// ZCard 获取有序集合的成员数，当 key 存在且是有序集类型时，返回有序集的基数。 当 key 不存在时，返回 0 。
func (c *LoggingConn) ZCard(ctx context.Context, key string) (int, error) {
	return redis.Int(c.DoWithTimeout(ctx, "ZCARD", key))
}

// ZCount 计算在有序集合中指定区间分数的成员数，返回分数值在 min 和 max 之间的成员的数量。
func (c *LoggingConn) ZCount(ctx context.Context, key string, minScore int, maxScore int) (int, error) {
	return redis.Int(c.DoWithTimeout(ctx, "ZCOUNT", key, minScore, maxScore))
}

// ZRange 计算在有序集合中指定区间分数的成员数，返回下标值在 min 和 max 之间的每个成员
// Redis ZRANGE 返回有序集中，指定区间内的成员。
// 其中成员的位置按分数值递增(从小到大)来排序。
// 具有相同分数值的成员按字典序(lexicographical order )来排列。
// 如果你需要成员按
// 值递减(从大到小)来排列，请使用 ZREVRANGE 命令。
// 下标参数 start 和 stop 都以 0 为底，也就是说，以 0 表示有序集第一个成员，以 1 表示有序集第二个成员，以此类推。
// 你也可以使用负数下标，以 -1 表示最后一个成员， -2 表示倒数第二个成员，以此类推。
func (c *LoggingConn) ZRange(ctx context.Context, key string, start int, stop int) ([]string, error) {
	reply, err := c.DoWithTimeout(ctx, "ZRANGE", key, start, stop)
	return redis.Strings(reply, err)
}

// ZRangeWithScores 计算在有序集合中指定区间分数的成员数，返回下标值在 min 和 max 之间的每个成员和分数关系
func (c *LoggingConn) ZRangeWithScores(ctx context.Context, key string, start int, stop int) (map[string]float64, error) {
	reply, err := c.DoWithTimeout(ctx, "ZRANGE", key, start, stop, "WITHSCORES")
	return Float64Map(reply, err)
}

// ZRangeByScore 通过分数返回有序集合指定区间内的成员，返回成员名称列表
// Redis ZRangeByScore 返回有序集合中指定分数区间的成员列表。有序集成员按分数值递增(从小到大)次序排列。
// 具有相同分数值的成员按字典序来排列(该属性是有序集提供的，不需要额外的计算)。
// 默认情况下，区间的取值使用闭区间 (小于等于或大于等于)，你也可以通过给参数前增加 ( 符号来使用可选的开区间 (小于或大于)。
func (c *LoggingConn) ZRangeByScore(ctx context.Context, key string, minScore interface{}, maxScore interface{}) ([]string, error) {
	reply, err := c.DoWithTimeout(ctx, "ZRANGEBYSCORE", key, minScore, maxScore)
	return redis.Strings(reply, err)
}

// ZRangeByScoreWithScores 返回有序集合中指定分数区间的成员列表。返回成员名称和分数的map
// Redis ZRangeByScoreWithScores 返回有序集合中指定分数区间的成员列表。有序集成员按分数值递增(从小到大)次序排列。
func (c *LoggingConn) ZRangeByScoreWithScores(ctx context.Context, key string, minScore interface{}, maxScore interface{}) (map[string]float64, error) {
	reply, err := c.DoWithTimeout(ctx, "ZRANGEBYSCORE", key, minScore, maxScore, "WITHSCORES")
	return Float64Map(reply, err)
}

// ZRangeByLex 通过字典区间返回有序集合的成员。返回指定区间内的元素列表。此指令适用于分数相同的有序集合中,LEX结尾的指令是要求分数必须相同.
// 注意：不要在分数不一致的SortSet集合中去使用 ZRANGEBYLEX 指令,因为获取的结果并不准确。
// 参考：https://blog.csdn.net/qq_32617703/article/details/103548791
// 指令	是否必须	说明
// ZRANGEBYLEX	是	指令
// key	 是	有序集合键名称
// min	 是	字典中排序位置较小的成员,必须以”[“开头,或者以”(“开头,可使用”-“代替
// max	 是	字典中排序位置较大的成员,必须以”[“开头,或者以”(“开头,可使用”+”代替
// LIMIT	 否	返回结果是否分页,指令中包含LIMIT后offset、count必须输入
// offset 否	返回结果起始位置
// count	 否	返回结果数量
//
// [符号
// [min 表示返回的结果中包含 min 值
// [max 表示返回的结果中包含 max 值
// ( 符号
// (min 表示返回的结果中不包含 min 值
// (max 表示返回的结果中不包含 max 值
func (c *LoggingConn) ZRangeByLex(ctx context.Context, key string, min string, max string) ([]string, error) {
	reply, err := c.DoWithTimeout(ctx, "ZRANGEBYLEX", key, min, max)
	return redis.Strings(reply, err)
}

// ZRangeByLexWithLimit 通过字典区间返回有序集合的成员。返回指定区间内的元素列表。此指令适用于分数相同的有序集合中,LEX结尾的指令是要求分数必须相同.
func (c *LoggingConn) ZRangeByLexWithLimit(ctx context.Context, key string, min string, max string, offset int, count int) ([]string, error) {
	reply, err := c.DoWithTimeout(ctx, "ZRANGEBYLEX", key, min, max, "LIMIT", offset, count)
	return redis.Strings(reply, err)
}

// ZRem 移除有序集合中的一个或多个成员，当 key 存在但不是有序集类型时，返回一个错误。返回被成功移除的成员的数量，不包括被忽略的成员。
func (c *LoggingConn) ZRem(ctx context.Context, key string, members ...interface{}) (int, error) {
	return redis.Int(c.DoWithTimeout(ctx, "ZRANK", redis.Args{}.Add(key).AddFlat(members)...))
}

// ZRemRangeByLex 命令用于删除成员名称按字典由低到高排序介于min 和 max 之间的所有成员（集合中所有成员的分数相同）。
// 不要在成员分数不同的有序集合中使用此命令，因为它是基于分数一致的有序集合设计的，如果使用，会导致删除的结果不正确。
// 返回被成功移除的成员的数量，不包括被忽略的成员。
func (c *LoggingConn) ZRemRangeByLex(ctx context.Context, key string, min string, max string) (int, error) {
	return redis.Int(c.DoWithTimeout(ctx, "ZREMRANGEBYLEX", key, min, max))
}

// ZRemRangeByRank 命令用于移除有序集中，指定排名(rank)区间内的所有成员。返回被移除成员的数量。
func (c *LoggingConn) ZRemRangeByRank(ctx context.Context, key string, start int, stop int) (int, error) {
	return redis.Int(c.DoWithTimeout(ctx, "ZREMRANGEBYRANK", key, start, stop))
}

// ZRemRangeByScore 命令用于移除有序集中，指定分数（score）区间内的所有成员。返回被移除成员的数量。
func (c *LoggingConn) ZRemRangeByScore(ctx context.Context, key string, start int, stop int) (int, error) {
	return redis.Int(c.DoWithTimeout(ctx, "ZREMRANGEBYSCORE", key, start, stop))
}

// ZrangeNoScore 返回有序集中指定区间内的成员，通过索引，分数从高到低
func (c *LoggingConn) ZrangeNoScore(ctx context.Context, key string, start, stop int) ([]string, error) {
	return redis.Strings(c.DoWithTimeout(ctx, "ZRANGE", key, start, stop))
}

func (c *LoggingConn) ZRevRange(ctx context.Context, key string, min, max int) (map[string]string, error) {
	byteList, err := redis.Strings(c.DoWithTimeout(ctx, "ZREVRANGE", key, min, max, "WITHSCORES"))
	if err != nil {
		return nil, err
	}

	list := make(map[string]string, len(byteList)/2)
	for k, v := range byteList {
		if k%2 != 0 {
			continue
		}
		list[v] = byteList[k+1]
	}
	return list, nil
}

// ZRevRangeByScore 返回有序集中指定分数区间内的所有的成员。有序集成员按分数值递减(从大到小)的次序排列。
// 具有相同分数值的成员按字典序的逆序(reverse lexicographical order )排列。
// 除了成员按分数值递减的次序排列这一点外， ZREVRANGEBYSCORE 命令的其他方面和 ZRANGEBYSCORE 命令一样。
func (c *LoggingConn) ZRevRangeByScore(ctx context.Context, key string, max interface{}, min interface{}) ([]string, error) {
	maxScore, err := cast.ToFloat64E(max)
	if err != nil {
		return nil, errors.New("ZRevRangeByScore param error: " + err.Error())
	}
	minScore, err := cast.ToFloat64E(min)
	if err != nil {
		return nil, errors.New("ZRevRangeByScore param error: " + err.Error())
	}
	if maxScore < minScore {
		tmp := minScore
		minScore = maxScore
		maxScore = tmp
	}
	reply, err := c.DoWithTimeout(ctx, "ZREVRANGEBYSCORE", key, maxScore, minScore)
	return redis.Strings(reply, err)
}

// ZRevRangeByScoreWithScores 返回有序集中指定分数区间内的所有的成员及其分数。有序集成员按分数值递减(从大到小)的次序排列。
func (c *LoggingConn) ZRevRangeByScoreWithScores(ctx context.Context, key string, max interface{}, min interface{}) (map[string]float64, error) {
	maxScore, err := cast.ToFloat64E(max)
	if err != nil {
		return nil, errors.New("ZRevRangeByScore param error: " + err.Error())
	}
	minScore, err := cast.ToFloat64E(min)
	if err != nil {
		return nil, errors.New("ZRevRangeByScore param error: " + err.Error())
	}
	if maxScore < minScore {
		tmp := minScore
		minScore = maxScore
		maxScore = tmp
	}
	reply, err := c.DoWithTimeout(ctx, "ZREVRANGEBYSCORE", key, maxScore, minScore, "WITHSCORES")
	return Float64Map(reply, err)
}

func (c *LoggingConn) PipelineZRange(ctx context.Context, scoreStart, scoreEnd interface{}, keys ...string) (map[string][]string, error) {
	startTime := time.Now()
	conn, err := c.Pool.GetContext(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	var args []any
	for _, key := range keys {
		args = append(args, fmt.Sprintf("ZRANGEBYSCORE %s %v %v", key, scoreStart, scoreEnd))
		_ = conn.Send("ZRANGEBYSCORE", key, scoreStart, scoreEnd)
	}

	err = conn.Flush()

	container := make(map[string][]string, len(keys))
	if err == nil {
		for _, key := range keys {
			reply, err := redis.Strings(redis.ReceiveContext(conn, ctx))
			if err != nil {
				continue
			}
			container[key] = reply
		}
	}

	c.log(ctx, logAttr{
		method:    "PipelineZRange",
		startTime: startTime,
		command:   "Pipeline",
		args:      args,
		reply:     stringifyReply(container),
		err:       err,
	})
	return container, nil
}

// ZScan 迭代有序集合中的元素（包括元素成员和元素分值），
// cursor - 游标。
// pattern - 匹配的模式。
// count - 指定从数据集里返回多少元素，默认值为 10
func (c *LoggingConn) ZScan(ctx context.Context, key string, cursor int, pattern string, count int) (fieldScores map[string]float64, latestCursor int, err error) {
	var args []any
	args = append(args, key, cursor)
	if pattern != "" && pattern != "*" {
		args = append(args, "MATCH", pattern)
	}
	if count > 0 {
		args = append(args, "COUNT", count)
	}

	values, err := redis.Values(c.DoWithTimeout(ctx, "ZSCAN", args...))
	if err != nil {
		return nil, 0, err
	}
	if values == nil || len(values) < 2 {
		return nil, 0, nil
	}
	latestCursor, _ = redis.Int(values[0], err)
	m, err := Float64Map(values[1], err)
	return m, latestCursor, err
}

func (c *LoggingConn) PipelineZAddWithDuration(ctx context.Context, keys []string, setData map[string]map[interface{}]int, duration time.Duration) error {
	startTime := time.Now()
	conn, err := c.Pool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	var args []any
	seconds := int(duration.Seconds())
	for _, key := range keys {
		vals := []interface{}{}
		vals = append(vals, key)
		for member, score := range setData[key] {
			vals = append(vals, score, member)
		}
		_ = conn.Send("ZADD", vals...)
		if seconds > 0 {
			_ = conn.Send("EXPIRE", redis.Args{}.Add(key).Add(seconds)...)
		}

		if seconds > 0 {
			args = append(args, fmt.Sprintf("ZADD %s EXPIRE %d", key, seconds))
		} else {
			args = append(args, fmt.Sprintf("ZADD %s", key))
		}
	}

	err = conn.Flush()
	c.log(ctx, logAttr{
		method:    "PipelineZAddWithDuration",
		startTime: startTime,
		command:   "Pipeline",
		args:      args,
		reply:     "",
		err:       err,
	})
	return err
}
