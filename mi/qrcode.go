package mi

import (
	"errors"
	"fmt"

	"github.com/duke-git/lancet/v2/datetime"
	"github.com/duke-git/lancet/v2/netutil"
	"github.com/google/go-querystring/query"
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
		TimeStamp:  datetime.GetNowDateTime(),
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

	client := netutil.NewHttpClient()
	resp, err := client.SendRequest(&netutil.HttpRequest{
		RawURL: fmt.Sprintf("https://openapi.alipay.com/gateway.do?%s", sign),
		Method: "GET",
	})
	if err != nil {
		return nil, err
	}

	result := Result{}
	err = client.DecodeResponse(resp, &result)
	if err != nil {
		return nil, err
	}

	response := result.AlipayOpenAppQrcodeCreateResponse
	if response.Code != "10000" {
		return nil, errors.New(response.Msg)
	}
	return response, nil
}
