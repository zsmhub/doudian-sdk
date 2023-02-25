package order_serviceDetail_request

import (
	"doudian.com/open/sdk_golang/api/order_serviceDetail/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderServiceDetailRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderServiceDetailParam 
}
func (c *OrderServiceDetailRequest) GetUrlPath() string{
	return "/order/serviceDetail"
}


func New() *OrderServiceDetailRequest{
	request := &OrderServiceDetailRequest{
		Param: &OrderServiceDetailParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderServiceDetailRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_serviceDetail_response.OrderServiceDetailResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_serviceDetail_response.OrderServiceDetailResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderServiceDetailRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderServiceDetailRequest) GetParams() *OrderServiceDetailParam{
	return c.Param
}


type OrderServiceDetailParam struct {
	// 服务请求列表中获取的id
	ServiceId *int64 `json:"service_id,omitempty"`
}
