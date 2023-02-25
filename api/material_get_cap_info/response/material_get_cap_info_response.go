package material_get_cap_info_response

import (
	"doudian.com/open/sdk_golang/core"
)

type MaterialGetCapInfoResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *MaterialGetCapInfoData `json:"data"`
}
type EquityInfoItem struct {
	// 权益类型
	EquityType string `json:"equity_type"`
	// 权益描述
	EquityTypeDesc string `json:"equity_type_desc"`
	// 权益开始时间
	BeginTime int64 `json:"begin_time"`
	// 权益结束时间
	EndTime int64 `json:"end_time"`
	// 该权益对应的总容量，单位KB
	TotalCapacity int64 `json:"total_capacity"`
	// 该权益对应的图片总容量，单位KB
	PhotoCapacity int64 `json:"photo_capacity"`
	// 该权益对应的视频总容量，单位KB
	VideoCapacity int64 `json:"video_capacity"`
}
type MaterialGetCapInfoData struct {
	// 总容量，单位KB
	TotalCapacity int64 `json:"total_capacity"`
	// 已使用容量，单位KB
	TotalCapacityUsed int64 `json:"total_capacity_used"`
	// 图片总容量，单位KB
	PhotoCapacity int64 `json:"photo_capacity"`
	// 图片已使用容量，单位KB
	PhotoCapacityUsed int64 `json:"photo_capacity_used"`
	// 视频总容量，单位KB
	VideoCapacity int64 `json:"video_capacity"`
	// 视频已使用容量，单位KB
	VideoCapacityUsed int64 `json:"video_capacity_used"`
	// 当前生效的权益列表
	EquityInfo []EquityInfoItem `json:"equity_info"`
}
