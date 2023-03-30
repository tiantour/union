package mx

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/duke-git/lancet/v2/netutil"
)

// https://developers.weixin.qq.com/doc/offiaccount/Account_Management/Generating_a_Parametric_QR_Code.html

type (
	Ticket struct {
		ActionName    string  `json:"action_name"`
		ExpireSeconds int     `json:"expire_seconds"`
		ActionInfo    *Action `json:"action_info"`
		Ticket        string  `json:"ticket"`
		URL           string  `json:"url"`
		Error
	}

	Action struct {
		Scene *Scene `json:"scene"`
	}

	Scene struct {
		SceneID  int    `json:"scene_id"`
		SceneSTR string `json:"scene_str"`
	}
)

func NewTicket() *Ticket {
	return new(Ticket)
}

func (t *Ticket) Get(args *Ticket) (*Ticket, error) {
	token, err := NewToken().Access()
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	client := netutil.NewHttpClient()
	resp, err := client.SendRequest(&netutil.HttpRequest{
		RawURL: fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=%s", token),
		Method: "POST",
		Body:   body,
	})
	if err != nil {
		return nil, err
	}

	result := Ticket{}
	err = client.DecodeResponse(resp, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMSG)
	}
	return &result, err
}
