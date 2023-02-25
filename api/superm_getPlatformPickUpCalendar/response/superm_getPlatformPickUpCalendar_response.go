package superm_getPlatformPickUpCalendar_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SupermGetPlatformPickUpCalendarResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SupermGetPlatformPickUpCalendarData `json:"data"`
}
type PeriodListItem struct {
	// 时间段开始时间；Unix时间戳，单位：秒
	TimeBeginTs int64 `json:"time_begin_ts"`
	// 时间段结束时间；Unix时间戳，单位：秒
	TimeEndTs int64 `json:"time_end_ts"`
}
type SupermGetPlatformPickUpCalendarData struct {
	// 是否支持查询
	IsSupport bool `json:"is_support"`
	// 不支持的原因
	UnSupportReason string `json:"un_support_reason"`
	// 时间段
	PeriodList []PeriodListItem `json:"period_list"`
}
