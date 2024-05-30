package logger

import (
	"context"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
)

const TimestampStdFormat = "2006-01-02 15:04:05"

// Caller returns returns a Valuer that returns a pkg/file:line description of the caller.
func Caller(depth int) log.Valuer {
	return func(context.Context) interface{} {
		_, file, line, _ := runtime.Caller(depth)
		idx := strings.LastIndexByte(file, '/')
		if idx == -1 {
			return file[idx+1:] + ":" + strconv.Itoa(line)
		}
		idx = strings.LastIndexByte(file[:idx], '/')
		return file[idx+1:] + ":" + strconv.Itoa(line)
	}
}

// Timestamp returns a timestamp Valuer with a custom time format.
func Date(layout string) log.Valuer {
	return func(context.Context) interface{} {
		return time.Now().In(time.Local).Format(layout)
	}
}

// Timestamp returns a timestamp Valuer with a custom time format.
func Timestamp() log.Valuer {
	return func(context.Context) interface{} {
		return time.Now().Unix()
	}
}

//TraceID 获取traceid
func TraceID(traceKey string) log.Valuer {
	return func(ctx context.Context) interface{} {
		if ctx == nil {
			return ""
		}
		if md, ok := metadata.FromServerContext(ctx); ok {
			extra := md.Clone().Get(traceKey)
			return extra
		}
		return ""
	}
}

//获取traceid
func IncrRpcId(rpcidKey string) log.Valuer {
	return func(ctx context.Context) interface{} {
		if ctx == nil {
			return ""
		}
		if md, ok := metadata.FromServerContext(ctx); ok {
			extra := md.Get(rpcidKey)
			last := strings.LastIndex(extra, ".")
			i, _ := strconv.Atoi(extra[last+1:])
			extra = extra[:last+1] + strconv.Itoa(i+1)
			md.Set(rpcidKey, extra)
			return extra
		}
		return ""
	}
}
