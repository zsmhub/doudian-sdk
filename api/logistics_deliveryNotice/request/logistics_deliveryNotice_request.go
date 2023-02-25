package logistics_deliveryNotice_request

import (
	"doudian.com/open/sdk_golang/api/logistics_deliveryNotice/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type LogisticsDeliveryNoticeRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *LogisticsDeliveryNoticeParam 
}
func (c *LogisticsDeliveryNoticeRequest) GetUrlPath() string{
	return "/logistics/deliveryNotice"
}


func New() *LogisticsDeliveryNoticeRequest{
	request := &LogisticsDeliveryNoticeRequest{
		Param: &LogisticsDeliveryNoticeParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *LogisticsDeliveryNoticeRequest) Execute(accessToken *doudian_sdk.AccessToken) (*logistics_deliveryNotice_response.LogisticsDeliveryNoticeResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &logistics_deliveryNotice_response.LogisticsDeliveryNoticeResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *LogisticsDeliveryNoticeRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *LogisticsDeliveryNoticeRequest) GetParams() *LogisticsDeliveryNoticeParam{
	return c.Param
}


type LogisticsDeliveryNoticeParam struct {
	// 运单号
	WaybillNo *string `json:"waybill_no,omitempty"`
	// 放行/回退
	NoticeType *string `json:"notice_type,omitempty"`
}
