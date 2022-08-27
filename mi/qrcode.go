package mi

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/go-querystring/query"
	"github.com/tiantour/fetch"
	"github.com/tiantour/tempo"
)

// QR qr
type QR struct {
	URLParam   string `json:"url_param,omitempty"`   // 小程序中能访问到的页面路径。
	QueryParam string `json:"query_param,omitempty"` // 小程序的启动参数，打开小程序的query，在小程序onLaunch的方法中获取。
	Describe   int    `json:"describe,omitempty"`    // 对应的二维码描述。
}

// Generate qrcode generate
func (q *QR) Generate(content string) (*Response, error) {
	args := &Request{
		AppID:      AppID,
		Method:     "alipay.open.app.qrcode.create",
		Format:     "JSON",
		Charset:    "utf-8",
		SignType:   "RSA2",
		TimeStamp:  tempo.NewNow().String(),
		Version:    "1.0",
		BizContent: content,
	}

	tmp, err := query.Values(args)
	if err != nil {
		return nil, err
	}

	sign, err := NewToken().Sign(&tmp, PrivatePath)
	if err != nil {
		return nil, err
	}

	body, err := fetch.Cmd(&fetch.Request{
		Method: "GET",
		URL: fmt.Sprintf("https://openapi.alipay.com/gateway.do?%s",
			sign,
		),
	})
	if err != nil {
		return nil, err
	}

	result := Result{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	response := result.AlipayOpenAppQrcodeCreateResponse
	if response.Code != "10000" {
		return nil, errors.New(response.Msg)
	}
	return response, nil
}
