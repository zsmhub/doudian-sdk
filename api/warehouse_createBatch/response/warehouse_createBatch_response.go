package warehouse_createBatch_response

import (
	"doudian.com/open/sdk_golang/core"
)

type WarehouseCreateBatchResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data map[string]bool `json:"data"`
}
type WarehouseCreateBatchData struct {
	// key是outWarehouseId,value代表成功/失败
	Data map[string]bool `json:"data"`
}
