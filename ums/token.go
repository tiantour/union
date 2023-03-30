package ums

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/duke-git/lancet/v2/netutil"
	"github.com/duke-git/lancet/v2/random"
	"github.com/tiantour/rsae"
	"github.com/tiantour/union/x/cache"
)

// Token token
type Token struct{}

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

func (t *Token) Get() (*Response, error) {
	data := &Request{
		AppID:      AppID,
		Timestamp:  time.Now().Format("20060102150405"),
		Nonce:      random.RandString(32),
		SignMethod: "SHA256",
	}
	sign := rsae.NewSHA().SHA256(AppID + data.Timestamp + data.Nonce + AppKey)
	data.Signature = string(hex.EncodeToString(sign))

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	header := http.Header{}
	header.Add("Accept", "application/json")
	header.Add("Content-Type", "application/json;charset=utf-8")

	client := netutil.NewHttpClient()
	resp, err := client.SendRequest(&netutil.HttpRequest{
		RawURL:  "https://api-mop.chinaums.com/v1/token/access",
		Method:  "POST",
		Body:    body,
		Headers: header,
	})
	if err != nil {
		return nil, err
	}

	result := Response{}
	err = client.DecodeResponse(resp, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != "0000" {
		return nil, errors.New(result.ErrInfo)
	}
	return &result, nil
}
