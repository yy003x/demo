package redis

import (
	"context"
	"errors"
	"github.com/gomodule/redigo/redis"
	"time"
)

//LSet 通过索引来设置元素的值。当索引参数超出范围，或对一个空列表进行 LSET 时，返回一个错误。
func (c *LoggingConn) LSet(ctx context.Context, key string, index int, value string) error {
	_, err := c.DoWithTimeout(ctx, "LSET", key, index, value)
	return err
}

//LRem 移除列表元素,根据参数 COUNT 的值，移除列表中与参数 VALUE 相等的元素。返回被移除元素的数量。列表不存在时返回 0 。
//COUNT 的值可以是以下几种：
//count > 0 : 从表头开始向表尾搜索，移除与 VALUE 相等的元素，数量为 COUNT 。
//count < 0 : 从表尾开始向表头搜索，移除与 VALUE 相等的元素，数量为 COUNT 的绝对值。
//count = 0 : 移除表中所有与 VALUE 相等的值。
func (c *LoggingConn) LRem(ctx context.Context, key string, count int, value string) (int, error) {
	return redis.Int(c.DoWithTimeout(ctx, "LREM", key, count, value))
}

//LTrim 对一个列表进行修剪(trim)，就是说，让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除。
//下标 0 表示列表的第一个元素，以 1 表示列表的第二个元素，以此类推。 你也可以使用负数下标，以 -1 表示列表的最后一个元素， -2 表示列表的倒数第二个元素，以此类推。
func (c *LoggingConn) LTrim(ctx context.Context, key string, start int, stop int) error {
	_, err := c.DoWithTimeout(ctx, "LTRIM", key, start, stop)
	return err
}

//LPop 移出并获取列表的第一个元素，返回列表的第一个元素。当列表 key 不存在时，返回 nil 。
func (c *LoggingConn) LPop(ctx context.Context, key string) (string, error) {
	return redis.String(c.DoWithTimeout(ctx, "LPOP", key))
}

//BLPop 移出并获取列表的第一个元素，如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
//如果列表为空，返回一个 nil 。 否则，返回一个含有两个元素的列表，第一个元素是被弹出元素所属的 key ，第二个元素是被弹出元素的值。
//
//非阻塞行为
//当 BLPOP 被调用时，如果给定 key 内至少有一个非空列表，那么弹出遇到的第一个非空列表的头元素，并和被弹出元素所属的列表的名字 key 一起，组成结果返回给调用者。
//当存在多个给定 key 时， BLPOP 按给定 key 参数排列的先后顺序，依次检查各个列表。 我们假设 key list1 不存在，而 list2 和 list3 都是非空列表。考虑以下的命令：
//BLPOP list1 list2 list3 0
//BLPOP 保证返回一个存在于 list2 里的元素（因为它是从 list1 –> list2 –> list3 这个顺序查起的第一个非空列表）。
//
//阻塞行为
//如果所有给定 key 都不存在或包含空列表，那么 BLPOP 命令将阻塞连接， 直到有另一个客户端对给定的这些 key 的任意一个执行 LPUSH 或 RPUSH 命令为止。
//一旦有新的数据出现在其中一个列表里，那么这个命令会解除阻塞状态，并且返回 key 和弹出的元素值。
//当 BLPOP 命令引起客户端阻塞并且设置了一个非零的超时参数 timeout 的时候， 若经过了指定的 timeout 仍没有出现一个针对某一特定 key 的 push 操作，则客户端会解除阻塞状态并且返回一个 nil 的多组合值(multi-bulk value)。
//timeout 参数表示的是一个指定阻塞的最大秒数的整型值。当 timeout 为 0 是表示阻塞时间无限制。
func (c *LoggingConn) BLPop(ctx context.Context, timeout time.Duration, key string) (string, error) {
	seconds := int(timeout.Seconds())
	vals, err := redis.Strings(c.DoWithTimeout(ctx, "BLPOP", key, seconds))
	if err != nil {
		return "", err
	}
	if len(vals) > 1 {
		return vals[1], err
	} else {
		return "", errors.New("BLPop failed")
	}
}

