package wechat

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tiantour/fetch"
)

/*

微信网页授权

https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1421140842

*/

// Token token
type Token struct {
	AccessToken  string `json:"access_token"`  // 网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
	ExpiresIn    int    `json:"expires_in"`    // access_token接口调用凭证超时时间，单位（秒）
	RefreshToken string `json:"refresh_token"` // 用户刷新access_token
	OpenID       string `json:"openid"`        // 用户唯一标识
	Scope        string `json:"scope"`         // 用户授权的作用域，使用逗号（,）分隔
	ErrCode      int    `json:"errcode"`       // 错误代码
	ErrMsg       string `json:"errmsg"`        // 错误消息
}

// NewToken new token
func NewToken() *Token {
	return &Token{}
}

// Access token
func (t *Token) Access(code string) (*Token, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		AppID,
		AppSecret,
		code,
	)
	return t.do(url)
}

// Refresh token
func (t *Token) Refresh(refreshToken string) (*Token, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s",
		AppID,
		refreshToken,
	)
	return t.do(url)
}

// Verify token
func (t *Token) Verify(accessToken, openID string) (*Token, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s",
		accessToken,
		openID,
	)
	return t.do(url)
}

// do do
func (t *Token) do(url string) (*Token, error) {
	result := Token{}
	body, err := fetch.Cmd(&fetch.Request{
		Method: "GET",
		URL:    url,
	})
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, err
}
