package logistics_queryPackageRoute_response

import (
	"doudian.com/open/sdk_golang/core"
)

type LogisticsQueryPackageRouteResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *Data `json:"data"`
}
type TrackRoutesItem struct {
	// 内容
	Content string `json:"content"`
	// 状态
	State string `json:"state"`
	// 时间
	StateTime int64 `json:"state_time"`
	// code
	Opcode string `json:"opcode"`
}
type Data struct {
	// 运单号
	TrackNo string `json:"track_no"`
	// 快递商
	Express string `json:"express"`
	// 轨迹
	TrackRoutes []TrackRoutesItem `json:"track_routes"`
}
type LogisticsQueryPackageRouteData struct {
	// 数据
	Data *Data `json:"data"`
}
