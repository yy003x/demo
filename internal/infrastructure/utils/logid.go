// GenerateLogID 生成唯一的日志标识符（Log ID）
package utils

import (
	"fmt"
	"math"
	"time"
)

func GenerateLogID() int {
	uniPartID := int(math.Abs(float64(crc32(uniqid(fmt.Sprintf("%d", time.Now().UnixNano()), true)))))
	year := time.Now().Year()%9 + 1
	dayOfYear := time.Now().YearDay()
	secondOfDay := time.Now().Hour()*3600 + time.Now().Minute()*60 + time.Now().Second()
	micro := time.Now().Nanosecond() / 1000
	// 将 uniPartID 整合到 logID 中
	logID := year*10000000 + dayOfYear*10000 + secondOfDay*100 + micro + uniPartID
	return logID
}

// 模拟 PHP 的 crc32 函数
func crc32(s string) uint32 {
	var crc uint32 = 0xFFFFFFFF

	for _, char := range s {
		crc ^= uint32(char)
		for j := 0; j < 8; j++ {
			mask := uint32(-(int(crc) & 1))
			crc = (crc >> 1) ^ (0xEDB88320 & mask)
		}
	}

	return ^crc
}

// 模拟 PHP 的 uniqid 函数
func uniqid(prefix string, moreEntropy bool) string {
	now := time.Now()
	sec := now.Unix()
	usec := now.Nanosecond() / 1000
	return fmt.Sprintf("%s%08x%05x", prefix, sec, usec)
}
