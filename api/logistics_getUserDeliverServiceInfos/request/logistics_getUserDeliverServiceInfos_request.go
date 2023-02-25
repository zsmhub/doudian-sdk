package logistics_getUserDeliverServiceInfos_request

import (
	"doudian.com/open/sdk_golang/api/logistics_getUserDeliverServiceInfos/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type LogisticsGetUserDeliverServiceInfosRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *LogisticsGetUserDeliverServiceInfosParam 
}
func (c *LogisticsGetUserDeliverServiceInfosRequest) GetUrlPath() string{
	return "/logistics/getUserDeliverServiceInfos"
}


func New() *LogisticsGetUserDeliverServiceInfosRequest{
	request := &LogisticsGetUserDeliverServiceInfosRequest{
		Param: &LogisticsGetUserDeliverServiceInfosParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *LogisticsGetUserDeliverServiceInfosRequest) Execute(accessToken *doudian_sdk.AccessToken) (*logistics_getUserDeliverServiceInfos_response.LogisticsGetUserDeliverServiceInfosResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &logistics_getUserDeliverServiceInfos_response.LogisticsGetUserDeliverServiceInfosResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *LogisticsGetUserDeliverServiceInfosRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *LogisticsGetUserDeliverServiceInfosRequest) GetParams() *LogisticsGetUserDeliverServiceInfosParam{
	return c.Param
}


type LogisticsGetUserDeliverServiceInfosParam struct {
	// 商家信息
	BizInfo *BizInfo `json:"biz_info,omitempty"`
	// GetRecommendedAndDeliveryExpressByOrder：By单快递推荐服务
	ServiceCodes []string `json:"service_codes,omitempty"`
}
type BizInfo struct {
	// 1：抖店
	BizType *int32 `json:"biz_type,omitempty"`
}
