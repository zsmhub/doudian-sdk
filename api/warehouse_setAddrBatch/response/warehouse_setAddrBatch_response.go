package warehouse_setAddrBatch_response

import (
	"doudian.com/open/sdk_golang/core"
)

type WarehouseSetAddrBatchResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *WarehouseSetAddrBatchData `json:"data"`
}
type WarehouseSetAddrBatchData struct {
}
