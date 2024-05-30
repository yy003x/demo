package api

import (
	"be_demo/internal/service/api_serv"

	"github.com/gin-gonic/gin"
)

func LoadPingService(r *gin.RouterGroup, serv *api_serv.PingService) {
	r.GET("/ping", serv.Ping)
}
