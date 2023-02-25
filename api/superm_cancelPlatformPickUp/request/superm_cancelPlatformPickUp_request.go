package superm_cancelPlatformPickUp_request

import (
	"doudian.com/open/sdk_golang/api/superm_cancelPlatformPickUp/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SupermCancelPlatformPickUpRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SupermCancelPlatformPickUpParam 
}
func (c *SupermCancelPlatformPickUpRequest) GetUrlPath() string{
	return "/superm/cancelPlatformPickUp"
}


func New() *SupermCancelPlatformPickUpRequest{
	request := &SupermCancelPlatformPickUpRequest{
		Param: &SupermCancelPlatformPickUpParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SupermCancelPlatformPickUpRequest) Execute(accessToken *doudian_sdk.AccessToken) (*superm_cancelPlatformPickUp_response.SupermCancelPlatformPickUpResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &superm_cancelPlatformPickUp_response.SupermCancelPlatformPickUpResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SupermCancelPlatformPickUpRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SupermCancelPlatformPickUpRequest) GetParams() *SupermCancelPlatformPickUpParam{
	return c.Param
}


type SupermCancelPlatformPickUpParam struct {
	// 售后单ID
	AftersaleId *int64 `json:"aftersale_id,omitempty"`
	// 取消运力的原因，开发者自定义传入，最大支持50个字符；填写后商家可以在抖店后台查看到。
	CancelReason string `json:"cancel_reason,omitempty"`
}
