package order_logisticsEdit_request

import (
	"doudian.com/open/sdk_golang/api/order_logisticsEdit/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderLogisticsEditRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderLogisticsEditParam 
}
func (c *OrderLogisticsEditRequest) GetUrlPath() string{
	return "/order/logisticsEdit"
}


func New() *OrderLogisticsEditRequest{
	request := &OrderLogisticsEditRequest{
		Param: &OrderLogisticsEditParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderLogisticsEditRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_logisticsEdit_response.OrderLogisticsEditResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_logisticsEdit_response.OrderLogisticsEditResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderLogisticsEditRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderLogisticsEditRequest) GetParams() *OrderLogisticsEditParam{
	return c.Param
}


type OrderLogisticsEditParam struct {
	// 订单ID
	OrderId *string `json:"order_id,omitempty"`
	// 已废弃。物流公司ID。请使用company_code字段。
	LogisticsId int64 `json:"logistics_id,omitempty"`
	// 物流公司code,由接口/order/logisticsCompanyList返回的物流公司列表获得，必填
	CompanyCode string `json:"company_code,omitempty"`
	// 快递单号
	LogisticsCode string `json:"logistics_code,omitempty"`
	// 已废弃。物流公司名称
	Company string `json:"company,omitempty"`
	// 门店ID
	StoreId int64 `json:"store_id,omitempty"`
	// 退货地址id，通过地址库列表【/address/list】接口查询
	AfterSaleAddressId int64 `json:"after_sale_address_id,omitempty"`
}
