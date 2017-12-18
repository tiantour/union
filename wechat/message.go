package wechat

import (
	"fmt"

	"github.com/tiantour/fetch"
)

// Message message
type Message struct{}

// NewMessage new message
func NewMessage() *Message {
	return &Message{}
}

// WAP wap
func (m Message) WAP(body []byte) ([]byte, error) {
	return m.do(0, body)
}

// MP mp
func (m Message) MP(body []byte) ([]byte, error) {
	return m.do(1, body)
}

// do do
func (m Message) do(category int, body []byte) ([]byte, error) {
	token, err := NewToken().Cache()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s", token)
	if category != 0 {
		url = fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token=%s", token)
	}
	return fetch.Cmd(fetch.Request{
		Method: "POST",
		URL:    url,
		Body:   body,
	})
}
