package crontab

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/robfig/cron/v3"
)

var (
	_ transport.Server = (*Server)(nil)
)

type ServerOption func(o *Server)

type Server struct {
	log  *log.Helper
	ms   []middleware.Middleware
	cron *cron.Cron
}

// WithLogger Logger with server logger.
func WithLogger(logger log.Logger) ServerOption {
	return func(s *Server) {
		s.log = log.NewHelper(log.With(logger, "moudle", "crontab/server"))
	}
}

// WithMiddleware with middleware
func WithMiddleware(m ...middleware.Middleware) ServerOption {
	return func(s *Server) {
		s.ms = m
	}
}

// NewServer create cron server
func NewServer(ops ...ServerOption) *Server {
	logger := log.NewHelper(log.GetLogger())
	cronOps := []cron.Option{
		cron.WithSeconds(),
		cron.WithLocation(time.Local),
		cron.WithChain(jobSkipIfStillRunning(logger)),
	}
	srv := &Server{
		cron: cron.New(cronOps...),
		log:  logger,
	}
	for _, o := range ops {
		o(srv)
	}
	return srv
}

// Start 启动cron
func (s *Server) Start(c context.Context) error {
	s.log.WithContext(c).Info("cron start begin...")
	s.cron.Start()
	s.log.WithContext(c).Info("cron start")
	return nil
}

// Stop 结束cron
func (s *Server) Stop(c context.Context) error {
	s.log.WithContext(c).Infof("cron stop begin...")
	ctx := s.cron.Stop()
	<-ctx.Done()
	s.log.WithContext(c).Infof("cron stop [%v]", ctx.Err())
	return nil
}

type CronCall func(context.Context, interface{}) (interface{}, error)

func (s *Server) RegisterCron(spec string, call CronCall) error {
	callFunc := func() {
		tr := &Transport{
			reqHeader:   header{},
			replyHeader: header{},
		}
		ctx := transport.NewServerContext(context.Background(), tr)
		handler := middleware.Chain(s.ms...)(middleware.Handler(call))
		handler(ctx, "")
	}
	_, err := s.cron.AddFunc(spec, callFunc)
	return err
}

// jobSkipIfStillRunning 任务维度 防止同时执行
func jobSkipIfStillRunning(logger *log.Helper) cron.JobWrapper {
	return func(j cron.Job) cron.Job {
		var ch = make(chan struct{}, 1)
		ch <- struct{}{}
		return cron.FuncJob(func() {
			select {
			case v := <-ch:
				j.Run()
				ch <- v
			default:
				logger.Info("skip")
			}
		})
	}
}
