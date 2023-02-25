package warehouse_setAddr_response

import (
	"doudian.com/open/sdk_golang/core"
)

type WarehouseSetAddrResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *WarehouseSetAddrData `json:"data"`
}
type WarehouseSetAddrData struct {
}
