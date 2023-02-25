package order_updatePostAmount_request

import (
	"doudian.com/open/sdk_golang/api/order_updatePostAmount/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderUpdatePostAmountRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderUpdatePostAmountParam 
}
func (c *OrderUpdatePostAmountRequest) GetUrlPath() string{
	return "/order/updatePostAmount"
}


func New() *OrderUpdatePostAmountRequest{
	request := &OrderUpdatePostAmountRequest{
		Param: &OrderUpdatePostAmountParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderUpdatePostAmountRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_updatePostAmount_response.OrderUpdatePostAmountResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_updatePostAmount_response.OrderUpdatePostAmountResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderUpdatePostAmountRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderUpdatePostAmountRequest) GetParams() *OrderUpdatePostAmountParam{
	return c.Param
}


type OrderUpdatePostAmountParam struct {
	OrderId *string `json:"order_id,omitempty"`
	PostAmount *int64 `json:"post_amount,omitempty"`
	PostAmount *int64 `json:"post_amount,omitempty"`
}
