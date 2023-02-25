package logistics_trackNoRouteDetail_response

import (
	"doudian.com/open/sdk_golang/core"
)

type LogisticsTrackNoRouteDetailResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *LogisticsTrackNoRouteDetailData `json:"data"`
}
type RouteNodeListItem struct {
	// 轨迹内容
	Content string `json:"content"`
	// 时间戳；单位：秒
	Timestamp string `json:"timestamp"`
	// 轨迹状态code；枚举值详见：https://op.jinritemai.com/docs/question-docs/94/1642
	State string `json:"state"`
	// 轨迹状态描述
	StateDescription string `json:"state_description"`
	// 站点名称
	SiteName string `json:"site_name"`
}
type TrackInfo struct {
	// 运单号
	TrackNo string `json:"track_no"`
	// 物流商编码
	Express string `json:"express"`
}
type LogisticsTrackNoRouteDetailData struct {
	// 轨迹信息
	RouteNodeList []RouteNodeListItem `json:"route__node_list"`
	// 运单信息
	TrackInfo *TrackInfo `json:"track_info"`
}
