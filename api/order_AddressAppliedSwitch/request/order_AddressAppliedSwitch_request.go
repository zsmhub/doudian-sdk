package order_AddressAppliedSwitch_request

import (
	"doudian.com/open/sdk_golang/api/order_AddressAppliedSwitch/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderAddressAppliedSwitchRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderAddressAppliedSwitchParam 
}
func (c *OrderAddressAppliedSwitchRequest) GetUrlPath() string{
	return "/order/AddressAppliedSwitch"
}


func New() *OrderAddressAppliedSwitchRequest{
	request := &OrderAddressAppliedSwitchRequest{
		Param: &OrderAddressAppliedSwitchParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderAddressAppliedSwitchRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_AddressAppliedSwitch_response.OrderAddressAppliedSwitchResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_AddressAppliedSwitch_response.OrderAddressAppliedSwitchResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderAddressAppliedSwitchRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderAddressAppliedSwitchRequest) GetParams() *OrderAddressAppliedSwitchParam{
	return c.Param
}


type OrderAddressAppliedSwitchParam struct {
	// 0代表关闭，不需要审核 1代表开启审核，买家变更需要审核
	IsAllowed *int64 `json:"is_allowed,omitempty"`
	// 0代表异步审核，1表示实时审核
	ReviewType int64 `json:"review_type,omitempty"`
}
