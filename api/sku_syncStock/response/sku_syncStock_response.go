package sku_syncStock_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SkuSyncStockResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SkuSyncStockData `json:"data"`
}
type SkuSyncStockData struct {
}
