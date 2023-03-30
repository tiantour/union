package mp

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/duke-git/lancet/v2/netutil"
	"github.com/tiantour/rsae"
)

// https://developers.weixin.qq.com/miniprogram/dev/OpenApiDoc/user-info/phone-number/getPhoneNumber.html

var (
	AppID string // AppID appid

	AppSecret string // AppSecret app secret

	SessionKey string // SessionKey sessionKey
)

type (
	// MP mp
	MP struct{}

	// Error Error
	Error struct {
		ErrCode int    `json:"errcode"` // 错误代码
		ErrMsg  string `json:"errmsg"`  // 错误消息
	}

	// User user
	User struct {
		NickName  string    `json:"nickName"`          // 用户昵称
		Gender    int       `json:"gender"`            // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
		Language  string    `json:"language"`          // 语言
		City      string    `json:"city"`              // 普通用户个人资料填写的城市
		Province  string    `json:"province"`          // 用户个人资料填写的省份
		Country   string    `json:"country"`           // 国家，如中国为CN
		AvatarURL string    `json:"avatarUrl"`         // 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
		OpenID    string    `json:"openid,omitempty"`  // 用户的唯一标识
		UnionID   string    `json:"unionid,omitempty"` // 只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。详见：获取用户个人信息（UnionID机制）
		OpenGID   string    `json:"openGId"`           // 群对当前小程序的唯一 ID
		Watermark Watermark `json:"watermark"`         // 水印
	}

	// Phone phone
	Phone struct {
		PhoneNumber     string    `json:"phoneNumber"`     // 用户绑定的手机号
		PurePhoneNumber string    `json:"purePhoneNumber"` // 没有区号的手机号
		CountryCode     string    `json:"countryCode"`     // 区号
		Watermark       Watermark `json:"watermark"`       // 水印
	}

	// Watermark watermark
	Watermark struct {
		AppID     string `json:"appid,omitempty"`
		TimeStamp int    `json:"timestamp,omitempty"`
	}

	// WP wechat mobile
	WP struct {
		Error
		PhoneInfo Phone `json:"phone_info"` // 手机
	}
)

// NewMP new mini program
func NewMP() *MP {
	return &MP{}
}

// User user
func (m *MP) User(encryptedData, iv string) (*User, error) {
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

	data := User{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	if data.Watermark.AppID != AppID {
		return nil, errors.New("appid not match")
	}
	return &data, nil
}

// Phone phone
func (m *MP) Phone(code string) (*Phone, error) {
	token, err := NewToken().Access()
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(map[string]string{
		"code": code,
	})
	if err != nil {
		return nil, err
	}

	client := netutil.NewHttpClient()
	resp, err := client.SendRequest(&netutil.HttpRequest{
		RawURL: fmt.Sprintf("https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=%s", token),
		Method: "POST",
		Body:   body,
	})
	if err != nil {
		return nil, err
	}

	result := WP{}
	err = client.DecodeResponse(resp, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}

	return &result.PhoneInfo, nil
}
