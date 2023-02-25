package superm_getPlatformPickUpCalendar_request

import (
	"doudian.com/open/sdk_golang/api/superm_getPlatformPickUpCalendar/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SupermGetPlatformPickUpCalendarRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SupermGetPlatformPickUpCalendarParam 
}
func (c *SupermGetPlatformPickUpCalendarRequest) GetUrlPath() string{
	return "/superm/getPlatformPickUpCalendar"
}


func New() *SupermGetPlatformPickUpCalendarRequest{
	request := &SupermGetPlatformPickUpCalendarRequest{
		Param: &SupermGetPlatformPickUpCalendarParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SupermGetPlatformPickUpCalendarRequest) Execute(accessToken *doudian_sdk.AccessToken) (*superm_getPlatformPickUpCalendar_response.SupermGetPlatformPickUpCalendarResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &superm_getPlatformPickUpCalendar_response.SupermGetPlatformPickUpCalendarResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SupermGetPlatformPickUpCalendarRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SupermGetPlatformPickUpCalendarRequest) GetParams() *SupermGetPlatformPickUpCalendarParam{
	return c.Param
}


type SupermGetPlatformPickUpCalendarParam struct {
	// 售后单号
	AftersaleId *int64 `json:"aftersale_id,omitempty"`
}
