package mp

import (
	"encoding/json"
	"errors"
	"fmt"

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

// Watermark water mark
type Watermark struct {
	AppID     string `json:"appid,omitempty"`
	TimeStamp int    `json:"timestamp,omitempty"`
}

// NewMP new mini program
func NewMP() *MP {
	return &MP{}
}

// Verify verify
// date 2017-06-19
// author andy.jiang
func (m MP) Verify(rawData MP, signature string) bool {
	body, err := json.Marshal(rawData)
	if err != nil {
		return false
	}
	data := rsae.NewRsae().SHA1(fmt.Sprintf("%s%s",
		string(body),
		SessionKey,
	))
	if signature != fmt.Sprintf("%x", data) {
		return false
	}
	return true
}

// User user
// date 2017-06-19
// author andy.jiang
func (m MP) User(encryptedData, iv string) (MP, error) {
	result := MP{}
	encryptedByte, err := rsae.NewRsae().Base64Decode(encryptedData)
	if err != nil {
		return result, err
	}
	sessionByte, err := rsae.NewRsae().Base64Decode(SessionKey)
	if err != nil {
		return result, err
	}
	ivByte, err := rsae.NewRsae().Base64Decode(iv)
	if err != nil {
		return result, err
	}
	body, err := rsae.NewRsae().AESDecrypt(encryptedByte, sessionByte, ivByte)
	if err != nil {
		return result, err
	}
	data := WMP{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return result, err
	}
	if data.Watermark.AppID != AppID {
		return result, errors.New("appid not match")
	}
	return data.MP, nil
}
