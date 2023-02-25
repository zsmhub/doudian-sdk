package order_orderDetail_request

import (
	"doudian.com/open/sdk_golang/api/order_orderDetail/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderOrderDetailRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderOrderDetailParam 
}
func (c *OrderOrderDetailRequest) GetUrlPath() string{
	return "/order/orderDetail"
}


func New() *OrderOrderDetailRequest{
	request := &OrderOrderDetailRequest{
		Param: &OrderOrderDetailParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderOrderDetailRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_orderDetail_response.OrderOrderDetailResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_orderDetail_response.OrderOrderDetailResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderOrderDetailRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderOrderDetailRequest) GetParams() *OrderOrderDetailParam{
	return c.Param
}


type OrderOrderDetailParam struct {
	// 店铺父订单号，抖店平台生成，平台下唯一；
	ShopOrderId *string `json:"shop_order_id,omitempty"`
	IsSearchable bool `json:"is_searchable,omitempty,omitempty"`
}
