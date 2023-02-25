package superm_getPlatformPickUpEstimatedCharge_request

import (
	"doudian.com/open/sdk_golang/api/superm_getPlatformPickUpEstimatedCharge/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SupermGetPlatformPickUpEstimatedChargeRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SupermGetPlatformPickUpEstimatedChargeParam 
}
func (c *SupermGetPlatformPickUpEstimatedChargeRequest) GetUrlPath() string{
	return "/superm/getPlatformPickUpEstimatedCharge"
}


func New() *SupermGetPlatformPickUpEstimatedChargeRequest{
	request := &SupermGetPlatformPickUpEstimatedChargeRequest{
		Param: &SupermGetPlatformPickUpEstimatedChargeParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SupermGetPlatformPickUpEstimatedChargeRequest) Execute(accessToken *doudian_sdk.AccessToken) (*superm_getPlatformPickUpEstimatedCharge_response.SupermGetPlatformPickUpEstimatedChargeResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &superm_getPlatformPickUpEstimatedCharge_response.SupermGetPlatformPickUpEstimatedChargeResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SupermGetPlatformPickUpEstimatedChargeRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SupermGetPlatformPickUpEstimatedChargeRequest) GetParams() *SupermGetPlatformPickUpEstimatedChargeParam{
	return c.Param
}


type SelectedCalendarPeriod struct {
	// 时间段开始时间；Unix时间戳，单位：秒
	TimeBeginTs *int64 `json:"time_begin_ts,omitempty"`
	// 时间段结束时间；Unix时间戳，单位：秒
	TimeEndTs *int64 `json:"time_end_ts,omitempty"`
}
type SupermGetPlatformPickUpEstimatedChargeParam struct {
	// 售后单ID
	AftersaleId *int64 `json:"aftersale_id,omitempty"`
	// 查询的时间段
	SelectedCalendarPeriod *SelectedCalendarPeriod `json:"selected_calendar_period,omitempty"`
	// 查询的费用类型；1 -呼叫运力费用；2-取消运力费用
	DispatcherFeeType *int32 `json:"dispatcher_fee_type,omitempty"`
}
