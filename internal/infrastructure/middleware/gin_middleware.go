package middleware

import (
	"be_demo/internal/data"
	"be_demo/internal/infrastructure/define"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

func RecoveryMiddleware(logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 处理 panic，例如记录日志
				buf := make([]byte, 64<<10) //nolint:gomnd
				n := runtime.Stack(buf, false)
				buf = buf[:n]
				log.WithContext(c, logger).Log(log.LevelError,
					"x_panic", err,
					"x_info", string(buf),
				)
				data.NewStdOut(logger).ApiStdOut(c, nil, define.ExtError_500000.Newf("panic"))
			}
		}()
		// 执行下一个中间件或处理函数
		c.Next()
	}
}
