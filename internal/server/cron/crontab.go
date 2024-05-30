package cron

import (
	"be_demo/internal/infrastructure/crontab"
	"be_demo/internal/service/cron_serv"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
)

func NewCron(
	demoCronService *cron_serv.DemoCronService,
) *crontab.Server {
	opts := []crontab.ServerOption{
		crontab.WithMiddleware(
			metadata.Server(metadata.WithPropagatedPrefix("x-")),
		),
	}
	srv := crontab.NewServer(opts...)
	// https://pkg.go.dev/github.com/robfig/cron#section-readme cron表达式
	if err := srv.RegisterCron("0 0 11 * * *", demoCronService.CronNotice11); err != nil {
		log.Errorf("cron.RegisterCron for sms_order_completed_7days err:%v", err)
		panic(err)
	}
	if err := srv.RegisterCron("0 0 15 * * *", demoCronService.CronNotice15); err != nil {
		log.Errorf("cron.RegisterCron for sms_order_completed_7days err:%v", err)
		panic(err)
	}
	if err := srv.RegisterCron("0 0 19 * * *", demoCronService.CronNotice19); err != nil {
		log.Errorf("cron.RegisterCron for sms_order_completed_7days err:%v", err)
		panic(err)
	}
	if err := srv.RegisterCron("0 0 22 * * *", demoCronService.CronNotice23); err != nil {
		log.Errorf("cron.RegisterCron for sms_order_completed_7days err:%v", err)
		panic(err)
	}
	return srv
}
