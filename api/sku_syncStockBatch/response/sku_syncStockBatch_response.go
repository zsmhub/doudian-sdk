package sku_syncStockBatch_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SkuSyncStockBatchResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SkuSyncStockBatchData `json:"data"`
}
type SkuSyncStockBatchData struct {
}
