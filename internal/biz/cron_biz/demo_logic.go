package cron_biz

import (
	"be_demo/internal/data"
	"be_demo/internal/infrastructure/times"
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type DemoCrontab struct {
	log    *log.Helper
	notice *data.Notice
}

// NewGreeterUsecase new a Greeter usecase.
func NewDemoCrontab(
	logger log.Logger,
	notice *data.Notice,
) *DemoCrontab {
	return &DemoCrontab{
		log:    log.NewHelper(log.With(logger, "x_module", "cli_biz/DemoLogic")),
		notice: notice,
	}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *DemoCrontab) Notice(ctx context.Context, msg string) (interface{}, error) {
	dt := times.Time(time.Now()).String()
	uc.notice.SendMsg("[❤️提醒❤️]", "现在时间:"+dt, msg)
	uc.log.WithContext(ctx).Info(msg)
	return nil, nil
}
