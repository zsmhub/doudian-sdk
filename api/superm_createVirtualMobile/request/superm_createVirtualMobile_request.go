package superm_createVirtualMobile_request

import (
	"doudian.com/open/sdk_golang/api/superm_createVirtualMobile/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SupermCreateVirtualMobileRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SupermCreateVirtualMobileParam 
}
func (c *SupermCreateVirtualMobileRequest) GetUrlPath() string{
	return "/superm/createVirtualMobile"
}


func New() *SupermCreateVirtualMobileRequest{
	request := &SupermCreateVirtualMobileRequest{
		Param: &SupermCreateVirtualMobileParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SupermCreateVirtualMobileRequest) Execute(accessToken *doudian_sdk.AccessToken) (*superm_createVirtualMobile_response.SupermCreateVirtualMobileResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &superm_createVirtualMobile_response.SupermCreateVirtualMobileResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SupermCreateVirtualMobileRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SupermCreateVirtualMobileRequest) GetParams() *SupermCreateVirtualMobileParam{
	return c.Param
}


type SupermCreateVirtualMobileParam struct {
	// 店铺订单编号
	ShopOrderId *int64 `json:"shop_order_id,omitempty"`
}
