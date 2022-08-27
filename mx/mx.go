package mx

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tiantour/fetch"
)

// https://developers.weixin.qq.com/doc/offiaccount/User_Management/Get_users_basic_information_UnionID.html#UinonId
var (
	AppID string // AppID appid

	AppSecret string // AppSecret app secret
)

type (
	MX struct{}

	Error struct {
		ErrCode int    `json:"errcode,omitempty"`
		ErrMSG  string `json:"errmsg,omitempty"`
	}

	User struct {
		Subscribe      int    `json:"subscribe,omitempty"`
		OpenID         string `json:"openid,omitempty"`
		Language       string `json:"language,omitempty"`
		SubscribeTime  int    `json:"subscribe_time,omitempty"`
		UnionID        string `json:"unionid,omitempty"`
		Remark         string `json:"remark,omitempty"`
		GroupID        int    `json:"groupid,omitempty"`
		TagIDList      []int  `json:"tagid_list,omitempty"`
		SubscribeBcene string `json:"subscribe_scene,omitempty"`
		QRScene        int    `json:"qr_scene,omitempty"`
		QRSceneStr     string `json:"qr_scene_str,omitempty"`
		Error
	}
)

func NewMX() *MX {
	return new(MX)
}

func (m *MX) User(openID string) (*User, error) {
	token, err := NewToken().Access()
	if err != nil {
		return nil, err
	}

	body, err := fetch.Cmd(&fetch.Request{
		Method: "GET",
		URL:    fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN", token.AccessToken, openID),
	})
	if err != nil {
		return nil, err
	}

	result := User{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMSG)
	}
	return &result, err

}
