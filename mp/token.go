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
	Result
}

// NewToken new token
func NewToken() *Token {
	return &Token{}
}

// Access access token
func (t *Token) Access() (string, error) {
	key := fmt.Sprintf("string:data:bind:access:token:%s", AppID)
	var token string
	err := cache.NewString().GET(&token, key)
	if err != nil {
		url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
			AppID,
			AppSecret,
		)
		result, err := t.do(url)
		if err != nil {
			return "", err
		}
		_ = cache.NewString().SET(nil, key, result.AccessToken, "EX", 7200)
		return result.AccessToken, nil
	}
	return token, nil
}

// do do
func (t *Token) do(url string) (*Token, error) {
	result := Token{}
	body, err := fetch.Cmd(fetch.Request{
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
