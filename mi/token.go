package mi

import (
	"errors"
	"fmt"
	"net/url"
	"os"

	"github.com/duke-git/lancet/v2/datetime"
	"github.com/duke-git/lancet/v2/netutil"
	"github.com/google/go-querystring/query"
	"github.com/tiantour/rsae"
)

// Token token
type Token struct{}

// NewToken new token
func NewToken() *Token {
	return &Token{}
}

// Access access token
func (t *Token) Access(code string) (*Response, error) {
	args := &Request{
		AppID:     AppID,
		Method:    "alipay.system.oauth.token",
		Format:    "JSON",
		Charset:   "utf-8",
		SignType:  "RSA2",
		TimeStamp: datetime.GetNowDateTime(),
		Version:   "1.0",
		GrantType: "authorization_code",
		Code:      code,
	}

	tmp, err := query.Values(args)
	if err != nil {
		return nil, err
	}

	sign, err := t.Sign(&tmp, PrivatePath)
	if err != nil {
		return nil, err
	}

	client := netutil.NewHttpClient()
	resp, err := client.SendRequest(&netutil.HttpRequest{
		RawURL: fmt.Sprintf("https://openapi.alipay.com/gateway.do?%s", sign),
		Method: "GET",
	})
	if err != nil {
		return nil, err
	}

	result := Result{}
	err = client.DecodeResponse(resp, &result)
	if err != nil {
		return nil, err
	}

	response := result.AlipaySystemOauthTokenResponse
	if response.Code != "" {
		return nil, errors.New(response.Msg)
	}
	return response, nil
}

// Sign trade sign
func (t *Token) Sign(args *url.Values, privatePath string) (string, error) {
	query, err := url.QueryUnescape(args.Encode())
	if err != nil {
		return "", err
	}

	privateKey, err := os.ReadFile(privatePath)
	if err != nil {
		return "", err
	}

	sign, err := rsae.NewRSA().Sign(query, privateKey)
	if err != nil {
		return "", err
	}
	args.Add("sign", sign)

	return args.Encode(), nil
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

	publicKey, err := os.ReadFile(publicPath)
	if err != nil {
		return err
	}

	ok, err := rsae.NewRSA().Verify(query, sign, publicKey)
	if !ok {
		return err
	}
	return nil
}
