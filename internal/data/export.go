package data

import (
	"be_demo/internal/infrastructure/define"
	"be_demo/internal/infrastructure/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

type StdOut struct {
	log *log.Helper
}

func NewStdOut(
	logger log.Logger,
) *StdOut {
	return &StdOut{
		log: log.NewHelper(log.With(logger, "x_module", "data/StdOut")),
	}
}

func (e *StdOut) ApiStdOut(c *gin.Context, data interface{}, err error) {
	ctx := utils.TransferGinContext(c)
	//默认值
	resp := define.StdResponse{}
	resp.Msg = define.SuccessMsg
	resp.Time = time.Now().UnixNano() / 1e6
	resp.Trace = c.GetHeader(define.XTraceId)
	resp.Data = data
	if err != nil {
		se := errors.FromError(err)
		resp.Code = se.Code
		resp.Msg = se.Reason
		e.log.WithContext(ctx).Errorf("ApiStdOut %+v", err)
	}
	c.JSON(http.StatusOK, resp)
}
