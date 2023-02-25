package order_insurance_request

import (
	"doudian.com/open/sdk_golang/api/order_insurance/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderInsuranceRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderInsuranceParam 
}
func (c *OrderInsuranceRequest) GetUrlPath() string{
	return "/order/insurance"
}


func New() *OrderInsuranceRequest{
	request := &OrderInsuranceRequest{
		Param: &OrderInsuranceParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderInsuranceRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_insurance_response.OrderInsuranceResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_insurance_response.OrderInsuranceResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderInsuranceRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderInsuranceRequest) GetParams() *OrderInsuranceParam{
	return c.Param
}


type OrderInsuranceParam struct {
	// 订单id(可以加A也可以不加）
	OrderId *string `json:"order_id,omitempty"`
}
