package union

import (
	"github.com/tiantour/conf"
	"github.com/tiantour/requests"
)

// code 获取授权
func (w *wechat) Code(uri, scope string) string {
	return "https://open.weixin.qq.com/connect/oauth2/authorize?appid=" + conf.Options.Wechat.AppID + "&redirect_uri=" + uri + "&response_type=code&scope=snsapi_" + scope + "&state=STATE#wechat_redirect"
}

// accessToken 获取token
func (w *wechat) AccessToken(code string) ([]byte, error) {
	url := "https://api.weixin.qq.com/sns/oauth2/access_token?appid=" + conf.Options.Wechat.AppID + "&secret=" + conf.Options.Wechat.AppSecret + "&code=" + code + "&grant_type=authorization_code"
	return w.requestOperate(url)
}

// refreshToken 刷新token
func (w *wechat) RefreshToken(refreshToken string) ([]byte, error) {
	url := "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=" + conf.Options.Wechat.AppID + "&grant_type=refresh_token&refresh_token=" + refreshToken
	return w.requestOperate(url)
}

// verifyToken
func (w *wechat) VerifyToken(accessToken, openID string) ([]byte, error) {
	url := "https://api.weixin.qq.com/sns/auth?access_token=" + accessToken + "&openid=" + openID
	return w.requestOperate(url)
}

// userInfo 用户资料
func (w *wechat) UserInfo(accessToken, openID string) ([]byte, error) {
	url := "https://api.weixin.qq.com/sns/userinfo?access_token=" + accessToken + "&openid=" + openID
	return w.requestOperate(url)
}

// requestOperate
func (w *wechat) requestOperate(url string) ([]byte, error) {
	requestURL, requestData, requestHeader := requests.Options()
	requestURL = url
	return requests.Get(requestURL, requestData, requestHeader)
}
