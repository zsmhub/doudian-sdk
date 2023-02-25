package warehouse_edit_response

import (
	"doudian.com/open/sdk_golang/core"
)

type WarehouseEditResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data bool `json:"data"`
}
type WarehouseEditData struct {
	// 修改结果
	Data bool `json:"data"`
}
