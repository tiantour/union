package ums

var (
	// AppID appid
	AppID string

	// AppKey app key
	AppKey string
)

type (
	// Request request
	Request struct {
		AppID      string `json:"appId,omitempty"`      // 是 AppId
		Timestamp  string `json:"timestamp,omitempty"`  // 是 时间戳
		Nonce      string `json:"nonce,omitempty"`      // 是 随机数
		SignMethod string `json:"signMethod,omitempty"` // 是 签名方法
		Signature  string `json:"signature,omitempty"`  // 是 签名
	}

	// Response response
	Response struct {
		ErrCode     string `json:"errCode,omitempty"`     // 是 错误代码
		ErrInfo     string `json:"errInfo,omitempty"`     // 是 错误说明
		AccessToken string `json:"accessToken,omitempty"` // 是 授权令牌
		ExpiresIn   int    `json:"expiresIn,omitempty"`   // 是 失效时间
	}
)
