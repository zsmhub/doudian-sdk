package order_getSettleBillDetailV3_request

import (
	"doudian.com/open/sdk_golang/api/order_getSettleBillDetailV3/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderGetSettleBillDetailV3Request struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderGetSettleBillDetailV3Param 
}
func (c *OrderGetSettleBillDetailV3Request) GetUrlPath() string{
	return "/order/getSettleBillDetailV3"
}


func New() *OrderGetSettleBillDetailV3Request{
	request := &OrderGetSettleBillDetailV3Request{
		Param: &OrderGetSettleBillDetailV3Param{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderGetSettleBillDetailV3Request) Execute(accessToken *doudian_sdk.AccessToken) (*order_getSettleBillDetailV3_response.OrderGetSettleBillDetailV3Response, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_getSettleBillDetailV3_response.OrderGetSettleBillDetailV3Response{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderGetSettleBillDetailV3Request) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderGetSettleBillDetailV3Request) GetParams() *OrderGetSettleBillDetailV3Param{
	return c.Param
}


type OrderGetSettleBillDetailV3Param struct {
	// 页数，支持范围1~100
	Size int64 `json:"size,omitempty"`
	// 查询开始时间，格式为：yyyy-MM-dd HH:mm:ss，订单号未传的情况下，开始时间必须传，注意：分页查询时，除首次查询外，应填入上一次返回的next_start_time
	StartTime string `json:"start_time,omitempty"`
	// 查询结束时间，和end_time的时间间隔建议不超过7天，格式为：yyyy-MM-dd HH:mm:ss，订单号未传的情况下，结束时间必须传
	EndTime string `json:"end_time,omitempty"`
	// SKU单，子订单号，支持通过英文逗号分隔传入多个参数
	OrderId string `json:"order_id,omitempty"`
	// 商品id
	ProductId string `json:"product_id,omitempty"`
	// 结算账户，不传则默认为全部，枚举： 1（微信：升级前）、 2（微信）、 3（支付宝）、 4（合众支付）、 5（聚合账户），支持通过英文逗号分隔传入多个参数
	PayType string `json:"pay_type,omitempty"`
	// 业务类型，不传则默认为全部，枚举： 1（鲁班广告）、2（值点商城）, 3（精选联盟）、 4（小店自卖）
	FlowType string `json:"flow_type,omitempty"`
	// 时间类型 ，不传则默认为结算时间，枚举： 0（结算时间） 1（下单时间）
	TimeType string `json:"time_type,omitempty"`
	// 查询开始索引，注意：分页查询时，除首次查询可不填外，应填入上一次返回的next_start_index
	StartIndex string `json:"start_index,omitempty"`
}
