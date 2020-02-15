package mp

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tiantour/fetch"
)

// Session session
type Session struct {
	OpenID     string `json:"openid"`      // openid
	SessionKey string `json:"session_key"` // session
	UnionID    string `json:"unionid"`     // unionid
	Result
}

// NewSession new session
func NewSession() *Session {
	return &Session{}
}

// Get get
// date 2017-06-19
// author andy.jiang
func (s *Session) Get(code string) (*Session, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		AppID,
		AppSecret,
		code,
	)
	return s.do(url)
}

// do
func (s *Session) do(url string) (*Session, error) {
	result := Session{}
	body, err := fetch.Cmd(&fetch.Request{
		URL:    url,
		Method: "GET",
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
