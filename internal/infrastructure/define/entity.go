package define

type StdResponse struct {
	Code  int32       `json:"code"`  //错误码
	Msg   string      `json:"msg"`   //错误信息
	Time  int64       `json:"time"`  //时间戳
	Trace string      `json:"trace"` //链路追踪
	Data  interface{} `json:"data"`  //业务数据
}

const (
	SuccessMsg = "SUCCESS"
	XTraceId   = "x-traceid"
	XRpcId     = "x-rpcid"
	XAppid     = "x-appid"
	XTime      = "x-time"
	XSign      = "x-sign"
)
