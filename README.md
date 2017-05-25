# union
union login for wechat,weibo,qq with go

###　weibo

```
package main

import (
	"fmt"

	"github.com/tiantour/union/weibo"
)

func main() {
	user, err := weibo.NewWeibo().User("your AccessToken", "your UID")
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
	user, err := qq.NewQQ().User("your AccessToken", "your OpenID")
	fmt.Println(user, err)
}
```

### wechat


user 

```
package main

import (
	"fmt"

	"github.com/tiantour/union/wechat"
)

func main() {
	// access token
	token, err := wechat.NewToken().Access("your code")
	if err != nil {
		fmt.Println(err)
	}
	// user info
	user, err := wechat.NewWechat().User(token.AccessToken, token.OpenID)
	fmt.Println(user, err)
}
```


share 


```
package main

import (
	"fmt"

	"github.com/tiantour/union/wechat"
)

func main() {
	url := "your url"
	share, err := wechat.NewShare().Do(url)
	fmt.Println(share, err)
}
```

message

```
package main

import (
	"fmt"

	"github.com/tiantour/union/wechat"
)

func main() {
	data := []byte{}
	message, err := wechat.NewMessage().Do(data)
	fmt.Println(message, err)
}
```