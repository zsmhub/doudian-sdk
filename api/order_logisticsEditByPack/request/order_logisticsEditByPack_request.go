package order_logisticsEditByPack_request

import (
	"doudian.com/open/sdk_golang/api/order_logisticsEditByPack/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderLogisticsEditByPackRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderLogisticsEditByPackParam 
}
func (c *OrderLogisticsEditByPackRequest) GetUrlPath() string{
	return "/order/logisticsEditByPack"
}


func New() *OrderLogisticsEditByPackRequest{
	request := &OrderLogisticsEditByPackRequest{
		Param: &OrderLogisticsEditByPackParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderLogisticsEditByPackRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_logisticsEditByPack_response.OrderLogisticsEditByPackResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_logisticsEditByPack_response.OrderLogisticsEditByPackResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderLogisticsEditByPackRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderLogisticsEditByPackRequest) GetParams() *OrderLogisticsEditByPackParam{
	return c.Param
}


type OrderLogisticsEditByPackParam struct {
	// 父订单ID，由orderList接口返回
	OrderId *string `json:"order_id,omitempty"`
	// 包裹ID
	PackId *string `json:"pack_id,omitempty"`
	// 运单号
	LogisticsCode *string `json:"logistics_code,omitempty"`
	// 物流公司ID，由接口/order/logisticsCompanyList返回的物流公司列表获得，必填
	CompanyCode string `json:"company_code,omitempty"`
	// 已废弃。物流公司ID。请使用company_code字段。
	LogisticsId string `json:"logistics_id,omitempty"`
	// 门店ID
	StoreId int64 `json:"store_id,omitempty"`
	// 退货地址id，通过地址库列表【/address/list】接口查询
	AfterSaleAddressId int64 `json:"after_sale_address_id,omitempty"`
}
