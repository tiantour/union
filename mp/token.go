package mp

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/tiantour/fetch"
	"github.com/tiantour/union/x/cache"
)

// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/mp-access-token/getAccessToken.html

// Token token
type Token struct {
	AccessToken  string `json:"access_token"`  // 网页授权接口调用凭证,注意：此access_token与基础支持的access_token不同
	ExpiresIn    int    `json:"expires_in"`    // access_token接口调用凭证超时时间，单位（秒）
	RefreshToken string `json:"refresh_token"` // 用户刷新access_token
	OpenID       string `json:"openid"`        // 用户唯一标识
	Scope        string `json:"scope"`         // 用户授权的作用域，使用逗号（,）分隔
	Error
}

// NewToken new token
func NewToken() *Token {
	return &Token{}
}

// Access access token
func (t *Token) Access() (string, error) {
	token, ok := cache.NewString().Get(AppID)
	if ok && token != "" {
		return token.(string), nil
	}

	result, err := t.Get()
	if err != nil {
		return "", err
	}

	_ = cache.NewString().Set(AppID, result.AccessToken, 1, 7200*time.Second)
	return result.AccessToken, nil
}

func (t *Token) Get() (*Token, error) {
	body, err := fetch.Cmd(&fetch.Request{
		Method: "GET",
		URL:    fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", AppID, AppSecret),
	})
	if err != nil {
		return nil, err
	}

	result := Token{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, err
}
