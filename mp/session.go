package mp

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tiantour/buffer"
	"github.com/tiantour/fetch"
)

// Session session
type Session struct {
	OpenID     string `json:"openid"`      // openid
	SessionKey string `json:"session_key"` // session
	UnionID    string `json:"unionid"`     // unionid
	ErrCode    int    `json:"errcode"`     // 错误码
	ErrMsg     string `json:"errmsg"`      // 错误提示
}

// NewSession new session
func NewSession() *Session {
	return &Session{}
}

// Get get
// date 2017-06-19
// author andy.jiang
func (s Session) Get(code string) (Session, error) {
	result := Session{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		AppID,
		AppSecret,
		code,
	)
	body, err := fetch.Cmd(fetch.Request{
		URL:    url,
		Method: "GET",
	})
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	if result.ErrCode != 0 {
		return result, errors.New(result.ErrMsg)
	}
	// 写入缓存
	_, err = buffer.NewHash().Add("mp", AppID, map[string]interface{}{
		result.OpenID: result.SessionKey,
	})
	return result, err
}