//BLPopMulti 移出并获取列表的第一个元素，如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
//如果列表为空，返回一个 nil 。 否则，返回一个含有两个元素的列表，第一个元素是被弹出元素所属的 key ，第二个元素是被弹出元素的值。
//当给定多个 key 参数时，按参数 key 的先后顺序依次检查各个列表，弹出第一个非空列表的头元素。
//
//非阻塞行为
//当 BLPOP 被调用时，如果给定 key 内至少有一个非空列表，那么弹出遇到的第一个非空列表的头元素，并和被弹出元素所属的列表的名字 key 一起，组成结果返回给调用者。
//当存在多个给定 key 时， BLPOP 按给定 key 参数排列的先后顺序，依次检查各个列表。 我们假设 key list1 不存在，而 list2 和 list3 都是非空列表。考虑以下的命令：
//BLPOP list1 list2 list3 0
//BLPOP 保证返回一个存在于 list2 里的元素（因为它是从 list1 –> list2 –> list3 这个顺序查起的第一个非空列表）。
//
//阻塞行为
//如果所有给定 key 都不存在或包含空列表，那么 BLPOP 命令将阻塞连接， 直到有另一个客户端对给定的这些 key 的任意一个执行 LPUSH 或 RPUSH 命令为止。
//一旦有新的数据出现在其中一个列表里，那么这个命令会解除阻塞状态，并且返回 key 和弹出的元素值。
//当 BLPOP 命令引起客户端阻塞并且设置了一个非零的超时参数 timeout 的时候， 若经过了指定的 timeout 仍没有出现一个针对某一特定 key 的 push 操作，则客户端会解除阻塞状态并且返回一个 nil 的多组合值(multi-bulk value)。
//timeout 参数表示的是一个指定阻塞的最大秒数的整型值。当 timeout 为 0 是表示阻塞时间无限制。
func (c *LoggingConn) BLPopMulti(ctx context.Context, timeout time.Duration, keys ...string) (map[string]string, error) {
	seconds := int(timeout.Seconds())
	return redis.StringMap(c.DoWithTimeout(ctx, "BLPOP", redis.Args{}.AddFlat(keys).Add(seconds)...))
}

//RPop 移除列表的最后一个元素，返回值为移除的元素。返回被移除的元素。当列表不存在时，返回 nil 。
func (c *LoggingConn) RPop(ctx context.Context, key string) (string, error) {
	return redis.String(c.DoWithTimeout(ctx, "RPOP", key))
}

//BRPop 移出并获取列表的最后一个元素，如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。非阻塞行为
//当 BRPOP 被调用时，如果指定 key 内至少有一个非空列表，那么弹出第一个非空列表的头元素，并和被弹出元素所属的列表的名字一起，组成结果返回给调用者。
//当存在多个指定的 key 时，BRPOP 按指定 key 参数排列的先后顺序，依次检查各个列表。
//
//非阻塞行为
//当 BRPOP 被调用时，如果指定 key 内至少有一个非空列表，那么弹出第一个非空列表的头元素，并和被弹出元素所属的列表的名字一起，组成结果返回给调用者。
//当存在多个指定的 key 时，BRPOP 按指定 key 参数排列的先后顺序，依次检查各个列表。
//
//阻塞行为
//如果所有指定的 key 都不存在元素，那么 BRPOP 命令将阻塞连接，直到等待超时，或有另一个客户端对指定 key 的任意一个执行 LPUSH 或 RPUSH 命令为止。
func (c *LoggingConn) BRPop(ctx context.Context, timeout time.Duration, key string) (string, error) {
	seconds := int(timeout.Seconds())
	vals, err := redis.Strings(c.DoWithTimeout(ctx, "BRPOP", key, seconds))
	if err != nil {
		return "", err
	}
	if len(vals) > 1 {
		return vals[1], err
	} else {
		return "", errors.New("BRPop failed")
	}
}

//BRPopMulti 移出并获取列表的最后一个元素，如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
//当给定多个 key 参数时，按参数 key 的先后顺序依次检查各个列表，弹出第一个非空列表的头元素。
//
//非阻塞行为
//当 BRPOP 被调用时，如果指定 key 内至少有一个非空列表，那么弹出第一个非空列表的头元素，并和被弹出元素所属的列表的名字一起，组成结果返回给调用者。
//当存在多个指定的 key 时，BRPOP 按指定 key 参数排列的先后顺序，依次检查各个列表。
//
//阻塞行为
//如果所有指定的 key 都不存在元素，那么 BRPOP 命令将阻塞连接，直到等待超时，或有另一个客户端对指定 key 的任意一个执行 LPUSH 或 RPUSH 命令为止。
func (c *LoggingConn) BRPopMulti(ctx context.Context, timeout time.Duration, keys ...string) (map[string]string, error) {
	seconds := int(timeout.Seconds())
	return redis.StringMap(c.DoWithTimeout(ctx, "BRPOP", redis.Args{}.AddFlat(keys).Add(seconds)...))
}

//RPopLPush 命令用于移除列表的最后一个元素，并将该元素添加到另一个列表并返回。
func (c *LoggingConn) RPopLPush(ctx context.Context, source string, destination string) (string, error) {
	return redis.String(c.DoWithTimeout(ctx, "RPOPLPUSH", source, destination))
}

//BRPopLPush 移除列表的最后一个元素，返回值为移除的元素。返回被移除的元素。当列表不存在时，返回 nil。
//返回值:假如在指定时间内没有任何元素被弹出，则返回一个 nil 和等待时长。 反之，返回一个含有两个元素的列表，第一个元素是被弹出元素的值，第二个元素是等待时长。
func (c *LoggingConn) BRPopLPush(ctx context.Context, source string, destination string, timeout time.Duration) (string, error) {
	seconds := int(timeout.Seconds())
	v, err := c.DoWithTimeout(ctx, "BRPOPLPUSH", redis.Args{}.Add(source).Add(destination).Add(seconds)...)
	values, err := redis.Bytes(v, err)
	return string(values), err
}

