package order_addressConfirm_request

import (
	"doudian.com/open/sdk_golang/api/order_addressConfirm/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderAddressConfirmRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderAddressConfirmParam 
}
func (c *OrderAddressConfirmRequest) GetUrlPath() string{
	return "/order/addressConfirm"
}


func New() *OrderAddressConfirmRequest{
	request := &OrderAddressConfirmRequest{
		Param: &OrderAddressConfirmParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderAddressConfirmRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_addressConfirm_response.OrderAddressConfirmResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_addressConfirm_response.OrderAddressConfirmResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderAddressConfirmRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderAddressConfirmRequest) GetParams() *OrderAddressConfirmParam{
	return c.Param
}


type OrderAddressConfirmParam struct {
	// 主订单id，注意请求时order/list或 order/detail中返回的主订单id带的‘A’需要截断掉
	OrderId *string `json:"order_id,omitempty"`
	// 0:同意; 拒绝需要传入以下参数： 1001:订单已进入拣货环节 1002:订单已进入配货环节 1003:订单已进入仓库环节 1004:订单已进入出库环节 1005:订单已进入发货环节
	IsApproved *int64 `json:"is_approved,omitempty"`
}
