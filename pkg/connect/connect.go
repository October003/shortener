package connect

import (
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// client 定义全局的HTTP客户端
var client = &http.Client{
	Transport: &http.Transport{
		DisableKeepAlives: true,
	},
	Timeout: 2 * time.Second,
}

// Get 判断url是否能请求通
func Get(url string) bool {
	resp, err := client.Get(url)
	if err != nil {
		logx.Errorw("connect cient.Get failed", logx.LogField{Key: "err", Value: err.Error()})
		return false
	}
	resp.Body.Close()
	// 别人给我发一个跳转响应 这里也不算过
	return resp.StatusCode == http.StatusOK
}
