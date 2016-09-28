package union

import (
	"github.com/tiantour/union/qq"
	"github.com/tiantour/union/wechat"
	"github.com/tiantour/union/weibo"
)

// union union
var (
	Wechat = &tWechat{}
	Weibo  = &tWeibo{}
	QQ     = &tQQ{}
)

type (
	tWechat struct {
		User    wechat.User
		Token   wechat.Token
		Message wechat.Message
	}
	tWeibo struct {
		User weibo.User
	}
	tQQ struct {
		User qq.User
	}
)
