package order_logisticsAddSinglePack_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderLogisticsAddSinglePackResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderLogisticsAddSinglePackData `json:"data"`
}
type OrderLogisticsAddSinglePackData struct {
	// 包裹id
	PackId string `json:"pack_id"`
}
