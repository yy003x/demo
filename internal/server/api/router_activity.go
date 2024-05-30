package api

import (
	"be_demo/internal/service/api_serv"

	"github.com/gin-gonic/gin"
)

func LoadActivityService(noneGroup *gin.RouterGroup, signGroup *gin.RouterGroup, serv *api_serv.ActivityService) {
	{
		noneGroup.POST("/api/activity/add", serv.AddActivity)
		noneGroup.POST("/api/activity/edit", serv.EditActivity)
		noneGroup.POST("/api/activity/list", serv.ListActivity)
		noneGroup.POST("/api/activity/detail", serv.DetailActivity)
		noneGroup.POST("/api/activity/remove", serv.RemoveActivity)
	}
}
