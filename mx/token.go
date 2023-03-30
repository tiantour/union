package mx

import (
	"errors"
	"fmt"
	"time"

	"github.com/duke-git/lancet/v2/netutil"
	"github.com/tiantour/union/x/cache"
)

type Token struct {
	AccessToken string `json:"access_token,omitempty"`
	ExpiresIN   int    `json:"expires_in,omitempty"`
	Error
}

func NewToken() *Token {
	return new(Token)
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
	client := netutil.NewHttpClient()
	resp, err := client.SendRequest(&netutil.HttpRequest{
		RawURL: fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", AppID, AppSecret),
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
		return nil, errors.New(result.ErrMSG)
	}
	return &result, err
}
