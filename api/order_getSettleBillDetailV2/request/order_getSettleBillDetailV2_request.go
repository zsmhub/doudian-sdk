package order_getSettleBillDetailV2_request

import (
	"doudian.com/open/sdk_golang/api/order_getSettleBillDetailV2/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderGetSettleBillDetailV2Request struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderGetSettleBillDetailV2Param 
}
func (c *OrderGetSettleBillDetailV2Request) GetUrlPath() string{
	return "/order/getSettleBillDetailV2"
}


func New() *OrderGetSettleBillDetailV2Request{
	request := &OrderGetSettleBillDetailV2Request{
		Param: &OrderGetSettleBillDetailV2Param{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderGetSettleBillDetailV2Request) Execute(accessToken *doudian_sdk.AccessToken) (*order_getSettleBillDetailV2_response.OrderGetSettleBillDetailV2Response, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_getSettleBillDetailV2_response.OrderGetSettleBillDetailV2Response{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderGetSettleBillDetailV2Request) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderGetSettleBillDetailV2Request) GetParams() *OrderGetSettleBillDetailV2Param{
	return c.Param
}


type OrderGetSettleBillDetailV2Param struct {
	// 查询页大小(默认100,最大为200)
	Size int64 `json:"size,omitempty"`
	// 查询开始时间。注:订单号未传的情况下，时间必须传
	StartTime string `json:"start_time,omitempty"`
	// 查询结束时间。和end_time的时间间隔建议不超过7天。注:订单号未传的情况下，时间必须传
	EndTime string `json:"end_time,omitempty"`
	// SKU单，子订单号
	OrderId string `json:"order_id,omitempty"`
	// 商品id
	ProductId string `json:"product_id,omitempty"`
	// 结算账户 0:全部 1:微信（升级前） 2:微信 3:支付宝 4:合众支付 5:聚合账户
	PayType string `json:"pay_type,omitempty"`
	// 业务类型，不传则默认为0 0:全部 1:鲁班广告, 2:精选联盟 ,3:值点商城,  4:小店自卖
	FlowType string `json:"flow_type,omitempty"`
	// 时间类型 0:结算时间 1：下单时间
	TimeType string `json:"time_type,omitempty"`
	// 查询开始索引,值为上一次请求的结果next_start_index,第一次查询可以不填
	StartIndex string `json:"start_index,omitempty"`
}
