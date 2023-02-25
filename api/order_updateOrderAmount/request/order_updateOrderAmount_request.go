package order_updateOrderAmount_request

import (
	"doudian.com/open/sdk_golang/api/order_updateOrderAmount/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderUpdateOrderAmountRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderUpdateOrderAmountParam 
}
func (c *OrderUpdateOrderAmountRequest) GetUrlPath() string{
	return "/order/updateOrderAmount"
}


func New() *OrderUpdateOrderAmountRequest{
	request := &OrderUpdateOrderAmountRequest{
		Param: &OrderUpdateOrderAmountParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderUpdateOrderAmountRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_updateOrderAmount_response.OrderUpdateOrderAmountResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_updateOrderAmount_response.OrderUpdateOrderAmountResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderUpdateOrderAmountRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderUpdateOrderAmountRequest) GetParams() *OrderUpdateOrderAmountParam{
	return c.Param
}


type UpdateDetailItem struct {
	OrderId string `json:"order_id,omitempty"`
	DiscountAmount int64 `json:"discount_amount,omitempty"`
}
type OrderUpdateOrderAmountParam struct {
	Pid *string `json:"pid,omitempty"`
	UpdateDetail []UpdateDetailItem `json:"update_detail,omitempty"`
}
