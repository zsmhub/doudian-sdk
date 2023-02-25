package order_logisticsAddMultiPack_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderLogisticsAddMultiPackResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderLogisticsAddMultiPackData `json:"data"`
}
type PackListItem struct {
	// 发货的订单信息
	ShippedOrderInfo []ShippedOrderInfoItem `json:"shipped_order_info"`
	// 物流单号
	LogisticsCode string `json:"logistics_code"`
	// 包裹id
	PackId string `json:"pack_id"`
}
type OrderLogisticsAddMultiPackData struct {
	// 包裹信息
	PackList []PackListItem `json:"pack_list"`
}
type ShippedOrderInfoItem struct {
	// 发货的子单id
	ShippedOrderId string `json:"shipped_order_id"`
	// 发货子单数量
	ShippedNum int64 `json:"shipped_num"`
	// 发货的四层单id
	ShippedItemIds []string `json:"shipped_item_ids"`
}
