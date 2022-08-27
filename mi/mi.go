package mi

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tiantour/imago"
	"github.com/tiantour/rsae"
)

var (
	AppID string // AppID appid

	AesKey string // AesKey aes key

	PrivatePath string // PrivatePath private path

	PublicPath string // PublicPath public path
)

type (
	// MI mi
	MI struct{}

	// User user
	User struct {
		UserID      string `json:"user_id,omitempty"`     // 是 用户ID
		NickName    string `json:"nickName,omitempty"`    // 是 用户暱称
		Avatar      string `json:"avatar,omitempty"`      // 是 用户头像
		Gender      string `json:"gender,omitempty"`      // 否 用户性别
		CountryCode string `json:"countryCode,omitempty"` // 国家编码
		Province    string `json:"province,omitempty"`    // 是 省份
		City        string `json:"city,omitempty"`        // 是 城市
	}

	// Request request
	Request struct {
		AppID        string `json:"app_id,omitempty" url:"app_id,omitempty"`                 // 是 应用ID
		Method       string `json:"method,omitempty" url:"method,omitempty"`                 // 是 接口名称
		Format       string `json:"format,omitempty" url:"format,omitempty"`                 // 否 JSON
		Charset      string `json:"charset,omitempty" url:"charset,omitempty"`               // 是 utf-8
		SignType     string `json:"sign_type,omitempty" url:"sign_type,omitempty"`           // 是 RSA2
		Sign         string `json:"sign,omitempty" url:"sign,omitempty"`                     // 是 签名
		TimeStamp    string `json:"timestamp,omitempty" url:"timestamp,omitempty"`           // 是 时间
		Version      string `json:"version,omitempty" url:"version,omitempty"`               // 是 1.0
		AuthToken    string `json:"auth_token,omitempty" url:"auth_token,omitempty"`         // 是 用户授权
		AppAuthToken string `json:"app_auth_token,omitempty" url:"app_auth_token,omitempty"` // 否 应用授权
		GrantType    string `json:"grant_type,omitempty" url:"grant_type,omitempty"`         // 是 值为authorization_code时，代表用code换取；值为refresh_token时，代表用refresh_token换取
		Code         string `json:"code,omitempty" url:"code,omitempty"`                     // 否 授权码
		RefreshToken string `json:"refresh_token,omitempty" url:"refresh_token,omitempty"`   // 否 刷新令牌
		BizContent   string `json:"biz_content,omitempty" url:"biz_content,omitempty"`       // 请求参数的集合
	}

	// Response response
	Response struct {
		Code    string `json:"code,omitempty"`     // 是 网关返回码
		Msg     string `json:"msg,omitempty"`      // 是 网关返回码描述
		SubCode string `json:"sub_code,omitempty"` // 否 业务返回码
		SubMsg  string `json:"sub_msg,omitempty"`  // 是 业务返回码描述
		Sign    string `json:"sign,omitempty"`     // 是 签名
		*Next
	}

	// Next next
	Next struct {
		UserID       string `json:"user_id,omitempty"`       // 是 支付宝用户的唯一userId
		AccessToken  string `json:"access_token,omitempty"`  // 是 访问令牌。通过该令牌调用需要授权类接口
		ExpiresIn    int32  `json:"expires_in,omitempty"`    // 是 访问令牌的有效时间，单位是秒。
		RefreshToken string `json:"refresh_token,omitempty"` // 是 刷新令牌。通过该令牌可以刷新access_token
		ReExpiresIn  int32  `json:"re_expires_in,omitempty"` // 是 刷新令牌的有效时间，单位是秒。
		Moblie       string `json:"mobile"`                  // 手机号
		Response     string `json:"response,omitempty"`      // 否 内容
		QrCodeURL    string `json:"qr_code_url,omitempty"`   // 二维码图片链接地址
	}

	// Result result
	Result struct {
		AlipaySystemOauthTokenResponse    *Response `json:"alipay_system_oauth_token_response,omitempty"`     // 内容
		AlipayUserInfoShareResponse       *Response `json:"alipay_user_info_share_response,omitempty"`        // 内容
		AlipayOpenAppQrcodeCreateResponse *Response `json:"alipay_open_app_qrcode_create_response,omitempty"` // 内容
		Sign                              string    `json:"sign,omitempty"`                                   // 签名
	}
)

// NewMI new mi
func NewMI() *MI {
	return &MI{}
}

// User user
func (m *MI) User(code, content string) (*User, error) {
	response, err := NewToken().Access(code)
	if err != nil {
		return nil, err
	}

	user := User{}
	err = json.Unmarshal([]byte(content), &user)
	if err != nil {
		return nil, err
	}
	user.UserID = response.UserID

	return &user, err
}

// Phone phone
func (m *MI) Phone(content string) (*Response, error) {
	data := Response{}
	err := json.Unmarshal([]byte(content), &data)
	if err != nil {
		return nil, err
	}

	publicKey, err := imago.NewFile().Read(PublicPath)
	if err != nil {
		return nil, err
	}

	origdata := fmt.Sprintf(`"%s"`, data.Response)
	ok, err := rsae.NewRSA().Verify(origdata, data.Sign, publicKey)
	if !ok {
		return nil, err
	}

	iv := []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	ciphertext, err := rsae.NewBase64().Decode(data.Response)
	if err != nil {
		return nil, err
	}

	key, err := rsae.NewBase64().Decode(AesKey)
	if err != nil {
		return nil, err
	}

	body, err := rsae.NewAES().Decrypt(ciphertext, key, iv)
	if err != nil {
		return nil, err
	}

	result := Response{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	if result.Code != "10000" {
		return nil, errors.New(result.Msg)
	}
	return &result, nil
}
