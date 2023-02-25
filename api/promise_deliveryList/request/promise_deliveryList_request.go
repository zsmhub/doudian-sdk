package promise_deliveryList_request

import (
	"doudian.com/open/sdk_golang/api/promise_deliveryList/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type PromiseDeliveryListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *PromiseDeliveryListParam 
}
func (c *PromiseDeliveryListRequest) GetUrlPath() string{
	return "/promise/deliveryList"
}


func New() *PromiseDeliveryListRequest{
	request := &PromiseDeliveryListRequest{
		Param: &PromiseDeliveryListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *PromiseDeliveryListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*promise_deliveryList_response.PromiseDeliveryListResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &promise_deliveryList_response.PromiseDeliveryListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *PromiseDeliveryListRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *PromiseDeliveryListRequest) GetParams() *PromiseDeliveryListParam{
	return c.Param
}


type PromiseDeliveryListParam struct {
	// 商品ID
	ProductId *int64 `json:"product_id,omitempty"`
	// 页码，从1开始
	Page *int64 `json:"page,omitempty"`
	// 每页数量
	Size *int64 `json:"size,omitempty"`
}
