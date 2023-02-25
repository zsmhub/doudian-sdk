package warehouse_removeAddr_response

import (
	"doudian.com/open/sdk_golang/core"
)

type WarehouseRemoveAddrResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *WarehouseRemoveAddrData `json:"data"`
}
type WarehouseRemoveAddrData struct {
}
