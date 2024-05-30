package cron_serv

import (
	"be_demo/internal/biz/cron_biz"
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

//DemoCmdService  ./bin/command --conf=./configs/test --log=./logs test
type DemoCronService struct {
	log         *log.Helper
	desc        string
	demoCrontab *cron_biz.DemoCrontab
}

// NewGreeterService new a greeter service.
func NewDemoCronService(
	logger log.Logger,
	demoCrontab *cron_biz.DemoCrontab,
) *DemoCronService {
	return &DemoCronService{
		log:         log.NewHelper(log.With(logger, "x_module", "cron_serv/DemoLogic")),
		desc:        "test",
		demoCrontab: demoCrontab,
	}
}

func (c *DemoCronService) GetDesc() string {
	return c.desc
}

// cron处理
func (c *DemoCronService) CronNotice11(ctx context.Context, req interface{}) (interface{}, error) {
	for i := 0; i < 3; i++ {
		c.demoCrontab.Notice(ctx, "中午了,喝点水,吃点DHA,走一走")
		time.Sleep(5 * time.Second)
	}
	return nil, nil
}

// cron处理
func (c *DemoCronService) CronNotice15(ctx context.Context, req interface{}) (interface{}, error) {
	for i := 0; i < 3; i++ {
		c.demoCrontab.Notice(ctx, "下午了,喝点水,吃点铁,走一走")
		time.Sleep(5 * time.Second)
	}
	return nil, nil
}

// cron处理
func (c *DemoCronService) CronNotice19(ctx context.Context, req interface{}) (interface{}, error) {
	for i := 0; i < 3; i++ {
		c.demoCrontab.Notice(ctx, "吃饭了,吃饭不积极,脑袋有问题")
		time.Sleep(5 * time.Second)
	}
	return nil, nil
}

// cron处理
func (c *DemoCronService) CronNotice23(ctx context.Context, req interface{}) (interface{}, error) {
	for i := 0; i < 3; i++ {
		c.demoCrontab.Notice(ctx, "晚上不早了,吃点钙,休息了")
		time.Sleep(5 * time.Second)
	}
	return nil, nil
}
