package order_logisticsAddMultiPack_request

import (
	"doudian.com/open/sdk_golang/api/order_logisticsAddMultiPack/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderLogisticsAddMultiPackRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderLogisticsAddMultiPackParam 
}
func (c *OrderLogisticsAddMultiPackRequest) GetUrlPath() string{
	return "/order/logisticsAddMultiPack"
}


func New() *OrderLogisticsAddMultiPackRequest{
	request := &OrderLogisticsAddMultiPackRequest{
		Param: &OrderLogisticsAddMultiPackParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderLogisticsAddMultiPackRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_logisticsAddMultiPack_response.OrderLogisticsAddMultiPackResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_logisticsAddMultiPack_response.OrderLogisticsAddMultiPackResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderLogisticsAddMultiPackRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderLogisticsAddMultiPackRequest) GetParams() *OrderLogisticsAddMultiPackParam{
	return c.Param
}


type BundleListItem struct {
	// 组套商品子商品ID
	SubProductId *string `json:"sub_product_id,omitempty"`
	// 组套商品子商品SKU_ID
	SubSkuId *string `json:"sub_sku_id,omitempty"`
	// 组套商品子商品发货数量
	ComboNum *int64 `json:"combo_num,omitempty"`
}
type ShippedOrderInfoItem struct {
	// 需要发货的子订单id
	ShippedOrderId *string `json:"shipped_order_id,omitempty"`
	// 上述子订单的待发货数
	ShippedNum int64 `json:"shipped_num,omitempty"`
	// 已废弃
	ShippedItemIds []string `json:"shipped_item_ids,omitempty"`
	// 组套商品参数列表
	BundleList []BundleListItem `json:"bundle_list,omitempty"`
}
type PackListItem struct {
	// 需要发货的子订单信息列表
	ShippedOrderInfo *[]ShippedOrderInfoItem `json:"shipped_order_info,omitempty"`
	// 运单号
	LogisticsCode *string `json:"logistics_code,omitempty"`
	// 物流公司名称
	Company string `json:"company,omitempty"`
	// 物流公司code，字段必传。可从/order/logisticsCompanyList接口获取。
	CompanyCode string `json:"company_code,omitempty"`
	// 已废弃。物流公司ID。请使用company_code字段。
	LogisticsId string `json:"logistics_id,omitempty"`
}
type OrderLogisticsAddMultiPackParam struct {
	// 父订单ID
	OrderId *string `json:"order_id,omitempty"`
	// 包裹list
	PackList *[]PackListItem `json:"pack_list,omitempty"`
	// 是否拒绝退款申请（true表示拒绝退款，并继续发货；不传或为false表示有退款需要处理，拒绝发货），is_refund_reject和is_reject_refund随机使用一个即可
	IsRejectRefund bool `json:"is_reject_refund,omitempty"`
	// 请求唯一标识，相同request_id多次请求，第一次请求成功后，后续的请求会触发幂等，会直接返回第一次请求成功的结果，不会实际触发发货。
	RequestId *string `json:"request_id,omitempty"`
	// 发货地址id，使用/address/list接口获取
	AddressId string `json:"address_id,omitempty"`
	// 商品序列号，单个序列号长度不能超过30位字符，其中手机序列号仅支持填写15～17位数字
	SerialNumberList []string `json:"serial_number_list,omitempty"`
	// 是否拒绝退款申请（true表示拒绝退款，并继续发货；不传或为false表示有退款需要处理，拒绝发货），is_refund_reject和is_reject_refund随机使用一个即可
	IsRefundReject bool `json:"is_refund_reject,omitempty"`
	// 门店id
	StoreId int64 `json:"store_id,omitempty"`
	// 退货地址ID,通过地址库列表【/address/list】接口查询
	AfterSaleAddressId int64 `json:"after_sale_address_id,omitempty"`
}
