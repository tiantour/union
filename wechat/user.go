package wechat

import (
	"encoding/json"
	"fmt"
)

// Info user
func (u *User) Info(accessToken, openID string) (User, error) {
	result := User{}
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s",
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
