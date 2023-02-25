package order_addSerialNumber_request

import (
	"doudian.com/open/sdk_golang/api/order_addSerialNumber/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderAddSerialNumberRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderAddSerialNumberParam 
}
func (c *OrderAddSerialNumberRequest) GetUrlPath() string{
	return "/order/addSerialNumber"
}


func New() *OrderAddSerialNumberRequest{
	request := &OrderAddSerialNumberRequest{
		Param: &OrderAddSerialNumberParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderAddSerialNumberRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_addSerialNumber_response.OrderAddSerialNumberResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_addSerialNumber_response.OrderAddSerialNumberResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderAddSerialNumberRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderAddSerialNumberRequest) GetParams() *OrderAddSerialNumberParam{
	return c.Param
}


type OrderAddSerialNumberParam struct {
	// 订单号
	OrderId *string `json:"order_id,omitempty"`
	// 商品序列号，序列号长度不能超过30位字符，其中手机序列号仅支持填写15～17位数字
	SerialNumberList *[]string `json:"serial_number_list,omitempty"`
}
