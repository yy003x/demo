package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// Float64Map is a helper that converts an array of strings (alternating key, value)
// into a map[string]float64.
func Float64Map(result interface{}, err error) (map[string]float64, error) {
	values, err := redis.Values(result, err)
	if err != nil {
		return nil, err
	}

	if len(values)%2 != 0 {
		return nil, fmt.Errorf("redis: Float64Map expects even number of values result, got %d", len(values))
	}

	m := make(map[string]float64, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].([]byte)
		if !ok {
			return nil, fmt.Errorf("redigo: Float64Map key[%d] not a bulk string value, got %T", i, values[i])
		}

		value, err := redis.Float64(values[i+1], nil)
		if err != nil {
			return nil, err
		}

		m[string(key)] = value
	}
	return m, nil
}
