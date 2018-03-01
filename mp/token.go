package mp

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tiantour/cache"
	"github.com/tiantour/fetch"
)

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
func (t Token) Access(code string) (Token, error) {
	result := Token{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		AppID,
		AppSecret,
		code,
	)
	body, err := fetch.Cmd(fetch.Request{
		Method: "GET",
		URL:    url,
	})
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	if err == nil && result.ErrCode != 0 {
		return result, errors.New(result.ErrMsg)
	}
	return result, err
}

// Refresh token
func (t Token) Refresh(refreshToken string) (Token, error) {
	result := Token{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s",
		AppID,
		refreshToken,
	)
	body, err := fetch.Cmd(fetch.Request{
		Method: "GET",
		URL:    url,
	})
	err = json.Unmarshal(body, &result)
	return result, err
}

// Verify token
func (t Token) Verify(accessToken, openID string) (Token, error) {
	result := Token{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s",
		accessToken,
		openID,
	)
	body, err := fetch.Cmd(fetch.Request{
		Method: "GET",
		URL:    url,
	})
	err = json.Unmarshal(body, &result)
	return result, err
}

// Cache cache
func (t Token) Cache() (string, error) {
	key := fmt.Sprintf("string:data:wechat:access:token:%s", AppID)
	token, err := cache.NewString().GET(key).Str()
	if err != nil || token == "" {
		result, err := t.Data()
		if err != nil {
			return token, err
		}
		_ = cache.NewString().SET(key, result.AccessToken)
		_ = cache.NewKey().EXPIRE(key, 7200)
		return result.AccessToken, nil
	}
	return token, err
}

// Data data
func (t Token) Data() (Token, error) {
	result := Token{}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
		AppID,
		AppSecret,
	)
	body, err := fetch.Cmd(fetch.Request{
		Method: "GET",
		URL:    url,
	})
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	return result, err
}
