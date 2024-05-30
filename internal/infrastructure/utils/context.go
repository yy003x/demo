package utils

import (
	"be_demo/internal/infrastructure/define"
	"context"

	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/metadata"
)

//不支持 协程并发
func TransferGinContext(c *gin.Context) context.Context {
	ctx := c.Request.Context()
	if md, ok := metadata.FromServerContext(ctx); ok {
		if len(md.Get(define.XTraceId)) == 0 {
			traceId := c.GetHeader(define.XTraceId)
			md.Set(define.XTraceId, c.GetHeader(traceId))
			ctx = metadata.NewServerContext(ctx, md)
		}
		return ctx
	}
	ctx = metadata.NewServerContext(ctx, metadata.New(
		map[string][]string{
			define.XTraceId: {c.GetHeader(define.XTraceId)},
		},
	),
	)
	if _, ok := kgin.FromGinContext(ctx); !ok {
		ctx = kgin.NewGinContext(ctx, c)
	}
	c.Request = c.Request.WithContext(ctx)
	return ctx
}

//协程并发 使用
func BackgroundContext(ctx context.Context) context.Context {
	if md, ok := metadata.FromServerContext(ctx); ok {
		if len(md.Get(define.XTraceId)) == 0 {
			traceId := NewUUID()
			md.Set(define.XTraceId, traceId)
			ctx = metadata.NewServerContext(ctx, md)
		}
		return ctx
	}
	ctx = metadata.NewServerContext(ctx, metadata.New(
		map[string][]string{
			define.XTraceId: {NewUUID()},
		},
	),
	)
	return ctx
}
