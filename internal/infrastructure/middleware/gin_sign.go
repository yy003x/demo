package middleware

import (
	"be_demo/internal/conf"
	"be_demo/internal/data"
	"be_demo/internal/infrastructure/define"
	"be_demo/internal/infrastructure/nacosx"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

func AuthSignMiddleware(nconf *nacosx.NacosConf[conf.NacosConfig], logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		conf := nconf.Get().Root.GetSign()
		skip := conf.GetSkip()
		if !skip {
			err := VerifySign(c, conf.GetKeys())
			if err != nil {
				data.NewStdOut(logger).ApiStdOut(c, nil, err)
				c.Abort()
			}
		}
		c.Next()
	}
}

func VerifySign(c *gin.Context, keys map[string]string) error {
	if keys == nil {
		return define.ExtError_400002.Newf("签名配置为空 keys=%+v", keys)
	}
	appId := c.GetHeader(define.XAppid)
	appSecret, ok := keys[appId]
	if !ok {
		return define.ExtError_400002.Newf("不存在来源 appid=%s", appId)
	}
	ts := c.GetHeader(define.XTime)
	now := time.Now().Unix()
	rts, _ := strconv.ParseInt(ts, 10, 64)
	if diff := now - rts; diff > 86400 {
		return define.ExtError_400002.Newf("签名超时 app_id=%s ts=%s diff=%d", appId, ts, diff)
	}
	w := md5.New()
	w.Write([]byte(appId))
	w.Write([]byte("&"))
	w.Write([]byte(ts))
	w.Write([]byte(appSecret))
	genSign := hex.EncodeToString(w.Sum(nil))
	sign := c.GetHeader(define.XSign)
	if len(sign) == 0 {
		return define.ExtError_400002.Newf("签名为空 app_id=%s ts=%s sign=%s", appId, ts, sign)
	}
	if sign != genSign {
		return define.ExtError_400002.Newf("签名错误 app_id=%s ts=%s sign=%s gen_sign=%s", appId, ts, sign, genSign)
	}
	return nil
}
