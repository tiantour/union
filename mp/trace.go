package mp

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/duke-git/lancet/v2/netutil"
)

type (
	Trace struct{}

	TraceRequest struct {
		OpenID          string     `json:"openid,omitempty"`            // 是 用户openid
		SenderPhone     string     `json:"sender_phone,omitempty"`      // 否 寄件人手机号
		ReceiverPhone   string     `json:"receiver_phone,omitempty"`    // 是 收件人手机号，部分运力需要用户手机号作为查单依据
		DeliveryID      string     `json:"delivery_id,omitempty"`       // 否	运力id（运单号所属运力公司id），该字段从 get_delivery_list 获取。
		WaybillID       string     `json:"waybill_id,omitempty"`        // 是 运单号
		TransID         string     `json:"trans_id,omitempty"`          // 是 微信支付id
		OrderTetailTath string     `json:"order_detail_path,omitempty"` // 否 点击落地页商品卡片跳转路径（建议为订单详情页path），不传默认跳转小程序首页。
		GoodsInfo       DetailList `json:"goods_info,omitempty"`        // 是	商品信息
	}

	TraceResponse struct {
		Error
		WaybillInfo  *WaybillInfo  `json:"waybill_info,omitempty"`  // 运单信息
		ShopInfo     *ShopInfo     `json:"shop_info,omitempty"`     // 店铺信息
		DeliveryInfo *DeliveryInfo `json:"delivery_info,omitempty"` // 运力信息
	}

	ShopInfo struct {
		GoodsInfo DetailList `json:"goods_info,omitempty"` // 店铺信息
	}

	DetailList struct {
		DetailList []*DetailItem `json:"detail_list,omitempty"` // 是 商品信息
	}

	DetailItem struct {
		GoodsName   string `json:"goods_name,omitempty"`    // 是 商品名称
		GoodsImgURL string `json:"goods_img_url,omitempty"` // 是 商品图片url
		GoodsDesc   string `json:"goods_desc,omitempty"`    // 否 商品详情描述，不传默认取“商品名称”值，最多40汉字
	}

	WaybillInfo struct {
		Status    int32  `json:"status,omitempty"`     // 是 运单状态，见运单状态
		WaybillID string `json:"waybill_id,omitempty"` // 是 运单号
	}

	DeliveryInfo struct {
		DeliveryID   string `json:"delivery_id,omitempty"`   // 是 运力公司 id
		DeliveryName string `json:"delivery_name,omitempty"` // 否 运力公司名称
	}
)

func NewTrace() *Trace {
	return &Trace{}
}

func (t *Trace) Waybill(args *TraceRequest) (*TraceResponse, error) {
	token, err := NewToken().Access()
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}

	request := &netutil.HttpRequest{
		RawURL: fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/express/delivery/open_msg/trace_waybill?access_token=%s", token),
		Method: "POST",
		Body:   body,
		Headers: http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		},
	}

	client := netutil.NewHttpClient()
	resp, err := client.SendRequest(request)
	if err != nil {
		return nil, err
	}

	result := TraceResponse{}
	err = client.DecodeResponse(resp, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, err
}

// Query query trace
func (t *Trace) Query(waybillToken string) (*TraceResponse, error) {
	token, err := NewToken().Access()
	if err != nil {
		return nil, err
	}

	body, err := json.Marshal(map[string]string{
		"waybill_token": waybillToken,
	})
	if err != nil {
		return nil, err
	}

	request := &netutil.HttpRequest{
		RawURL: fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/express/delivery/open_msg/query_trace?access_token=%s", token),
		Method: "POST",
		Body:   body,
		Headers: http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		},
	}

	client := netutil.NewHttpClient()
	resp, err := client.SendRequest(request)
	if err != nil {
		return nil, err
	}

	result := TraceResponse{}
	err = client.DecodeResponse(resp, &result)
	if err != nil {
		return nil, err
	}

	if result.ErrCode != 0 {
		return nil, errors.New(result.ErrMsg)
	}
	return &result, err
}
