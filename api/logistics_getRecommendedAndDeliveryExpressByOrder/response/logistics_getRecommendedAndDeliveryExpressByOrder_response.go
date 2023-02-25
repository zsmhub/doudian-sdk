package logistics_getRecommendedAndDeliveryExpressByOrder_response

import (
	"doudian.com/open/sdk_golang/core"
)

type LogisticsGetRecommendedAndDeliveryExpressByOrderResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data []DataItem `json:"data"`
}
type CollectSignInfo struct {
	// 揽签时长,单位小时
	AvgCostHours string `json:"avg_cost_hours"`
	// 该线路商超过其他物流公司23.86%
	LevelPercent string `json:"level_percent"`
	// 优化百分比，时长下降6.1%
	OptimizedPercent string `json:"optimized_percent"`
}
type ExpressInfoListItem struct {
	// 物流商编码
	Express string `json:"express"`
	// 是否可达
	IsDeliverable bool `json:"is_deliverable"`
	// 是否订购电子面单
	IsShopEBill bool `json:"is_shop_eBill"`
	// 是否推荐
	IsRecommended bool `json:"is_recommended"`
	// 时长维度的推荐原因
	CollectSignInfo *CollectSignInfo `json:"collect_sign_info"`
}
type DataItem struct {
	// 订单号
	OrderId string `json:"order_id"`
	// 推荐及可达物流商信息集合
	ExpressInfoList []ExpressInfoListItem `json:"express_info_list"`
}
type LogisticsGetRecommendedAndDeliveryExpressByOrderData struct {
	// 返回值
	Data []DataItem `json:"data"`
}
