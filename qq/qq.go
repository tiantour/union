package qq

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tiantour/fetch"
)

var (
	AppID string // AppID appID
)

type (
	// QQ qq
	QQ struct{}

	// User user
	User struct {
		Ret             int    `json:"ret"`                // 返回码
		Msg             string `json:"msg"`                // 如果ret<0，会有相应的错误信息提示，返回数据全部用UTF-8编码。
		NickName        string `json:"nickname"`           // 用户在QQ空间的昵称。
		FigureURL       string `json:"figureurl"`          // 大小为30×30像素的QQ空间头像URL。
		FigureURL1      string `json:"figureurl_1"`        // 大小为50×50像素的QQ空间头像URL。
		FigureURL2      string `json:"figureurl_2"`        // 大小为100×100像素的QQ空间头像URL。
		FigureURLQQ1    string `json:"figureurl_qq_1"`     // 大小为40×40像素的QQ头像URL。
		FigureURLQQ2    string `json:"figureurl_qq_2"`     // 大小为100×100像素的QQ头像URL。需要注意，不是所有的用户都拥有QQ的100x100的头像，但40x40像素则是一定会有。
		Gender          string `json:"gender"`             // 性别。 如果获取不到则默认返回"男"
		ISYellowVip     string `json:"is_yellow_vip"`      // 标识用户是否为黄钻用户（0：不是；1：是）。
		Vip             string `json:"vip"`                // 标识用户是否为黄钻用户（0：不是；1：是）
		YelloVipLevel   string `json:"yellow_vip_level"`   // 黄钻等级
		Level           string `json:"level"`              // 黄钻等级
		IsYellowYearVip string `json:"is_yellow_year_vip"` // 标识是否为年费黄钻用户（0：不是； 1：是）
	}
)

// NewQQ new qq
func NewQQ() *QQ {
	return &QQ{}
}

// User user
func (q *QQ) User(accessToken, openID string) (*User, error) {
	body, err := fetch.Cmd(&fetch.Request{
		Method: "GET",
		URL:    fmt.Sprintf("https://graph.qq.com/user/get_user_info?access_token=%s&oauth_consumer_key=%s&openid=%s", accessToken, AppID, openID),
	})
	if err != nil {
		return nil, err
	}

	result := User{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if result.Ret != 0 {
		return nil, errors.New(result.Msg)
	}
	return &result, err
}
