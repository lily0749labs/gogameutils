// Package httpUtil 提供游戏服务常用的 HTTP 调用工具。
package httpUtil

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// JSON 使用 http.DefaultClient 发送 JSON 请求并返回完整响应体。
// 此函数保留迁移前行为：方法名转为大写，不依据 HTTP 状态码返回错误。
func JSON(method, url string, data any, headers map[string]string) ([]byte, error) {
	return JSONWithClient(http.DefaultClient, method, url, data, headers)
}

// JSONWithClient 使用指定客户端发送 JSON 请求并返回完整响应体。
func JSONWithClient(client *http.Client, method, url string, data any, headers map[string]string) ([]byte, error) {
	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(strings.ToUpper(method), url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	for key, value := range headers {
		request.Header.Set(key, value)
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer func() { _ = response.Body.Close() }()

	return io.ReadAll(response.Body)
}
