package alipay

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/tiantour/fetch"
	"github.com/tiantour/imago"
	"github.com/tiantour/rsae"
	"github.com/tiantour/tempo"
)

// Token token
type Token struct {
	UserID       int64  `json:"user_id,omitempty"`       // 是 支付宝用户的唯一userId
	AccessToken  string `json:"access_token,omitempty"`  // 是 访问令牌。通过该令牌调用需要授权类接口
	ExpiresIn    int32  `json:"expires_in,omitempty"`    // 是 访问令牌的有效时间，单位是秒。
	RefreshToken string `json:"refresh_token,omitempty"` // 是 刷新令牌。通过该令牌可以刷新access_token
	ReExpiresIn  int32  `json:"re_expires_in,omitempty"` // 是 刷新令牌的有效时间，单位是秒。
}

// NewToken new token
func NewToken() *Token {
	return &Token{}
}

// Access access token
func (t *Token) Access(code string) (string, error) {
	args := &Request{
		AppID:     AppID,
		Method:    "alipay.system.oauth.token",
		Format:    "JSON",
		Charset:   "utf-8",
		SignType:  "RSA2",
		TimeStamp: tempo.NewNow().String(),
		Version:   "1.0",
		Parameter: &Parameter{
			GrantType: "authorization_code",
			Code:      code,
		},
	}
	sign, err := t.Sign(args, PrivatePath)
	if err != nil {
		return "", err
	}
	url := fmt.Sprintf("https://openapi.alipay.com/gateway.do?%s", sign)
	body, err := fetch.Cmd(fetch.Request{
		Method: "GET",
		URL:    url,
	})
	if err != nil {
		return "", err
	}
	result := Token{}
	err = json.Unmarshal(body, &result)
	return result.AccessToken, err
}

// Sign trade sign
func (t *Token) Sign(args interface{}, privatePath string) (string, error) {
	fmt.Println(0, args, privatePath)
	params, err := query.Values(args)
	fmt.Println(1, params)
	if err != nil {
		return "", err
	}
	query, err := url.QueryUnescape(params.Encode())
	if err != nil {
		return "", err
	}
	fmt.Println(2, query)
	if err != nil {
		return "", err
	}
	privateKey, err := imago.NewFile().Read(privatePath)
	fmt.Println(3, privateKey)
	if err != nil {
		return "", err
	}
	sign, err := rsae.NewRSA().Sign(query, privateKey)
	fmt.Println(4, sign, err)
	if err != nil {
		return "", err
	}
	fmt.Println(5, fmt.Sprintf("%s&sign=%s",
		query,
		url.QueryEscape(sign),
	))
	return fmt.Sprintf("%s&sign=%s",
		query,
		url.QueryEscape(sign),
	), nil
}

// Verify verify
func (t *Token) Verify(args url.Values, publicPath string) error {
	sign := args.Get("sign")
	args.Del("sign")
	args.Del("sign_type")
	query, err := url.QueryUnescape(args.Encode())
	if err != nil {
		return err
	}
	publicKey, err := imago.NewFile().Read(publicPath)
	if err != nil {
		return err
	}
	ok, err := rsae.NewRSA().Verify(query, sign, publicKey)
	if !ok {
		return errors.New("签名错误")
	}
	return nil
}
