package order_addresSwitchConfig_request

import (
	"doudian.com/open/sdk_golang/api/order_addresSwitchConfig/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderAddresSwitchConfigRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderAddresSwitchConfigParam 
}
func (c *OrderAddresSwitchConfigRequest) GetUrlPath() string{
	return "/order/addresSwitchConfig"
}


func New() *OrderAddresSwitchConfigRequest{
	request := &OrderAddresSwitchConfigRequest{
		Param: &OrderAddresSwitchConfigParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderAddresSwitchConfigRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_addresSwitchConfig_response.OrderAddresSwitchConfigResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_addresSwitchConfig_response.OrderAddresSwitchConfigResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderAddresSwitchConfigRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderAddresSwitchConfigRequest) GetParams() *OrderAddresSwitchConfigParam{
	return c.Param
}


type OrderAddresSwitchConfigParam struct {
}
