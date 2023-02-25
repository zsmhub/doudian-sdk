package superm_shopOrderDispatcher_request

import (
	"doudian.com/open/sdk_golang/api/superm_shopOrderDispatcher/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SupermShopOrderDispatcherRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SupermShopOrderDispatcherParam 
}
func (c *SupermShopOrderDispatcherRequest) GetUrlPath() string{
	return "/superm/shopOrderDispatcher"
}


func New() *SupermShopOrderDispatcherRequest{
	request := &SupermShopOrderDispatcherRequest{
		Param: &SupermShopOrderDispatcherParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SupermShopOrderDispatcherRequest) Execute(accessToken *doudian_sdk.AccessToken) (*superm_shopOrderDispatcher_response.SupermShopOrderDispatcherResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &superm_shopOrderDispatcher_response.SupermShopOrderDispatcherResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SupermShopOrderDispatcherRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SupermShopOrderDispatcherRequest) GetParams() *SupermShopOrderDispatcherParam{
	return c.Param
}


type SupermShopOrderDispatcherParam struct {
	// 店铺ID
	StoreId *int64 `json:"store_id,omitempty"`
	// 订单号
	ShopOrderId *string `json:"shop_order_id,omitempty"`
	// enum DispatcherType {     Call   = 1 // 呼叫运力     Cancel = 2 // 取消运力 }
	DispatcherType *int32 `json:"dispatcher_type,omitempty"`
	// 序列码列表，如果包含有多个序列码，请用"_"英文下划线分割（属于数码手机类目的商品订单才需要传序列码。）
	SerialNumberList string `json:"serial_number_list,omitempty"`
	// 是否需要取件码(针对呼叫运力)0否1是
	VerifyCodeType int32 `json:"verify_code_type,omitempty"`
}
