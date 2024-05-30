package main

import (
	"be_demo/internal/conf"
	"be_demo/internal/infrastructure/crontab"
	"be_demo/internal/infrastructure/logger"
	"flag"
	"os"
	"path/filepath"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = "api_server"
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string
	// flaglog is the config flag.
	flaglog string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf=./configs/test")
	flag.StringVar(&flaglog, "log", "../../logs", "logs path, eg: -log=./logs")
}

func newApp(logger log.Logger, cs *crontab.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			cs,
		),
	)
}

func main() {
	flag.Parse()

	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	logConf := bc.GetPublic().GetLog()
	logger := log.With(
		logger.NewLogger(logger.Config{
			Path:         filepath.Join(flaglog, logConf.GetPath()),
			Level:        logConf.GetLevel(),
			Rotationtime: logConf.Rotationtime.AsDuration(),
			Maxage:       logConf.Maxage.AsDuration(),
			Stdout:       logConf.GetStdout(),
		}),
		"x_caller", log.DefaultCaller,
		"x_timestamp", logger.Timestamp(),
		"x_date", logger.Date(logger.TimestampStdFormat),
		"x_service_id", id,
		"x_service_name", Name,
		"x_service_version", Version,
		"x_trace_id", tracing.TraceID(),
		"x_span_id", tracing.SpanID(),
	)

	app, cleanup, err := wireApp(&bc, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
