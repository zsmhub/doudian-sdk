package superm_getDispatcherInfo_request

import (
	"doudian.com/open/sdk_golang/api/superm_getDispatcherInfo/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SupermGetDispatcherInfoRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SupermGetDispatcherInfoParam 
}
func (c *SupermGetDispatcherInfoRequest) GetUrlPath() string{
	return "/superm/getDispatcherInfo"
}


func New() *SupermGetDispatcherInfoRequest{
	request := &SupermGetDispatcherInfoRequest{
		Param: &SupermGetDispatcherInfoParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SupermGetDispatcherInfoRequest) Execute(accessToken *doudian_sdk.AccessToken) (*superm_getDispatcherInfo_response.SupermGetDispatcherInfoResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &superm_getDispatcherInfo_response.SupermGetDispatcherInfoResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SupermGetDispatcherInfoRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SupermGetDispatcherInfoRequest) GetParams() *SupermGetDispatcherInfoParam{
	return c.Param
}


type SupermGetDispatcherInfoParam struct {
	// 门店ID
	StoreId *int64 `json:"store_id,omitempty"`
	// 订单ID
	ShopOrderId *string `json:"shop_order_id,omitempty"`
	// 运力费用类型 1: 呼叫运力费用，2:取消运力费用
	DispatcherFeeType *int32 `json:"dispatcher_fee_type,omitempty"`
}
