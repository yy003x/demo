package define

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
)

type ExtError struct {
	code   int
	reason string
	msg    string
}

func (e *ExtError) Code() int {
	if e == nil {
		return 0
	}
	return e.code
}

func (e *ExtError) Reason() string {
	if e == nil {
		return ""
	}
	return e.reason
}

func (e *ExtError) Msg(format string, args ...interface{}) string {
	if e == nil {
		return "nil"
	}
	return e.msg + "[" + fmt.Sprintf(format, args...) + "]"
}

func (e ExtError) New(err error) error {
	if err != nil {
		return errors.New(e.Code(), e.Reason(), err.Error())
	}
	return errors.New(e.Code(), e.Reason(), "nil err")
}

func (e ExtError) Newf(format string, a ...interface{}) error {
	return errors.Newf(e.Code(), e.Reason(), e.Msg(format, a...))
}

var (
	/***** 客户端错误 *****/
	ExtError_400000 = ExtError{400000, "请登录", ""}
	ExtError_400001 = ExtError{400001, "请求数据有问题,请检查后重试", "request data error"}
	ExtError_400002 = ExtError{400002, "请求签名有问题,请检查后重试", "middleware sign error"}
	ExtError_400003 = ExtError{400003, "请求参数有问题,请检查后重试", "service argument error"}
	ExtError_400004 = ExtError{400004, "请求的数据重复,请重新尝试", "repeat error"}
	ExtError_400005 = ExtError{400005, "请求频次太快了,请稍等再试", "limit error"}

	/***** 服务端错误 *****/
	ExtError_500000 = ExtError{500000, "服务遇到了问题，请稍后再试", "system panic error"}
	ExtError_500001 = ExtError{500001, "数据获取有问题，请稍后再试", "service call error"}
	ExtError_510001 = ExtError{510001, "服务遇到了问题，请稍后再试", "mysql error"}
	ExtError_520001 = ExtError{520001, "服务遇到了问题，请稍后再试", "redis error"}

	/***** 业务错误 *****/
	//立减活动 300000 -- 399999
	ErrCode_300000 = ExtError{300000, "抱歉...", "zero_token error"}
	ErrCode_399999 = ExtError{399999, "抱歉...", "qual error"}
)
