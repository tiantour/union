package wechat

import (
	"encoding/json"
	"fmt"

	"github.com/tiantour/cache"
	"github.com/tiantour/fetch"
)

// Token message
func (m *Message) token(appID, appSecret string) (string, error) {
	key := "wechat_access_token"
	token, err := cache.String.GET(key).Str()
	if err != nil {
		result := Token{}
		url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
			appID,
			appSecret,
		)
		body, err := request(url)
		if err == nil && json.Unmarshal(body, &result) == nil {
			_ = cache.String.SET(key, result.AccessToken)
			_ = cache.Key.EXPIRE(key, 7200)
			return result.AccessToken, nil
		}
		return "", err
	}
	return token, nil

}

// Push message
func (m *Message) Push(appID, appSecret string, data interface{}) ([]byte, error) {
	accessToken, err := m.token(appID, appSecret)
	if err != nil {
		return nil, err
	}
	requestURL := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", accessToken)
	requestData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return fetch.Cmd("post", requestURL, requestData)
}
