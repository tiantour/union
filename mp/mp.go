package mp

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tiantour/fetch"
	"github.com/tiantour/rsae"
)

var (
	// AppID appid
	AppID string

	// AppSecret app secret
	AppSecret string

	// SessionKey sessionKey
	SessionKey string
)

// WMP watermark mp
type WMP struct {
	MP
	Watermark Watermark `json:"watermark"` // 水印
}

// WP wechat phone
type WP struct {
	Phone
	Watermark Watermark `json:"watermark"` // 水印
}

// MP  mini program
type MP struct {
	NickName  string `json:"nickName"`          // 用户昵称
	Gender    int    `json:"gender"`            // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Language  string `json:"language"`          // 语言
	City      string `json:"city"`              // 普通用户个人资料填写的城市
	Province  string `json:"province"`          // 用户个人资料填写的省份
	Country   string `json:"country"`           // 国家，如中国为CN
	AvatarURL string `json:"avatarUrl"`         // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	OpenID    string `json:"openid,omitempty"`  // 用户的唯一标识
	UnionID   string `json:"unionid,omitempty"` // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。详见：获取用户个人信息（UnionID机制）
	OpenGID   string `json:"openGId"`           // 群对当前小程序的唯一 ID
}

// Phone phone
type Phone struct {
	PhoneNumber     string `json:"phoneNumber"`     // 用户绑定的手机号
	PurePhoneNumber string `json:"purePhoneNumber"` // 没有区号的手机号
	CountryCode     string `json:"countryCode"`     // 区号
}

// QR qr
type QR struct {
	Scene     string                 `json:"scene"`      // 场景
	Page      string                 `json:"page"`       // 页面
	Width     int                    `json:"width"`      // 宽度
	AutoColor bool                   `json:"auto_color"` // 默认颜色
	LineColor map[string]interface{} `json:"line_color"` // 线条颜色
}

// Watermark water mark
type Watermark struct {
	AppID     string `json:"appid,omitempty"`
	TimeStamp int    `json:"timestamp,omitempty"`
}

// NewMP new mini program
func NewMP() *MP {
	return &MP{}
}

// User user
func (m *MP) User(encryptedData, iv string) (*MP, error) {
	result := MP{}
	encryptedByte, err := rsae.NewBase64().Decode(encryptedData)
	if err != nil {
		return nil, err
	}
	sessionByte, err := rsae.NewBase64().Decode(SessionKey)
	if err != nil {
		return nil, err
	}
	ivByte, err := rsae.NewBase64().Decode(iv)
	if err != nil {
		return nil, err
	}
	body, err := rsae.NewAES().Decrypt(encryptedByte, sessionByte, ivByte)
	if err != nil {
		return nil, err
	}
	data := WMP{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	if data.Watermark.AppID != AppID {
		return nil, errors.New("appid not match")
	}
	return &data.MP, nil
}

// Verify verify
func (m *MP) Verify(rawData *MP, signature string) bool {
	body, err := json.Marshal(rawData)
	if err != nil {
		return false
	}
	tmp := fmt.Sprintf("%s%s", string(body), SessionKey)
	data := rsae.NewSHA().SHA1(tmp)
	if signature != string(data) {
		return false
	}
	return true
}

// Phone phone
func (m *MP) Phone(encryptedData, iv string) (*Phone, error) {
	result := Phone{}
	encryptedByte, err := rsae.NewBase64().Decode(encryptedData)
	if err != nil {
		return nil, err
	}
	sessionByte, err := rsae.NewBase64().Decode(SessionKey)
	if err != nil {
		return nil, err
	}
	ivByte, err := rsae.NewBase64().Decode(iv)
	if err != nil {
		return nil, err
	}
	body, err := rsae.NewAES().Decrypt(encryptedByte, sessionByte, ivByte)
	if err != nil {
		return nil, err
	}
	data := WP{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	if data.Watermark.AppID != AppID {
		return nil, errors.New("appid not match")
	}
	return &data.Phone, nil
}

// QR qr
func (m *MP) QR(args *QR) ([]byte, error) {
	body, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	token, err := NewToken().Access()
	if err != nil {
		return nil, err
	}
	return fetch.Cmd(fetch.Request{
		Method: "POST",
		URL:    fmt.Sprintf("https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s", token),
		Body:   body,
	})
}
