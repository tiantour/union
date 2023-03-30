package wechat

import (
	"errors"
	"fmt"

	"github.com/duke-git/lancet/v2/netutil"
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
	Error
}

// NewToken new token
func NewToken() *Token {
	return &Token{}
}

func (t *Token) Access(code string) (*Token, error) {
	client := netutil.NewHttpClient()
	resp, err := client.SendRequest(&netutil.HttpRequest{
		RawURL: fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", AppID, AppSecret, code),
		Method: "GET",
	})
	if err != nil {
		return nil, err
	}

	result := Token{}
	err = client.DecodeResponse(resp, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, err
}
