package order_getSettleBillDetail_request

import (
	"doudian.com/open/sdk_golang/api/order_getSettleBillDetail/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderGetSettleBillDetailRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderGetSettleBillDetailParam 
}
func (c *OrderGetSettleBillDetailRequest) GetUrlPath() string{
	return "/order/getSettleBillDetail"
}


func New() *OrderGetSettleBillDetailRequest{
	request := &OrderGetSettleBillDetailRequest{
		Param: &OrderGetSettleBillDetailParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderGetSettleBillDetailRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_getSettleBillDetail_response.OrderGetSettleBillDetailResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_getSettleBillDetail_response.OrderGetSettleBillDetailResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderGetSettleBillDetailRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderGetSettleBillDetailRequest) GetParams() *OrderGetSettleBillDetailParam{
	return c.Param
}


type OrderGetSettleBillDetailParam struct {
	// 页数（默认为0，第一页从0开始）
	Page *int64 `json:"page,omitempty"`
	// 每页结果数,默认为20
	Size *int64 `json:"size,omitempty"`
	// 查询开始时间
	StartTime *string `json:"start_time,omitempty"`
	// 查询结束时间，建议开始时间和结束时间间隔不超过7天
	EndTime *string `json:"end_time,omitempty"`
	// 子订单id
	OrderId string `json:"order_id,omitempty"`
	// 商品id
	ProductId string `json:"product_id,omitempty"`
	// 结算账户 0:全部 1:微信（升级前） 2:微信 3:支付宝 4:合众支付 5:聚合账户
	PayType string `json:"pay_type,omitempty"`
	// 业务类型，不传则默认为0 0:全部 1:鲁班广告, 2:值点商城, 3:精选联盟  4:小店自卖
	FlowType string `json:"flow_type,omitempty"`
	// 时间类型 0:结算时间 1：下单时间
	TimeType *string `json:"time_type,omitempty"`
}
