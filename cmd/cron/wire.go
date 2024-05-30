//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"be_demo/internal/biz/cron_biz"
	"be_demo/internal/conf"
	"be_demo/internal/data"
	"be_demo/internal/server/cron"
	"be_demo/internal/service/cron_serv"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Bootstrap, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		cron.ProviderSet,
		cron_serv.ProviderSet,
		cron_biz.ProviderSet,
		data.ProviderSet,
		newApp,
	))
}
