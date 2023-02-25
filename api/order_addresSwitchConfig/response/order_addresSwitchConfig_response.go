package order_addresSwitchConfig_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderAddresSwitchConfigResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderAddresSwitchConfigData `json:"data"`
}
type OrderAddresSwitchConfigData struct {
	// 0: 商家当前未开启审核 1:商家当前已开启审核，但本应用不可进行审核 2:商家当前已开启审核，且本应用可审核
	AuthorizedStatus int64 `json:"authorized_status"`
}
