package cli_biz

import (
	"be_demo/internal/data"
	"be_demo/internal/infrastructure/times"
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type DemoCliLogic struct {
	log    *log.Helper
	notice *data.Notice
}

// NewGreeterUsecase new a Greeter usecase.
func NewDemoCliLogic(
	logger log.Logger,
	notice *data.Notice,
) *DemoCliLogic {
	return &DemoCliLogic{
		log:    log.NewHelper(log.With(logger, "x_module", "cli_biz/DemoLogic")),
		notice: notice,
	}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *DemoCliLogic) Test(ctx context.Context, in interface{}) (interface{}, error) {
	uc.log.WithContext(ctx).Infof("Ping: %v", in)
	dt := times.Time(time.Now()).String()
	uc.notice.SendMsg("[❤️提醒❤️]", "现在时间:"+dt, "吃饭了")
	return nil, nil
}
