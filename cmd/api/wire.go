//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"be_demo/internal/biz/api_biz"
	"be_demo/internal/conf"
	"be_demo/internal/data"
	"be_demo/internal/infrastructure/nacosx"
	"be_demo/internal/repository"
	"be_demo/internal/server/api"
	"be_demo/internal/service/api_serv"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Bootstrap, *nacosx.NacosConf[conf.NacosConfig], log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		api.ProviderSet,
		api_serv.ProviderSet,
		data.ProviderSet,
		api_biz.ProviderSet,
		repository.ProviderSet,
		newApp,
	))
}
