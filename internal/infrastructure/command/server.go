package command

import (
	"context"
	"syscall"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/spf13/cobra"
)

var (
	_ transport.Server = (*Server)(nil)
)

type ServerOption func(o *Server)

type Server struct {
	log     *log.Helper
	ms      []middleware.Middleware
	rootCmd *cobra.Command
}

// Logger with server logger.
func WithLogger(logger log.Logger) ServerOption {
	return func(s *Server) {
		s.log = log.NewHelper(log.With(logger, "moudle", "cmd"))
	}
}

//with middleware
func WithMiddleware(m ...middleware.Middleware) ServerOption {
	return func(s *Server) {
		s.ms = m
	}
}

func NewServer(ops ...ServerOption) *Server {
	srv := &Server{
		log: log.NewHelper(log.GetLogger()),
	}

	for _, o := range ops {
		o(srv)
	}

	srv.rootCmd = &cobra.Command{
		Use:     "task",
		Short:   "activity cmd.",
		Long:    `activity cmd.`,
		Version: "1.0",
	}
	return srv
}

func (s *Server) Start(ctx context.Context) error {
	s.log.WithContext(ctx).Info("开始执行 " + time.Now().Format("2006-01-02 15:04:05"))
	if err := s.rootCmd.Execute(); err != nil {
		s.log.WithContext(ctx).Error("执行失败 " + err.Error())
		return err
	}
	defer syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.log.WithContext(ctx).Info("执行完成 " + time.Now().Format("2006-01-02 15:04:05"))
	return nil
}

type ICliService interface {
	CmdService(ctx context.Context, value []string) error
	GetCmd() string
	GetDesc() string
	StringVar(cmdObj *cobra.Command)
}

func (s *Server) RegisterCmdService(f ICliService) {
	cmdObj := &cobra.Command{
		Use:   f.GetCmd(),
		Short: f.GetDesc(),
		Long:  f.GetDesc(),
		FParseErrWhitelist: cobra.FParseErrWhitelist{
			UnknownFlags: true,
		},
		Run: func(cmd *cobra.Command, args []string) {
			h := func(ctx context.Context, in interface{}) (interface{}, error) {
				err := f.CmdService(ctx, args)
				return nil, err
			}
			handler := middleware.Chain(s.ms...)(h)
			tr := &Transport{
				reqHeader:   header{},
				replyHeader: header{},
			}
			ctx := transport.NewServerContext(context.Background(), tr)
			_, _ = handler(ctx, args)
		},
	}
	f.StringVar(cmdObj)
	s.rootCmd.AddCommand(cmdObj)
}
