package wechat

import (
	"encoding/json"
	"fmt"

	"github.com/tiantour/conf"
)

// Access token
func (t *Token) Access(code string) (Token, error) {
	result := Token{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		conf.Options.Wechat.AppID,
		conf.Options.Wechat.AppSecret,
		code,
	)
	body, err := request(url)
	if err == nil && json.Unmarshal(body, &result) == nil {
		return result, nil
	}
	return result, err
}

// Refresh token
func (t *Token) Refresh(refreshToken string) (Token, error) {
	result := Token{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s",
		conf.Options.Wechat.AppID,
		refreshToken,
	)
	body, err := request(url)
	if err == nil && json.Unmarshal(body, &result) == nil {
		return result, nil
	}
	return result, err
}

// Verify token
func (t *Token) Verify(accessToken, openID string) (Message, error) {
	result := Message{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s",
		accessToken,
		openID,
	)
	body, err := request(url)
	if err == nil && json.Unmarshal(body, &result) == nil {
		return result, nil
	}
	return result, err
}
