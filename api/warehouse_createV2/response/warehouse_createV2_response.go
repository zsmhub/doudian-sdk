package warehouse_createV2_response

import (
	"doudian.com/open/sdk_golang/core"
)

type WarehouseCreateV2Response struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *WarehouseCreateV2Data `json:"data"`
}
type WarehouseCreateV2Data struct {
	// 内部仓id
	WarehouseId string `json:"warehouse_id"`
}
