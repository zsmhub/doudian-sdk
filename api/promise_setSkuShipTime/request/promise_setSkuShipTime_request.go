package promise_setSkuShipTime_request

import (
	"doudian.com/open/sdk_golang/api/promise_setSkuShipTime/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type PromiseSetSkuShipTimeRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *PromiseSetSkuShipTimeParam 
}
func (c *PromiseSetSkuShipTimeRequest) GetUrlPath() string{
	return "/promise/setSkuShipTime"
}


func New() *PromiseSetSkuShipTimeRequest{
	request := &PromiseSetSkuShipTimeRequest{
		Param: &PromiseSetSkuShipTimeParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *PromiseSetSkuShipTimeRequest) Execute(accessToken *doudian_sdk.AccessToken) (*promise_setSkuShipTime_response.PromiseSetSkuShipTimeResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &promise_setSkuShipTime_response.PromiseSetSkuShipTimeResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *PromiseSetSkuShipTimeRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *PromiseSetSkuShipTimeRequest) GetParams() *PromiseSetSkuShipTimeParam{
	return c.Param
}


type RulesItem struct {
	// skuid
	SkuId *string `json:"sku_id,omitempty"`
	// 外部仓库id
	OutWarehouseId *string `json:"out_warehouse_id,omitempty"`
	// 0表示现货模式，1表示全款预售模式
	PreSellType *int16 `json:"pre_sell_type,omitempty"`
	// 发货延迟时间：0表示当天发货，1表示24小时发货；当全款预售时，需传入值为(2,15)，即2-15中的任意值
	DelayDay *int16 `json:"delay_day,omitempty"`
	// 全款预售截止时间，和PreSellType=1时组合使用
	PreSellEndTime *int64 `json:"pre_sell_end_time,omitempty"`
}
type PromiseSetSkuShipTimeParam struct {
	// sku发货时效规则
	Rules *[]RulesItem `json:"rules,omitempty"`
}
