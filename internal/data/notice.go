package data

import (
	"be_demo/internal/conf"
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

// Data .
type Notice struct {
	log    *log.Helper
	config *NoticeConfig
}

// NewData .
func NewNotice(
	bc *conf.Bootstrap,
	logger log.Logger,
) *Notice {
	conf := bc.GetPublic()
	return &Notice{
		config: &NoticeConfig{
			MonitorUrl: conf.GetNotice().GetUrl(),
			Secret:     conf.GetNotice().Secret,
		},
		log: log.NewHelper(log.With(logger, "x_module", "data/Notice")),
	}
}

type NoticeConfig struct {
	MonitorUrl string
	Secret     string
}

//知音楼sigin加密
func (notice *Notice) zhiyinlouSign(secret, timestamp string) string {
	data := timestamp + "\n" + secret
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(data))
	return url.QueryEscape(base64.StdEncoding.EncodeToString(h.Sum(nil)))
}

//发送知音楼提醒
func (notice *Notice) SendMsg(args ...string) {
	//截取错误前1000个字符进行报警输出
	s := strings.Join(args, "\n")
	if len(s) > 1000 {
		s = s[:1000]
	}
	b := json.RawMessage(`
		{"msgtype": "text","text": {"content":` + strconv.Quote(s) + `}}`)

	timeStamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	notice.config.MonitorUrl += "&timestamp=" + timeStamp + "&sign=" + notice.zhiyinlouSign(notice.config.Secret, timeStamp)
	_, _ = postJson(notice.config.MonitorUrl, b)
}

//发送知音楼提醒
func (notice *Notice) SendMonitor(args ...interface{}) {
	slice := make([]string, len(args))
	//因堆栈无法发出, 如果为错误类型则进行转换获取错误文本
	for i, v := range args {
		convertErr, ok := v.(error)
		if ok {
			slice[i] = convertErr.Error()
		} else {
			slice[i] = fmt.Sprint(v)
		}
	}

	host, _ := os.Hostname()
	//截取错误前1000个字符进行报警输出
	s := "[" + host + "]\n" + strings.Join(slice, ",")
	if len(s) > 1000 {
		s = s[:1000]
	}
	b := json.RawMessage(`
		{"msgtype": "text","text": {"content":` + strconv.Quote(s) + `}}`)

	timeStamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	notice.config.MonitorUrl += "&timestamp=" + timeStamp + "&sign=" + notice.zhiyinlouSign(notice.config.Secret, timeStamp)
	_, _ = postJson(notice.config.MonitorUrl, b)
}

// HttpPost post请求
func postJson(url string, params []byte) ([]byte, error) {
	body := bytes.NewBuffer(params)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("httpcode error:" + fmt.Sprint(resp.StatusCode))
	}
	return respData, nil
}
