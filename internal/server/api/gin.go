package api

import (
	"be_demo/internal/conf"
	"be_demo/internal/infrastructure/middleware"
	"be_demo/internal/service/api_serv"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	kGin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewGinHttpServer(
	conf *conf.Bootstrap,
	logger log.Logger,
	pingService *api_serv.PingService,
	activityService *api_serv.ActivityService,
) *http.Server {
	var opts = []http.ServerOption{}
	if conf.Server.Api.Network != "" {
		opts = append(opts, http.Network(conf.Server.Api.Network))
	}
	if conf.Server.Api.Addr != "" {
		opts = append(opts, http.Address(conf.Server.Api.Addr))
	}
	if conf.Server.Api.Timeout != nil {
		opts = append(opts, http.Timeout(conf.Server.Api.Timeout.AsDuration()))
	}
	// 设置运行模式
	gin.SetMode(gin.ReleaseMode)
	// 禁用默认的日志输出
	gin.DefaultWriter = ioutil.Discard
	router := gin.Default()

	router.Use(kGin.Middlewares(
		metadata.Server(metadata.WithPropagatedPrefix("x-")),
		recovery.Recovery(),
	))
	noneGroup := router.Group("/", middleware.RecoveryMiddleware(logger))
	// signGroup := router.Group("/", middleware.RecoveryMiddleware(logger), middleware.AuthSignMiddleware(nconf, logger))
	signGroup := router.Group("/", middleware.RecoveryMiddleware(logger))
	//注册优化
	LoadPingService(noneGroup, pingService)
	LoadActivityService(noneGroup, signGroup, activityService)
	router.NoRoute(pingService.CustomNotFound)
	srv := http.NewServer(opts...)
	srv.HandlePrefix("/", router)
	return srv
}
