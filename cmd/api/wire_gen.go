// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"be_demo/internal/biz/api_biz"
	"be_demo/internal/conf"
	"be_demo/internal/data"
	"be_demo/internal/infrastructure/nacosx"
	"be_demo/internal/repository/api_repo"
	"be_demo/internal/server/api"
	"be_demo/internal/service/api_serv"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(bootstrap *conf.Bootstrap, nacosConf *nacosx.NacosConf[conf.NacosConfig], logger log.Logger) (*kratos.App, func(), error) {
	pingLogic := api_biz.NewPingLogic(logger)
	pingService := api_serv.NewPingService(logger, pingLogic)
	stdOut := data.NewStdOut(logger)
	dbs, cleanup, err := data.NewMysqlDBS(bootstrap, logger)
	if err != nil {
		return nil, nil, err
	}
	dmActivityRepo := api_repo.NewDmActivityRepo(logger, dbs)
	activityLogic := api_biz.NewActivityLogic(logger, dmActivityRepo)
	activityService := api_serv.NewActivityService(logger, stdOut, activityLogic)
	server := api.NewGinHttpServer(bootstrap, logger, nacosConf, pingService, activityService)
	app := newApp(logger, server)
	return app, func() {
		cleanup()
	}, nil
}