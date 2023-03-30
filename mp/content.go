package mp

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/duke-git/lancet/v2/netutil"
)

type (
	// Image image
	Image struct {
		Media []byte `json:"media"` // 媒体
	}

	// Message message
	Message struct {
		Content string `json:"content"` // 内容
	}
)

type Content struct{}

func NewContent() *Content {
	return new(Content)
}

// Image image
func (c *Content) Image(args *Image) ([]byte, error) {
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
		RawURL: fmt.Sprintf("https://api.weixin.qq.com/wxa/img_sec_check?access_token=%s", token),
		Method: "POST",
		Body:   body,
	})
	if err != nil {
		return nil, err
	}

	result := Error{}
	err = client.DecodeResponse(resp, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return body, err
}

// Message message
func (c *Content) Message(args *Message) ([]byte, error) {
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
		RawURL: fmt.Sprintf("https://api.weixin.qq.com/wxa/msg_sec_check?access_token=%s", token),
		Method: "POST",
		Body:   body,
	})
	if err != nil {
		return nil, err
	}

	result := Error{}
	err = client.DecodeResponse(resp, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return body, err
}
