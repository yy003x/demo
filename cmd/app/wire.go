//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"be_demo/internal/biz/app_biz"
	"be_demo/internal/conf"
	"be_demo/internal/repository"
	"be_demo/internal/server/app"
	"be_demo/internal/service/app_serv"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		app.ProviderSet,
		app_serv.ProviderSet,
		app_biz.ProviderSet,
		repository.ProviderSet,
		newApp,
	))
}
