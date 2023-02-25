package superm_getPlatformPickUpEstimatedCharge_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SupermGetPlatformPickUpEstimatedChargeResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SupermGetPlatformPickUpEstimatedChargeData `json:"data"`
}
type SupermGetPlatformPickUpEstimatedChargeData struct {
	// 运力费用预估，单位：分
	DispatcherFee int64 `json:"dispatcher_fee"`
}