//LPush 命令将一个或多个值插入到列表头部。 如果 key 不存在，一个空列表会被创建并执行 LPUSH 操作。 当 key 存在但不是列表类型时，返回一个错误。
//返回执行 LPUSH 命令后，列表的长度。
func (c *LoggingConn) LPush(ctx context.Context, key string, values ...interface{}) (int, error) {
	return redis.Int(c.DoWithTimeout(ctx, "LPUSH", redis.Args{}.Add(key).AddFlat(values)...))
}

//RPush 命令用于将一个或多个值插入到列表的尾部(最右边)。返回执行 RPUSH 操作后，列表的长度。
//如果列表不存在，一个空列表会被创建并执行 RPUSH 操作。 当列表存在但不是列表类型时，返回一个错误。
//注意：在 Redis 2.4 版本以前的 RPUSH 命令，都只接受单个 value 值。
func (c *LoggingConn) RPush(ctx context.Context, key string, values ...interface{}) (int, error) {
	return redis.Int(c.DoWithTimeout(ctx, "RPUSH", redis.Args{}.Add(key).AddFlat(values)...))
}

//RPushX 命令用于将一个值插入到已存在的列表尾部(最右边)。如果列表不存在，操作无效。返回执行 RPushX 操作后，列表的长度。
func (c *LoggingConn) RPushX(ctx context.Context, key string, value interface{}) (int, error) {
	return redis.Int(c.DoWithTimeout(ctx, "RPUSHX", key, value))
}

//LPushX 将一个值插入到已存在的列表头部，列表不存在时操作无效。
//返回LPUSHX 命令执行之后，列表的长度。
func (c *LoggingConn) LPushX(ctx context.Context, key string, value interface{}) (int, error) {
	return redis.Int(c.DoWithTimeout(ctx, "LPUSHX", key, value))
}

//LRange 返回列表中指定区间内的元素，区间以偏移量 START 和 END 指定。 其中 0 表示列表的第一个元素， 1 表示列表的第二个元素，以此类推。
//你也可以使用负数下标，以 -1 表示列表的最后一个元素， -2 表示列表的倒数第二个元素，以此类推。
func (c *LoggingConn) LRange(ctx context.Context, key string, start int, stop int) ([]string, error) {
	return redis.Strings(c.DoWithTimeout(ctx, "LRANGE", key, start, stop))
}

//LRangeAll 返回列表中的所有元素
func (c *LoggingConn) LRangeAll(ctx context.Context, key string) ([]string, error) {
	return redis.Strings(c.DoWithTimeout(ctx, "LRANGE", key, 0, -1))
}

//LLen 命令用于返回列表的长度。 如果列表 key 不存在，则 key 被解释为一个空列表，返回 0 。 如果 key 不是列表类型，返回一个错误。
func (c *LoggingConn) LLen(ctx context.Context, key string) (int, error) {
	return redis.Int(c.DoWithTimeout(ctx, "LLEN", key))
}

//LIndex 命令用于通过索引获取列表中的元素。你也可以使用负数下标，以 -1 表示列表的最后一个元素， -2 表示列表的倒数第二个元素，以此类推。
//返回列表中下标为指定索引值的元素。 如果指定索引值不在列表的区间范围内，返回 nil 。
func (c *LoggingConn) LIndex(ctx context.Context, key string, index int) (string, error) {
	return redis.String(c.DoWithTimeout(ctx, "LINDEX", key, index))
}

//LInsertBefore 命令用于在列表的元素前插入元素。当指定元素不存在于列表中时，不执行任何操作。
//当列表不存在时，被视为空列表，不执行任何操作。
//如果 key 不是列表类型，返回一个错误。
//返回：如果命令执行成功，返回插入操作完成之后，列表的长度。 如果没有找到指定元素 ，返回 -1 。 如果 key 不存在或为空列表，返回 0 。
func (c *LoggingConn) LInsertBefore(ctx context.Context, key string, pivot interface{}, value interface{}) (int, error) {
	return redis.Int(c.DoWithTimeout(ctx, "LINSERT", key, "BEFORE", pivot, value))
}

//LInsertAfter 命令用于在列表的元素后插入元素。当指定元素不存在于列表中时，不执行任何操作。
//当列表不存在时，被视为空列表，不执行任何操作。
//如果 key 不是列表类型，返回一个错误。
//返回：如果命令执行成功，返回插入操作完成之后，列表的长度。 如果没有找到指定元素 ，返回 -1 。 如果 key 不存在或为空列表，返回 0 。
func (c *LoggingConn) LInsertAfter(ctx context.Context, key string, pivot interface{}, value interface{}) (int, error) {
	return redis.Int(c.DoWithTimeout(ctx, "LINSERT", key, "AFTER", pivot, value))
}
