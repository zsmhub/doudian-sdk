package order_downloadSettleItemToShop_request

import (
	"doudian.com/open/sdk_golang/api/order_downloadSettleItemToShop/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderDownloadSettleItemToShopRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderDownloadSettleItemToShopParam 
}
func (c *OrderDownloadSettleItemToShopRequest) GetUrlPath() string{
	return "/order/downloadSettleItemToShop"
}


func New() *OrderDownloadSettleItemToShopRequest{
	request := &OrderDownloadSettleItemToShopRequest{
		Param: &OrderDownloadSettleItemToShopParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderDownloadSettleItemToShopRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_downloadSettleItemToShop_response.OrderDownloadSettleItemToShopResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_downloadSettleItemToShop_response.OrderDownloadSettleItemToShopResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderDownloadSettleItemToShopRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderDownloadSettleItemToShopRequest) GetParams() *OrderDownloadSettleItemToShopParam{
	return c.Param
}


type OrderDownloadSettleItemToShopParam struct {
	// 开始时间
	StartTime *string `json:"start_time,omitempty"`
	// 结束时间
	EndTime *string `json:"end_time,omitempty"`
	// 时间类型 0:结算时间 1：下单时间
	TimeType *string `json:"time_type,omitempty"`
	// 订单ID
	OrderId string `json:"order_id,omitempty"`
	// 字段已作废,勿填
	BillId string `json:"bill_id,omitempty"`
	// 商品ID
	ProductId string `json:"product_id,omitempty"`
	// 结算账户 0:全部 1:微信（升级前） 2:微信 3:支付宝 4:合众支付 5:聚合账户
	PayType string `json:"pay_type,omitempty"`
	// 业务类型，不传则默认为0 0:全部 1:鲁班广告, 2:值点商城, 3:精选联盟  4:小店自卖
	FlowType string `json:"flow_type,omitempty"`
}
