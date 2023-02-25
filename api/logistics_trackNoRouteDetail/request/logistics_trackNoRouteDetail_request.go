package logistics_trackNoRouteDetail_request

import (
	"doudian.com/open/sdk_golang/api/logistics_trackNoRouteDetail/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type LogisticsTrackNoRouteDetailRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *LogisticsTrackNoRouteDetailParam 
}
func (c *LogisticsTrackNoRouteDetailRequest) GetUrlPath() string{
	return "/logistics/trackNoRouteDetail"
}


func New() *LogisticsTrackNoRouteDetailRequest{
	request := &LogisticsTrackNoRouteDetailRequest{
		Param: &LogisticsTrackNoRouteDetailParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *LogisticsTrackNoRouteDetailRequest) Execute(accessToken *doudian_sdk.AccessToken) (*logistics_trackNoRouteDetail_response.LogisticsTrackNoRouteDetailResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &logistics_trackNoRouteDetail_response.LogisticsTrackNoRouteDetailResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *LogisticsTrackNoRouteDetailRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *LogisticsTrackNoRouteDetailRequest) GetParams() *LogisticsTrackNoRouteDetailParam{
	return c.Param
}


type LogisticsTrackNoRouteDetailParam struct {
	// 物流商编码；需使用【/order/logisticsCompanyList】接口响应参数中的code；
	LogisticsCode *string `json:"logistics_code,omitempty"`
	// 运单号；可使用电子面单接口获取返回的单号查询【/logistics/newCreateOrder】或商家店铺后台查看
	TrackNo *string `json:"track_no,omitempty"`
}
