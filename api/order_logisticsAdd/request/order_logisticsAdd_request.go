package order_logisticsAdd_request

import (
	"doudian.com/open/sdk_golang/api/order_logisticsAdd/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderLogisticsAddRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderLogisticsAddParam 
}
func (c *OrderLogisticsAddRequest) GetUrlPath() string{
	return "/order/logisticsAdd"
}


func New() *OrderLogisticsAddRequest{
	request := &OrderLogisticsAddRequest{
		Param: &OrderLogisticsAddParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderLogisticsAddRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_logisticsAdd_response.OrderLogisticsAddResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_logisticsAdd_response.OrderLogisticsAddResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderLogisticsAddRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderLogisticsAddRequest) GetParams() *OrderLogisticsAddParam{
	return c.Param
}


type OrderLogisticsAddParam struct {
	// 订单ID
	OrderId *string `json:"order_id,omitempty"`
	// 【已废弃】物流公司ID。请使用company_code字段。
	LogisticsId int64 `json:"logistics_id,omitempty"`
	// 物流公司名称
	Company string `json:"company,omitempty"`
	// (必填)物流公司code,可从/order/logisticsCompanyList接口获取。
	CompanyCode string `json:"company_code,omitempty"`
	// (必填)快递单号
	LogisticsCode string `json:"logistics_code,omitempty"`
	// 是否拒绝退款申请（true表示拒绝退款，并继续发货；不传或为false表示有退款需要处理，拒绝发货），is_refund_reject和is_reject_refund随机使用一个即可
	IsRefundReject bool `json:"is_refund_reject,omitempty"`
	// 是否拒绝退款申请（true表示拒绝退款，并继续发货；不传或为false表示有退款需要处理，拒绝发货），is_refund_reject和is_reject_refund随机使用一个即可
	IsRejectRefund bool `json:"is_reject_refund,omitempty"`
	// 商品序列号，15-17位数字
	SerialNumberList []string `json:"serial_number_list,omitempty"`
	// 发货地址ID,通过地址库列表【/address/list】接口查询
	AddressId int64 `json:"address_id,omitempty"`
	// 门店ID；抖超小时达店铺需要填写；
	StoreId int64 `json:"store_id,omitempty"`
	// 退货地址ID,通过地址库列表【/address/list】接口查询
	AfterSaleAddressId int64 `json:"after_sale_address_id,omitempty"`
}
