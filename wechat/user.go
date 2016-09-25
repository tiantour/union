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
	if err == nil && json.Unmarshal(body, &result) == nil {
		return result, nil
	}
	return result, err
}
