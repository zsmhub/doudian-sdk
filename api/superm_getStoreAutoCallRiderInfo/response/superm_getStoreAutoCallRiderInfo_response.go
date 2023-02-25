package superm_getStoreAutoCallRiderInfo_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SupermGetStoreAutoCallRiderInfoResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SupermGetStoreAutoCallRiderInfoData `json:"data"`
}
type AutocallInfo struct {
	// 自动呼叫运力设置状态；1：关闭 2：开启
	ServiceStatus int64 `json:"service_status"`
	// 自动呼叫运力策略； 0：接单后立即呼叫  1：接单后延迟呼叫
	ServiceType int64 `json:"service_type"`
	// 接单后延时呼叫时间，service_type为1时生效，单位min，范围[1,15]
	DelayTime int64 `json:"delay_time"`
}
type SupermGetStoreAutoCallRiderInfoData struct {
	// 自动呼叫设置信息
	AutocallInfo *AutocallInfo `json:"autocall_info"`
}
