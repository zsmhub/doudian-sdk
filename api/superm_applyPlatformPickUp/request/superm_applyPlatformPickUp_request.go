package superm_applyPlatformPickUp_request

import (
	"doudian.com/open/sdk_golang/api/superm_applyPlatformPickUp/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SupermApplyPlatformPickUpRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SupermApplyPlatformPickUpParam 
}
func (c *SupermApplyPlatformPickUpRequest) GetUrlPath() string{
	return "/superm/applyPlatformPickUp"
}


func New() *SupermApplyPlatformPickUpRequest{
	request := &SupermApplyPlatformPickUpRequest{
		Param: &SupermApplyPlatformPickUpParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SupermApplyPlatformPickUpRequest) Execute(accessToken *doudian_sdk.AccessToken) (*superm_applyPlatformPickUp_response.SupermApplyPlatformPickUpResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &superm_applyPlatformPickUp_response.SupermApplyPlatformPickUpResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SupermApplyPlatformPickUpRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SupermApplyPlatformPickUpRequest) GetParams() *SupermApplyPlatformPickUpParam{
	return c.Param
}


type SupermApplyPlatformPickUpParam struct {
	// 售后单ID
	AftersaleId *int64 `json:"aftersale_id,omitempty"`
	// 运力呼叫时间段
	SelectedCalendarPeriod *SelectedCalendarPeriod `json:"selected_calendar_period,omitempty"`
}
type SelectedCalendarPeriod struct {
	// 时间段开始时间
	TimeBeginTs *int64 `json:"time_begin_ts,omitempty"`
	// 时间段结束时间
	TimeEndTs *int64 `json:"time_end_ts,omitempty"`
}
