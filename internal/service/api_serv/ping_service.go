package api_serv

import (
	"be_demo/internal/biz/api_biz"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

// GreeterService is a greeter service.
type PingService struct {
	log *log.Helper
	uc  *api_biz.PingLogic
}

// NewGreeterService new a greeter service.
func NewPingService(
	logger log.Logger,
	uc *api_biz.PingLogic,
) *PingService {
	return &PingService{
		log: log.NewHelper(log.With(logger, "x_module", "api_serv/PingService")),
		uc:  uc,
	}
}

// ...
func (s *PingService) Ping(c *gin.Context) {
	s.log.WithContext(c).Info("test")
	c.JSON(http.StatusOK, "ok...")
}

// 自定义404错误处理器
func (sv *PingService) CustomNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, "error")
}
