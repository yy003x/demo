package times

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

// 时间格式常量
const (
	FmtDateRaw  = "2006-01-02T15:04:05.999999999Z07:00"
	FmtDateTime = "2006-01-02 15:04:05"
	FmtDateStd  = "2006-01-02"
	FmtDateStr  = "20060102"
)

// 自定义时间类型 Time
type Time time.Time

var (
	// ErrYearOutOfRange 表示年份超出范围的错误
	ErrYearOutOfRange = errors.New("year outside of range [0, 9999]")
)

// MarshalJSON 实现了 json.Marshaler 接口，将时间格式化为 JSON 字符串
func (t Time) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)
	if y := tt.Year(); y < 0 || y >= 10000 {
		return nil, ErrYearOutOfRange
	}

	b := make([]byte, 0, len(FmtDateTime)+2)
	b = append(b, '"')
	b = tt.AppendFormat(b, FmtDateTime)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON 实现了 json.Unmarshaler 接口，将 JSON 字符串解析为时间
func (t *Time) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	parsed, err := time.ParseInLocation(`"`+FmtDateTime+`"`, string(data), time.Local)
	*t = Time(parsed)
	return err
}

// String 方法将时间转换为字符串
func (t Time) String() string {
	return time.Time(t).Format(FmtDateTime)
}

// Value 实现了 database/sql/driver.Valuer 接口，将时间转换为数据库值
func (t Time) Value() (driver.Value, error) {
	return time.Time(t), nil
}

// Scan 实现了 database/sql.Scanner 接口，将数据库值转换为时间
func (t *Time) Scan(v interface{}) error {
	if vt, ok := v.(time.Time); ok {
		*t = Time(vt)
		return nil
	}
	return errors.New("时间类型错误")
}

// UnixTimestamp 返回时间的Unix时间戳（以秒为单位）
func (t Time) UnixTimestamp() int64 {
	x := time.Time(t).Unix()
	if x < 0 {
		return 0
	}
	return x
}

// RedisScan 用于从 Redis 中扫描时间值
func (t *Time) RedisScan(x interface{}) error {
	bs, ok := x.([]byte)
	if !ok {
		return fmt.Errorf("expected []byte, got %T", x)
	}
	tt, err := time.Parse(FmtDateTime, string(bs))
	if err != nil {
		return err
	}
	*t = Time(tt)
	return nil
}

func ToString(t time.Time) string {
	if !t.IsZero() {
		return time.Time(t).Format(FmtDateTime)
	}
	return ""
}
