package mx

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tiantour/fetch"
	"github.com/tiantour/imago"
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

func (t *Ticket) Get(token string) (*Ticket, error) {
	data := &Ticket{
		ActionName:    "QR_SCENE",
		ExpireSeconds: 648000,
		ActionInfo: &Action{
			Scene: &Scene{
				SceneID: imago.NewRandom().Number(16),
			},
		},
	}

	body, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	body, err = fetch.Cmd(&fetch.Request{
		Method: "POST",
		URL:    fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=%s", token),
		Body:   body,
	})
	if err != nil {
		return nil, err
	}

	result := Ticket{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMSG)
	}
	return &result, err
}
