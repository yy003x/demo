//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"be_demo/internal/biz/cli_biz"
	"be_demo/internal/conf"
	"be_demo/internal/data"
	"be_demo/internal/server/cli"
	"be_demo/internal/service/cli_serv"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Bootstrap, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		cli.ProviderSet,
		cli_serv.ProviderSet,
		cli_biz.ProviderSet,
		data.ProviderSet,
		newApp,
	))
}
