package logger

import (
	"be_demo/internal/infrastructure/logger/stat"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var _ log.Logger = (*Logger)(nil)

type Config struct {
	Path         string
	Level        string
	Rotationtime time.Duration
	Maxage       time.Duration
	OpenStat     bool
	Stdout       bool
}

type Logger struct {
	logger *logrus.Logger
	stat   *stat.Stat
	notice INoticeService
}

func NewLogger(config Config) *Logger {
	lrus := logrus.New()
	// 设置日志级别
	if level, err := logrus.ParseLevel(config.Level); err != nil {
		lrus.SetLevel(logrus.DebugLevel)
	} else {
		lrus.SetLevel(level)
	}
	// 设置日志格式为JSON
	lrus.SetFormatter(&logFormatter{})
	// lrus.SetFormatter(&logrus.JSONFormatter{})
	//设置日志滚动更新
	writer, err := rotatelogs.New(
		config.Path,
		rotatelogs.WithRotationTime(config.Rotationtime),
		rotatelogs.WithMaxAge(config.Maxage),
	)
	if err != nil {
		lrus.Fatalf("Failed to initialize log file rotatelogs: %v", err)
	}
	if config.Stdout {
		lrus.SetOutput(io.MultiWriter(os.Stdout, writer))
	} else {
		lrus.SetOutput(writer)
	}
	l := &Logger{
		logger: lrus,
	}
	//设置kratos底层 msg 字段名
	log.DefaultMessageKey = "x_msg"
	return l
}

func (lf *Logger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		lf.errorf("params error")
		return nil
	}

	buf := make(map[string]interface{})
	for i := 0; i < len(keyvals); i += 2 {
		if logKey := keyvals[i].(string); logKey != "" {
			buf[logKey] = keyvals[i+1]
		}
	}

	switch level {
	case log.LevelFatal:
		lf.fatalM(buf)
	case log.LevelWarn:
		lf.warnM(buf)
	case log.LevelError:
		lf.errorM(buf)
	case log.LevelDebug:
		lf.debugM(buf)
	default:
		lf.infoM(buf)
	}

	//开启统计
	if lf.stat != nil {
		lf.stat.Push(buf)
	}

	return nil
}

//log 服务统计功能
func (lf *Logger) logStat() {
	lf.stat = stat.NewStat(&stat.Config{
		SlowMsThreshold:       time.Millisecond * 800,
		DbSlowMsThreshold:     time.Millisecond * 100,
		RedisSlowMsThreshold:  time.Millisecond * 100,
		ClientSlowMsThreshold: time.Millisecond * 500,

		WindowSize:     10,
		WindowDuration: time.Second * 10,
	})
	lf.stat.Consumer()

	go func() {
		ticker := time.NewTicker(time.Second * 10)
		defer func() {
			ticker.Stop()
			if err := recover(); err != nil {
				lf.errorf("stat err: %+v", err)
			}
		}()
		for range ticker.C {
			summary := lf.stat.GetNoticeSummary()
			if summary != "" {
				if lf.notice != nil {
					lf.notice.SendMonitor(summary)
				}
			}
		}
	}()
}

func (l *Logger) errorf(format string, args ...interface{}) {
	l.errorM(map[string]interface{}{
		"x_msg": fmt.Sprintf(format, args...),
	})
}

func (l *Logger) debugM(msgs map[string]interface{}) {
	l.logger.WithFields(msgs).Debug()
}

func (l *Logger) infoM(msgs map[string]interface{}) {
	l.logger.WithFields(msgs).Info()
}

func (l *Logger) warnM(msgs map[string]interface{}) {
	l.logger.WithFields(msgs).Warn()
}

func (l *Logger) errorM(msgs map[string]interface{}) {
	msgs["x_extra"] = map[string]interface{}{"stack": callFrames(15)}
	l.logger.WithFields(msgs).Error()
}

func (l *Logger) fatalM(msgs map[string]interface{}) {
	msgs["x_extra"] = map[string]interface{}{"stack": callFrames(15)}
	l.logger.WithFields(msgs).Fatal()
}

var packageName string
var once sync.Once

func callFrames(maxDept int) []string {
	var stacks []string
	pcs := make([]uintptr, maxDept)
	depth := runtime.Callers(3, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	// cache this package's fully-qualified name
	once.Do(func() {
		packageName = getPackageName(runtime.FuncForPC(pcs[0]).Name())
	})
	for f, again := frames.Next(); again; f, again = frames.Next() {
		pkg := getPackageName(f.Function)

		// If the caller isn't part of this package, we're done
		if pkg != packageName {
			stacks = append(stacks, fmt.Sprintf("%s:%d", f.File, f.Line))
		}
	}
	return stacks
}

// getPackageName reduces a fully qualified function name to the package name
// There really ought to be to be a better way...
func getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
}
