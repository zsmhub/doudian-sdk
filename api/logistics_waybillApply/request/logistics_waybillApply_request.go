package logistics_waybillApply_request

import (
	"doudian.com/open/sdk_golang/api/logistics_waybillApply/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type LogisticsWaybillApplyRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *LogisticsWaybillApplyParam 
}
func (c *LogisticsWaybillApplyRequest) GetUrlPath() string{
	return "/logistics/waybillApply"
}


func New() *LogisticsWaybillApplyRequest{
	request := &LogisticsWaybillApplyRequest{
		Param: &LogisticsWaybillApplyParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *LogisticsWaybillApplyRequest) Execute(accessToken *doudian_sdk.AccessToken) (*logistics_waybillApply_response.LogisticsWaybillApplyResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &logistics_waybillApply_response.LogisticsWaybillApplyResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *LogisticsWaybillApplyRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *LogisticsWaybillApplyRequest) GetParams() *LogisticsWaybillApplyParam{
	return c.Param
}


type WaybillAppliesItem struct {
	// 物流公司编码
	LogisticsCode *string `json:"logistics_code,omitempty"`
	// 运单号
	TrackNo *string `json:"track_no,omitempty"`
}
type LogisticsWaybillApplyParam struct {
	// 请求结构体
	WaybillApplies *[]WaybillAppliesItem `json:"waybill_applies,omitempty"`
}
