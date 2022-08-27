package mx

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tiantour/fetch"
)

type Token struct {
	AccessToken string `json:"access_token,omitempty"`
	ExpiresIN   int    `json:"expires_in,omitempty"`
	Error
}

func NewToken() *Token {
	return new(Token)
}

func (t *Token) Access() (*Token, error) {
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
		return nil, errors.New(result.ErrMSG)
	}
	return &result, err
}
