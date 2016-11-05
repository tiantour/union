package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// Access token
func (t *Token) Access(appID, appSecret, code string) (Token, error) {
	result := Token{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		appID,
		appSecret,
		code,
	)
	body, err := request(url)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	if result.ErrCode != 0 {
		return result, errors.New(strconv.Itoa(result.ErrCode) + result.ErrMsg)
	}
	return result, nil
}

// Refresh token
func (t *Token) Refresh(appID, refreshToken string) (Token, error) {
	result := Token{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s",
		appID,
		refreshToken,
	)
	body, err := request(url)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Verify token
func (t *Token) Verify(accessToken, openID string) (Prompt, error) {
	result := Prompt{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s",
		accessToken,
		openID,
	)
	body, err := request(url)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
