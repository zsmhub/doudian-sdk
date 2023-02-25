package superm_getDispatcherInfo_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SupermGetDispatcherInfoResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SupermGetDispatcherInfoData `json:"data"`
}
type SupermGetDispatcherInfoData struct {
	// 运力费用，单位：分
	DispatcherFee int64 `json:"dispatcher_fee"`
}
