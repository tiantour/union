package union

import (
	"github.com/tiantour/conf"
	"github.com/tiantour/requests"
)

// UserInfo
func (q *qq) UserInfo(accessToken, openID string) ([]byte, error) {
	requestURL, requestData, requestHeader := requests.Options()
	requestURL = "https://graph.qq.com/user/get_user_info?" + "access_token=" + accessToken + "&oauth_consumer_key=" + conf.Options.QQ.AppID + "&openid=" + openID
	return requests.Get(requestURL, requestData, requestHeader)
}
