package api_serv

import (
	"be_demo/internal/biz/api_biz"
	"be_demo/internal/data"
	"be_demo/internal/entity/api_vo"
	"be_demo/internal/infrastructure/define"
	"be_demo/internal/infrastructure/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

// GreeterService is a greeter service.
type ActivityService struct {
	log           *log.Helper
	stdout        *data.StdOut
	activityLogic *api_biz.ActivityLogic
}

// NewGreeterService new a greeter service.
func NewActivityService(
	logger log.Logger,
	stdout *data.StdOut,
	activityLogic *api_biz.ActivityLogic,
) *ActivityService {
	return &ActivityService{
		log:           log.NewHelper(log.With(logger, "x_module", "api_serv/ActivityService")),
		stdout:        stdout,
		activityLogic: activityLogic,
	}
}

// ...
func (s *ActivityService) AddActivity(c *gin.Context) {
	var (
		in  = api_vo.AddActivityRequest{}
		ctx = utils.TransferGinContext(c)
	)
	if err := c.ShouldBindJSON(&in); err != nil {
		s.stdout.ApiStdOut(c, nil, define.ExtError_400001.New(err))
		return
	}
	out, err := s.activityLogic.AddActivity(ctx, &in)
	s.stdout.ApiStdOut(c, out, err)
}

// ...
func (s *ActivityService) EditActivity(c *gin.Context) {
	var (
		in  = api_vo.EditActivityRequest{}
		ctx = utils.TransferGinContext(c)
	)
	if err := c.ShouldBindJSON(&in); err != nil {
		s.stdout.ApiStdOut(c, nil, err)
		return
	}
	out, err := s.activityLogic.EditActivity(ctx, &in)
	s.stdout.ApiStdOut(c, out, err)
}

// ...
func (s *ActivityService) RemoveActivity(c *gin.Context) {
	var (
		in  = api_vo.AddActivityRequest{}
		ctx = utils.TransferGinContext(c)
	)
	if err := c.ShouldBindJSON(&in); err != nil {
		s.stdout.ApiStdOut(c, nil, err)
		return
	}
	out, err := s.activityLogic.RemoveActivity(ctx, &in)
	s.stdout.ApiStdOut(c, out, err)
}

// ...
func (s *ActivityService) ListActivity(c *gin.Context) {
	var (
		in  = api_vo.AddActivityRequest{}
		ctx = utils.TransferGinContext(c)
	)
	if err := c.ShouldBindJSON(&in); err != nil {
		s.stdout.ApiStdOut(c, nil, err)
		return
	}
	out, err := s.activityLogic.ListActivity(ctx, &in)
	s.stdout.ApiStdOut(c, out, err)
}

// ...
func (s *ActivityService) DetailActivity(c *gin.Context) {
	var (
		in  = api_vo.AddActivityRequest{}
		ctx = utils.TransferGinContext(c)
	)
	if err := c.ShouldBindJSON(&in); err != nil {
		s.stdout.ApiStdOut(c, nil, err)
		return
	}
	out, err := s.activityLogic.DetailActivity(ctx, &in)
	s.stdout.ApiStdOut(c, out, err)
}
