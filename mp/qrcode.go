package mp

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/duke-git/lancet/v2/netutil"
)

// QR qr
type QR struct {
	Scene     string                 `json:"scene"`      // 场景
	Page      string                 `json:"page"`       // 页面
	Width     int                    `json:"width"`      // 宽度
	AutoColor bool                   `json:"auto_color"` // 默认颜色
	LineColor map[string]interface{} `json:"line_color"` // 线条颜色
}

// NewQR new qr
func NewQR() *QR {
	return &QR{}
}

// Generate qr generate
func (q *QR) Generate(args *QR) ([]byte, error) {
	token, err := NewToken().Access()
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	client := netutil.NewHttpClient()
	resp, err := client.SendRequest(&netutil.HttpRequest{
		RawURL: fmt.Sprintf("https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s", token),
		Method: "POST",
		Body:   body,
	})
	if err != nil {
		return nil, err
	}

	result := Error{}
	err = client.DecodeResponse(resp, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return nil, err
}
