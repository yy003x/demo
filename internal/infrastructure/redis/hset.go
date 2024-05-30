package redis

import (
	"context"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

func (c *LoggingConn) HSet(ctx context.Context, key string, field, value interface{}) error {
	_, err := c.DoWithTimeout(ctx, "HSET", key, field, value)
	return err
}

//HSetNx 命令用于为哈希表中不存在的的字段赋值 。
//如果哈希表不存在，一个新的哈希表被创建并进行 HSET 操作。
//如果字段已经存在于哈希表中，操作无效。
//如果 key 不存在，一个新哈希表被创建并执行 HSETNX 命令。
func (c *LoggingConn) HSetNx(ctx context.Context, key string, field, value interface{}) (bool, error) {
	reply, err := c.DoWithTimeout(ctx, "HSETNX", key, field, value)
	return redis.Bool(reply, err)
}

//HSetNxWithDuration 命令用于为哈希表中不存在的的字段赋值 。
func (c *LoggingConn) HSetNxWithDuration(ctx context.Context, key string, field, value interface{}, duration time.Duration) (bool, error) {
	reply, err := c.DoWithTimeout(ctx, "HSETNX", key, field, value)
	if err == nil {
		err = c.Expire(ctx, key, duration)
	}
	return redis.Bool(reply, err)
}

func (c *LoggingConn) HMSet(ctx context.Context, key string, value interface{}) error {
	_, err := c.DoWithTimeout(ctx, "HMSET", redis.Args{}.Add(key).AddFlat(value)...)
	return err
}

func (c *LoggingConn) HMSetWithDuration(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	_, err := c.DoWithTimeout(ctx, "HMSET", redis.Args{}.Add(key).AddFlat(value)...)
	if err == nil {
		c.Expire(ctx, key, duration)
	}
	return err
}

func (c *LoggingConn) HDel(ctx context.Context, key string, field interface{}) error {
	_, err := c.DoWithTimeout(ctx, "HDEL", key, field)
	return err
}

func (c *LoggingConn) HGetString(ctx context.Context, key string, field interface{}) (string, error) {
	res, err := redis.String(c.DoWithTimeout(ctx, "HGET", key, field))
	return res, err
}

func (c *LoggingConn) HMGETString(ctx context.Context, key string, field interface{}) ([]string, error) {
	reply, err := c.DoWithTimeout(ctx, "HMGET", redis.Args{}.Add(key).AddFlat(field)...)
	return redis.Strings(reply, err)
}

func (c *LoggingConn) HMGET(ctx context.Context, key string, field interface{}) (interface{}, error) {
	reply, err := c.DoWithTimeout(ctx, "HMGET", redis.Args{}.Add(key).AddFlat(field)...)
	return reply, err
}

func (c *LoggingConn) HGetAll(ctx context.Context, key string, dest interface{}) error {
	reply, err := c.DoWithTimeout(ctx, "HGETALL", key)
	values, err := redis.Values(reply, err)
	if err != nil {
		return err
	}
	if len(values) == 0 {
		return redis.ErrNil
	}
	return redis.ScanStruct(values, dest)
}

func (c *LoggingConn) HGetAllMap(ctx context.Context, key string) (map[string]string, error) {
	return redis.StringMap(c.DoWithTimeout(ctx, "HGETALL", key))
}

func (c *LoggingConn) HGetAllMapInt(ctx context.Context, key string) (map[string]int, error) {
	return redis.IntMap(c.DoWithTimeout(ctx, "HGETALL", key))
}

//HExists 查看哈希表 key 中，指定的字段是否存在。
//如果哈希表含有给定字段，返回 1 。 如果哈希表不含有给定字段，或 key 不存在，返回 0 。
func (c *LoggingConn) HExists(ctx context.Context, key string, field string) (bool, error) {
	return redis.Bool(c.DoWithTimeout(ctx, "HEXISTS", key, field))
}

//HIncrBy 为哈希表 key 中的指定字段的整数值加上增量 increment。返回执行 HINCRBY 命令之后，哈希表中字段的值。
//Redis HINCRBY 命令用于为哈希表中的字段值加上指定增量值。
//增量也可以为负数，相当于对指定字段进行减法操作。
//如果哈希表的 key 不存在，一个新的哈希表被创建并执行 HINCRBY 命令。
//如果指定的字段不存在，那么在执行命令前，字段的值被初始化为 0 。
//对一个储存字符串值的字段执行 HINCRBY 命令将造成一个错误。
//本操作的值被限制在 64 位(bit)有符号数字表示之内。
func (c *LoggingConn) HIncrBy(ctx context.Context, key string, field string, increment int64) (int64, error) {
	return redis.Int64(c.DoWithTimeout(ctx, "HINCRBY", key, field, increment))
}

//HIncrByFloat 为哈希表 key 中的指定字段的浮点数值加上增量 increment 。
//Redis HINCRBYFLOAT 命令用于为哈希表中的字段值加上指定浮点数增量值。
//如果指定的字段不存在，那么在执行命令前，字段的值被初始化为0 。
//如果指定的字段值不是浮点数格式，那么将返回错误：ERR hash value is not a float
func (c *LoggingConn) HIncrByFloat(ctx context.Context, key string, field string, increment float64) (float64, error) {
	return redis.Float64(c.DoWithTimeout(ctx, "HINCRBYFLOAT", key, field, increment))
}

//HKeys 获取哈希表中的所有字段名称
func (c *LoggingConn) HKeys(ctx context.Context, key string) ([]string, error) {
	return redis.Strings(c.DoWithTimeout(ctx, "HKEYS", key))
}

//HValues 获取哈希表中的所有字段名称
func (c *LoggingConn) HValues(ctx context.Context, key string) ([]string, error) {
	values, err := redis.Strings(c.DoWithTimeout(ctx, "HVALS", key))
	if err == redis.ErrNil {
		return nil, nil
	}
	return values, err
}

//HScan 迭代哈希表中的键值对。返回的扫码到的字段和字段值的map数据以及最新的游标位置。
//注意 count值在Hash集合元素较少时不生效的
//redis HSCAN 命令基本语法如下：
//HSCAN key cursor [MATCH pattern] [COUNT count]
//cursor - 游标。
//pattern - 匹配的模式。
//count - 指定从数据集里返回多少元素，默认值为 10 。
func (c *LoggingConn) HScan(ctx context.Context, key string, cursor int, pattern string, count int) (fieldValues map[string]string, latestCursor int, err error) {
	var args []any
	args = append(args, key, cursor)
	if pattern != "" {
		args = append(args, "MATCH", pattern)
	}
	if count > 0 {
		args = append(args, "COUNT", count)
	}
	values, err := redis.Values(c.DoWithTimeout(ctx, "HSCAN", args...))
	if err != nil {
		return nil, 0, err
	}
	if values == nil || len(values) < 2 {
		return nil, 0, nil
	}
	latestCursor, _ = redis.Int(values[0], err)
	m, err := redis.StringMap(values[1], err)
	return m, latestCursor, err
}

//PipelineHGetAll
////key与 对应空结构一一对应
//	reply, err := redisConn.PipelineHgetAll(context.Background(), []string{"promotionv3_info_17577", "promotionv3_info_17578"}, map[string]interface{}{
//		"promotionv3_info_17577": &user{},
//		"promotionv3_info_17578": &user{},
//		"promotionv3_info_17579": &user{},
//	})
func (c *LoggingConn) PipelineHGetAll(ctx context.Context, keys []string, keyMapContainer map[string]interface{}) (map[string]interface{}, error) {
	startTime := time.Now()
	conn, err := c.Pool.GetContext(ctx)
	if err != nil {
		if err == redis.ErrPoolExhausted {
			c.logger.Errorf("PipelineHGetAll failed, %s", err.Error())
		}
		return nil, err
	}
	defer conn.Close()
	for _, key := range keys {
		_ = conn.Send("HGETALL", key)
	}
	if err := conn.Flush(); err != nil {
		return nil, err
	}

	var args []any
	containers := make(map[string]interface{}, len(keys))
	for _, key := range keys {
		args = append(args, fmt.Sprintf("HGETALL %s", key))
		values, err := redis.Values(redis.ReceiveContext(conn, ctx))
		if err != nil {
			c.logger.WithContext(ctx).Errorf("Redis PipelineHGetAll error %v", err)
			continue
		}

		if len(values) == 0 {
			c.logger.WithContext(ctx).Infof("Redis PipelineHGetAll value empty, key: %v", key)
			continue
		}
		err = redis.ScanStruct(values, keyMapContainer[key])
		if err != nil {
			c.logger.WithContext(ctx).Errorf("PipelineHGetAll.ScanStruct failed, %v", err)
			continue
		}
		containers[key] = keyMapContainer[key]
	}

	c.log(ctx, logAttr{
		method:    "PipelineHGetAll",
		startTime: startTime,
		command:   "Pipeline",
		args:      args,
		reply:     stringifyReply(containers),
		err:       err,
	})
	return containers, nil
}

func (c *LoggingConn) PipelineHGetAllMap(ctx context.Context, keys []string) (map[string]interface{}, error) {
	startTime := time.Now()
	conn, err := c.Pool.GetContext(ctx)
	if err != nil {
		if err == redis.ErrPoolExhausted {
			c.logger.Errorf("PipelineHGetAllMap failed, %s", err.Error())
		}
		return nil, err
	}
	defer conn.Close()
	for _, key := range keys {
		_ = conn.Send("HGETALL", key)
	}
	if err := conn.Flush(); err != nil {
		return nil, err
	}

	var args []any
	containers := make(map[string]interface{}, len(keys))
	for _, key := range keys {
		args = append(args, fmt.Sprintf("HGETALL %s", key))
		values, err := redis.Values(redis.ReceiveContext(conn, ctx))
		if err != nil {
			c.logger.WithContext(ctx).Errorf("Redis PipelineHGetAllMap error %v", err)
			continue
		}

		if len(values) == 0 {
			c.logger.WithContext(ctx).Infof("Redis PipelineHGetAllMap value empty, key: %v", key)
			continue
		}
		stringMap, _ := redis.StringMap(values, err)
		containers[key] = stringMap
	}

	c.log(ctx, logAttr{
		method:    "PipelineHGetAllMap",
		startTime: startTime,
		command:   "Pipeline",
		args:      args,
		reply:     stringifyReply(containers),
		err:       err,
	})
	return containers, nil
}

func (c *LoggingConn) PipelineHGetField(ctx context.Context, keyList []string, field string) (map[string]interface{}, error) {
	startTime := time.Now()
	conn, err := c.Pool.GetContext(ctx)
	if err != nil {
		if err == redis.ErrPoolExhausted {
			c.logger.Errorf("PipelineHGetAll failed, %s", err.Error())
		}
		return nil, err
	}
	defer conn.Close()
	for _, key := range keyList {
		_ = conn.Send("HGET", key, field)
	}
	if err := conn.Flush(); err != nil {
		return nil, err
	}

	var args []any
	containers := make(map[string]interface{}, len(keyList))
	for _, key := range keyList {
		args = append(args, fmt.Sprintf("HGET %s %s", key, field))
		value, err := redis.String(redis.ReceiveContext(conn, ctx))
		if err != nil {
			c.logger.WithContext(ctx).Errorf("Redis PipelineHGetAll error %v", err)
			continue
		}
		containers[key] = value
	}

	c.log(ctx, logAttr{
		method:    "PipelineHGetField",
		startTime: startTime,
		command:   "HGET",
		args:      args,
		reply:     stringifyReply(containers),
		err:       err,
	})
	return containers, nil
}

func (c *LoggingConn) PipelineHMSet(ctx context.Context, setData map[string]interface{}) error {
	startTime := time.Now()
	conn, err := c.Pool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	var args []any
	for key, val := range setData {
		args = append(args, fmt.Sprintf("HMSET %s", key))
		_ = conn.Send("HMSET", redis.Args{}.Add(key).AddFlat(val)...)
	}

	err = conn.Flush()
	c.log(ctx, logAttr{
		method:    "PipelineHMSet",
		startTime: startTime,
		command:   "Pipeline",
		args:      args,
		reply:     "",
		err:       err,
	})
	return err
}

func (c *LoggingConn) PipelineHMSetWithDuration(ctx context.Context, setData map[string]interface{}, duration time.Duration) error {
	startTime := time.Now()
	conn, err := c.Pool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	var args []any
	seconds := int(duration.Seconds())
	for key, val := range setData {
		_ = conn.Send("HMSET", redis.Args{}.Add(key).AddFlat(val)...)
		if seconds > 0 {
			_ = conn.Send("EXPIRE", redis.Args{}.Add(key).Add(seconds)...)
		}

		if seconds > 0 {
			args = append(args, fmt.Sprintf("HMSET %s EXPIRE %d", key, seconds))
		} else {
			args = append(args, fmt.Sprintf("HMSET %s", key))
		}
	}

	err = conn.Flush()
	c.log(ctx, logAttr{
		method:    "PipelineHMSetWithDuration",
		startTime: startTime,
		command:   "Pipeline",
		args:      args,
		reply:     "",
		err:       err,
	})
	return err
}

func (c *LoggingConn) PipelineHMSetWithDurations(ctx context.Context, setData map[string]interface{}, durations map[string]time.Duration) error {
	startTime := time.Now()
	conn, err := c.Pool.GetContext(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	var args []any
	for key, val := range setData {
		_ = conn.Send("HMSET", redis.Args{}.Add(key).AddFlat(val)...)
		seconds := 0
		if duration, ok := durations[key]; ok {
			seconds = int(duration.Seconds())
			if seconds > 0 {
				_ = conn.Send("EXPIRE", redis.Args{}.Add(key).Add(seconds)...)
			}
		}

		if seconds > 0 {
			args = append(args, fmt.Sprintf("HMSET %s EXPIRE %d", key, seconds))
		} else {
			args = append(args, fmt.Sprintf("HMSET %s", key))
		}
	}

	err = conn.Flush()
	c.log(ctx, logAttr{
		method:    "PipelineHMSetWithDurations",
		startTime: startTime,
		command:   "Pipeline",
		args:      args,
		reply:     "",
		err:       err,
	})
	return err
}

func (c *LoggingConn) HLen(ctx context.Context, key string) (len int, err error) {
	len, err = redis.Int(c.DoWithTimeout(ctx, "HLEN", key))
	if err != nil && !c.IsErrNil(err) {
		return 0, err
	}

	return len, nil
}

func (c *LoggingConn) PipelineHLenMap(ctx context.Context, keys []string) (map[string]int, error) {
	startTime := time.Now()
	conn, err := c.Pool.GetContext(ctx)
	if err != nil {
		if err == redis.ErrPoolExhausted {
			c.logger.Errorf("PipelineHMGetLenMap failed, %s", err.Error())
		}
		return nil, err
	}
	defer conn.Close()
	for _, key := range keys {
		_ = conn.Send("HLEN", key)
	}
	if err := conn.Flush(); err != nil {
		return nil, err
	}

	var args []any
	containers := make(map[string]int, len(keys))
	for _, key := range keys {
		args = append(args, fmt.Sprintf("HLEN %s", key))
		value, err := redis.Int(redis.ReceiveContext(conn, ctx))
		if err != nil {
			c.logger.WithContext(ctx).Errorf("Redis PipelineHMGetLenMap error %v", err)
			continue
		}

		containers[key] = value
	}

	c.log(ctx, logAttr{
		method:    "PipelineHLenMap",
		startTime: startTime,
		command:   "Pipeline",
		args:      args,
		reply:     stringifyReply(containers),
		err:       err,
	})
	return containers, nil
}
