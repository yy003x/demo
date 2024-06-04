package redis

import (
	"context"
	"time"
)

// 计算ctx剩余时间
func ShrinkDuration(ctx context.Context, maxTimeout time.Duration) time.Duration {
	if ctx == nil {
		return maxTimeout
	}
	var timeoutTime = time.Now().Add(maxTimeout)
	if deadline, ok := ctx.Deadline(); ok && timeoutTime.After(deadline) {
		return deadline.Sub(time.Now())
	}
	return maxTimeout
}
