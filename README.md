# union
union login component developed by go. support Alipay, WeChat, QQ, Weibo, UmsPay, Mini Program, Wxwork

### alipay
```
package main

import (
	"fmt"

	"github.com/tiantour/union/mi"
)

func main() {
    mi.AppID = "your AppID"
    mi.PrivatePath = "your PrivateKey path"
    mi.PublicPath = "your PublicKey path"
	user, err := mi.NewMI().User("your code", "your content")
	fmt.Println(user, err)
}
```

### wechat
```
package main

import (
	"fmt"

	"github.com/tiantour/union/wechat"
)

func main() {
    wechat.AppID = "your AppID"
    wechat.AppSecret = "your AppSecret"
	user, err := wechat.NewWechat().User("your code")
	fmt.Println(user, err)
}
```

### qq

```
package main

import (
	"fmt"

	"github.com/tiantour/union/qq"
)

func main() {
    qq.AppID = "your AppID"
	user, err := qq.NewQQ().User("your AccessToken", "your OpenID")
	fmt.Println(user, err)
}
```

### weibo

```
package main

import (
	"fmt"

	"github.com/tiantour/union/weibo"
)

func main() {
    weibo.AppID = "your AppID"
	user, err := weibo.NewWeibo().User("your accessToken", "your UID")
	fmt.Println(user, err)
}
```

### umsPay

```
package main

import (
	"fmt"

	"github.com/tiantour/union/ums"
)

func main() {
    ums.AppID = "your AppID"
    ums.AppKey = "your AppKey"
	user, err := ums.NewToken().Access()
	fmt.Println(user, err)
}
```

### mini program
```
package main

import (
	"fmt"

	"github.com/tiantour/union/mp"
)

func main() {
    mp.AppID = "your AppID"
    mp.AppSecret = "your AppSecret"

    // new 
	data, err := mp.NewSession().Get("your code")
	fmt.Println(data, err)

    // old
    data, err := mp.NewMP().User(encryptedData, iv)
	fmt.Println(data, err)

}
```

### wxwork
```
package main

import (
	"fmt"

	"github.com/tiantour/union/wxwork"
)

func main() {
    wxwork.CorpID = "your CorpID"
    wxwork.CorpSecret = "your CorpSecret" 
	user, err := mp.NewWxwork().User(code)
	fmt.Println(user, err)
}
```