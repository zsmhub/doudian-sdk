package logistics_cancelOrder_request

import (
	"doudian.com/open/sdk_golang/api/logistics_cancelOrder/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type LogisticsCancelOrderRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *LogisticsCancelOrderParam 
}
func (c *LogisticsCancelOrderRequest) GetUrlPath() string{
	return "/logistics/cancelOrder"
}


func New() *LogisticsCancelOrderRequest{
	request := &LogisticsCancelOrderRequest{
		Param: &LogisticsCancelOrderParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *LogisticsCancelOrderRequest) Execute(accessToken *doudian_sdk.AccessToken) (*logistics_cancelOrder_response.LogisticsCancelOrderResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &logistics_cancelOrder_response.LogisticsCancelOrderResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *LogisticsCancelOrderRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *LogisticsCancelOrderRequest) GetParams() *LogisticsCancelOrderParam{
	return c.Param
}


type LogisticsCancelOrderParam struct {
	// 物流公司
	LogisticsCode *string `json:"logistics_code,omitempty"`
	// 运单号
	TrackNo *string `json:"track_no,omitempty"`
	// 实际使用取号服务店铺user_id
	UserId int64 `json:"user_id,omitempty"`
}
