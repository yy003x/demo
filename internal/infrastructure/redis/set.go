package redis

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

//SAdd 向集合添加一个或多个成员，返回被添加到集合中的新元素的数量，不包括被忽略的元素。
//add 命令将一个或多个成员元素加入到集合中，已经存在于集合的成员元素将被忽略。
//假如集合 key 不存在，则创建一个只包含添加的元素作成员的集合。
//当集合 key 不是集合类型时，返回一个错误。
//注意：在 Redis2.4 版本以前， SADD 只接受单个成员值。
func (c *LoggingConn) SAdd(ctx context.Context, key string, members ...interface{}) (int, error) {
	reply, err := c.DoWithTimeout(ctx, "SADD", redis.Args{}.Add(key).AddFlat(members)...)
	return redis.Int(reply, err)
}

//SMembers 命令返回集合中的所有的成员。 不存在的集合 key 被视为空集合。
func (c *LoggingConn) SMembers(ctx context.Context, key string) ([]string, error) {
	reply, err := c.DoWithTimeout(ctx, "SMEMBERS", key)
	return redis.Strings(reply, err)
}

//PipelineSAdd 通过管道方式批量给集合添加成员数据
func (c *LoggingConn) PipelineSAdd(ctx context.Context, setData map[string]interface{}) error {
	startTime := time.Now()
	conn, err := c.Pool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	var args []any
	for key, val := range setData {
		args = append(args, fmt.Sprintf("SADD %s", key))
		_ = conn.Send("SADD", redis.Args{}.Add(key).AddFlat(val)...)
	}

	err = conn.Flush()
	c.log(ctx, logAttr{
		method:    "PipelineSAdd",
		startTime: startTime,
		command:   "Pipeline",
		args:      args,
		reply:     "",
		err:       err,
	})
	return err
}

//PipelineSMembers 通过管道方式获取多个集合的成员
func (c *LoggingConn) PipelineSMembers(ctx context.Context, keys []string) (map[string][]string, error) {
	startTime := time.Now()
	conn, err := c.Pool.GetContext(ctx)
	if err != nil {
		if err == redis.ErrPoolExhausted {
			c.logger.Errorf("PipelineSMembers failed, %s", err.Error())
		}
		return nil, err
	}
	defer conn.Close()
	for _, key := range keys {
		_ = conn.Send("SMEMBERS", key)
	}
	if err := conn.Flush(); err != nil {
		return nil, err
	}

	var args []any
	containers := make(map[string][]string, len(keys))
	for _, key := range keys {
		args = append(args, fmt.Sprintf("SMEMBERS %s", key))
		values, err := redis.Strings(redis.ReceiveContext(conn, ctx))
		if err != nil {
			c.logger.WithContext(ctx).Errorf("Redis PipelineSMembers error %v", err)
			continue
		}

		if len(values) == 0 {
			c.logger.WithContext(ctx).Infof("Redis PipelineSMembers value empty, key: %v", key)
			continue
		}

		containers[key] = values
	}

	c.log(ctx, logAttr{
		method:    "PipelineSMembers",
		startTime: startTime,
		command:   "Pipeline",
		args:      args,
		reply:     stringifyReply(containers),
		err:       err,
	})
	return containers, nil
}

//SRem 命令用于移除集合中的一个或多个成员元素，不存在的成员元素会被忽略。返回被成功移除的元素的数量，不包括被忽略的元素。
//当 key 不是集合类型，返回一个错误。
func (c *LoggingConn) SRem(ctx context.Context, key string, members ...interface{}) (int, error) {
	reply, err := c.DoWithTimeout(ctx, "SREM", redis.Args{}.Add(key).AddFlat(members)...)
	return redis.Int(reply, err)
}

//SIsMember 判断 member 元素是否是集合 key 的成员
func (c *LoggingConn) SIsMember(ctx context.Context, key string, member interface{}) (bool, error) {
	reply, err := c.DoWithTimeout(ctx, "SISMEMBER", key, member)
	return redis.Bool(reply, err)
}

//SCard 获取集合的成员数，集合的数量。当集合 key 不存在时，返回 0 。
func (c *LoggingConn) SCard(ctx context.Context, key string) (int, error) {
	reply, err := c.DoWithTimeout(ctx, "SCARD", key)
	return redis.Int(reply, err)
}
