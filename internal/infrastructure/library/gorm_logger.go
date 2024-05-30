package library

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type Gorm struct {
	logger.Config
	log *log.Helper
}

func NewGorm(lg log.Logger) *Gorm {
	return &Gorm{
		log: log.NewHelper(log.With(lg, "x_module", "pkg/gorm")),
		Config: logger.Config{
			SlowThreshold: time.Millisecond * 500,
			LogLevel:      logger.Info,
			Colorful:      false,
		},
	}
}

// LogMode log mode
func (l *Gorm) LogMode(level logger.LogLevel) logger.Interface {
	newlogger := *l
	newlogger.LogLevel = level
	return &newlogger
}

// Info print info
func (l *Gorm) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		l.log.WithContext(ctx).Info(append([]interface{}{msg, utils.FileWithLineNum()}, data...)...)
	}
}

// Warn print warn messages
func (l *Gorm) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		l.log.WithContext(ctx).Warn(append([]interface{}{msg, utils.FileWithLineNum()}, data...)...)
	}
}

// Error print error messages
func (l *Gorm) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		l.log.WithContext(ctx).Error(append([]interface{}{msg, utils.FileWithLineNum()}, data...)...)
	}
}

// Trace print sql message
func (l *Gorm) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= logger.Error && !errors.Is(err, logger.ErrRecordNotFound):
		sql, rows := fc()
		if rows == -1 {
			l.log.WithContext(ctx).Errorw("x_file", utils.FileWithLineNum(), "x_error", err, "x_duration", elapsed.Seconds(), "x_rows", "-", "x_action", sql)
		} else {
			l.log.WithContext(ctx).Errorw("x_file", utils.FileWithLineNum(), "x_error", err, "x_duration", elapsed.Seconds(), "x_rows", rows, "x_action", sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL: %v", l.SlowThreshold)
		if rows == -1 {
			l.log.WithContext(ctx).Errorw("x_file", utils.FileWithLineNum(), "x_error", slowLog, "x_duration", elapsed.Seconds(), "x_rows", "-", "x_action", sql)
		} else {
			l.log.WithContext(ctx).Errorw("x_file", utils.FileWithLineNum(), "x_error", slowLog, "x_duration", elapsed.Seconds(), "x_rows", rows, "x_action", sql)
		}
	case l.LogLevel == logger.Info:
		sql, rows := fc()
		if rows == -1 {
			l.log.WithContext(ctx).Infow("x_file", utils.FileWithLineNum(), "x_duration", elapsed.Seconds(), "x_rows", "-", "x_action", sql)
		} else {
			l.log.WithContext(ctx).Infow("x_file", utils.FileWithLineNum(), "x_duration", elapsed.Seconds(), "x_rows", rows, "x_action", sql)
		}
	}
}
