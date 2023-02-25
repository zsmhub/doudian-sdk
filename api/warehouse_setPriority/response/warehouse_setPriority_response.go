package warehouse_setPriority_response

import (
	"doudian.com/open/sdk_golang/core"
)

type WarehouseSetPriorityResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *WarehouseSetPriorityData `json:"data"`
}
type WarehouseSetPriorityData struct {
}
