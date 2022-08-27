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
	Error
}

// NewSession new session
func NewSession() *Session {
	return &Session{}
}

// Get get
func (s *Session) Get(code string) (*Session, error) {
	body, err := fetch.Cmd(&fetch.Request{
		Method: "GET",
		URL:    fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", AppID, AppSecret, code),
	})
	if err != nil {
		return nil, err
	}

	result := Session{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, err
}
