package superm_applyPlatformPickUp_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SupermApplyPlatformPickUpResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SupermApplyPlatformPickUpData `json:"data"`
}
type SupermApplyPlatformPickUpData struct {
	// 物流单ID
	LogisticsId string `json:"logistics_id"`
}
