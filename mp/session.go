package mp

import (
	"errors"
	"fmt"

	"github.com/duke-git/lancet/v2/netutil"
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
	client := netutil.NewHttpClient()
	resp, err := client.SendRequest(&netutil.HttpRequest{
		RawURL: fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", AppID, AppSecret, code),
		Method: "GET",
	})
	if err != nil {
		return nil, err
	}

	result := Session{}
	err = client.DecodeResponse(resp, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, err
}
