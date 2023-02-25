package warehouse_create_response

import (
	"doudian.com/open/sdk_golang/core"
)

type WarehouseCreateResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data int64 `json:"data"`
}
type WarehouseCreateData struct {
	// 仓库id
	Data int64 `json:"data"`
}
