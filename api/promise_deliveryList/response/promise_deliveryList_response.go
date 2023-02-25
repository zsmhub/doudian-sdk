package promise_deliveryList_response

import (
	"doudian.com/open/sdk_golang/core"
)

type PromiseDeliveryListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *PromiseDeliveryListData `json:"data"`
}
type ProductsItem struct {
	// 商品ID
	ProductId int64 `json:"product_id"`
	// 1：当日发；2：次日发
	ShipMode int32 `json:"ship_mode"`
}
type PromiseDeliveryListData struct {
	// 商品返回结构
	Products []ProductsItem `json:"products"`
	// 总数
	Total int64 `json:"total"`
}
