package wxwork

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tiantour/fetch"
)

var (
	// CorpID corpid
	CorpID string

	// CorpSecret corp secret
	CorpSecret string
)

// Wxwork Wxwork
type Wxwork struct {
	ErrCode        int    `json:"errcode"`         // 错误代码
	ErrMsg         string `json:"errmsg"`          // 错误消息
	UserID         string `json:"UserId"`          // 成员UserID
	DeviceID       string `json:"DeviceId"`        // 手机设备号
	OpenID         string `json:"OpenId"`          // 非企业成员的标识，对当前企业唯一
	ExternalUserID string `json:"external_userid"` // 外部联系人id，当且仅当用户是企业的客户，且跟进人在应用的可见范围内时返回
}

// NewWxwork new wxwork
func NewWxwork() *Wxwork {
	return &Wxwork{}
}

// User user
func (w *Wxwork) User(code string) (*Wxwork, error) {
	token, err := NewToken().Access()
	if err != nil {
		return nil, err
	}
	body, err := fetch.Cmd(&fetch.Request{
		Method: "GET",
		URL: fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=%s&code=%s",
			token.AccessToken,
			code,
		),
	})
	if err != nil {
		return nil, err
	}
	result := Wxwork{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, err
}
