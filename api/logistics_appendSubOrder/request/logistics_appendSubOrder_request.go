package logistics_appendSubOrder_request

import (
	"doudian.com/open/sdk_golang/api/logistics_appendSubOrder/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type LogisticsAppendSubOrderRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *LogisticsAppendSubOrderParam 
}
func (c *LogisticsAppendSubOrderRequest) GetUrlPath() string{
	return "/logistics/appendSubOrder"
}


func New() *LogisticsAppendSubOrderRequest{
	request := &LogisticsAppendSubOrderRequest{
		Param: &LogisticsAppendSubOrderParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *LogisticsAppendSubOrderRequest) Execute(accessToken *doudian_sdk.AccessToken) (*logistics_appendSubOrder_response.LogisticsAppendSubOrderResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &logistics_appendSubOrder_response.LogisticsAppendSubOrderResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *LogisticsAppendSubOrderRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *LogisticsAppendSubOrderRequest) GetParams() *LogisticsAppendSubOrderParam{
	return c.Param
}


type LogisticsAppendSubOrderParam struct {
	// 运单号
	TrackNo *string `json:"track_no,omitempty"`
	// 物流商编码
	LogisticsCode *string `json:"logistics_code,omitempty"`
	// 追加个数
	PackAddQuantity *int32 `json:"pack_add_quantity,omitempty"`
	// 是否返回全量的子单号
	IsReturnFullSubCodes bool `json:"is_return_full_sub_codes,omitempty"`
}
