package order_logisticsCompanyList_request

import (
	"doudian.com/open/sdk_golang/api/order_logisticsCompanyList/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderLogisticsCompanyListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderLogisticsCompanyListParam 
}
func (c *OrderLogisticsCompanyListRequest) GetUrlPath() string{
	return "/order/logisticsCompanyList"
}


func New() *OrderLogisticsCompanyListRequest{
	request := &OrderLogisticsCompanyListRequest{
		Param: &OrderLogisticsCompanyListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderLogisticsCompanyListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_logisticsCompanyList_response.OrderLogisticsCompanyListResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_logisticsCompanyList_response.OrderLogisticsCompanyListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderLogisticsCompanyListRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderLogisticsCompanyListRequest) GetParams() *OrderLogisticsCompanyListParam{
	return c.Param
}


type OrderLogisticsCompanyListParam struct {
}
