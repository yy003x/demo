package logger

import (
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
)

var _ INoticeService = (*zhiyinlouNotice)(nil)

type INoticeService interface {
	SendMonitor(args ...interface{})
}

type zhiyinlouNotice struct {
	config *ZhiyinlouConfig
}

type ZhiyinlouConfig struct {
	MonitorUrl string
	Secret     string
}

func NewZhiyinlouNotice(config *ZhiyinlouConfig) INoticeService {
	return &zhiyinlouNotice{config: config}
}

//知音楼sigin加密
func (notice *zhiyinlouNotice) zhiyinlouSign(secret, timestamp string) string {
	data := timestamp + "\n" + secret
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(data))
	return url.QueryEscape(base64.StdEncoding.EncodeToString(h.Sum(nil)))
}

//发送知音楼提醒
func (notice *zhiyinlouNotice) SendMonitor(args ...interface{}) {
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
	b := json.RawMessage(`{"msgtype": "text","text": {"content":` + strconv.Quote(s) + `}}`)
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
