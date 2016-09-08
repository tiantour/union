package union

import (
	"github.com/tiantour/conf"
	"github.com/tiantour/requests"
)

// AccessToken
func (w *weibo) UserInfo(accessToken, uID string) ([]byte, error) {
	requestURL, requestData, requestHeader := requests.Options()
	requestURL = "https://api.weibo.com/2/users/show.json?source=" + conf.Options.Weibo.AppID + "&access_token=" + accessToken + "&uid=" + uID
	return requests.Get(requestURL, requestData, requestHeader)
}
