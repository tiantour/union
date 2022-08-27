package mp

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tiantour/fetch"
)

// QR qr
type QR struct {
	Scene     string                 `json:"scene"`      // 场景
	Page      string                 `json:"page"`       // 页面
	Width     int                    `json:"width"`      // 宽度
	AutoColor bool                   `json:"auto_color"` // 默认颜色
	LineColor map[string]interface{} `json:"line_color"` // 线条颜色
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

	body, err = fetch.Cmd(&fetch.Request{
		Method: "POST",
		URL:    fmt.Sprintf("https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s", token),
		Body:   body,
	})

	// 判断 body 大小
	if len(body) >= 256 {
		return body, err
	}

	data := Error{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return body, err
	}

	if data.ErrCode != 0 {
		return nil, errors.New(data.ErrMsg)
	}
	return nil, err
}
