package superm_shopOrderDispatcher_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SupermShopOrderDispatcherResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SupermShopOrderDispatcherData `json:"data"`
}
type SupermShopOrderDispatcherData struct {
	// 取件码
	PickupCode string `json:"pickup_code"`
}
