package order_ordeReportList_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderOrdeReportListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderOrdeReportListData `json:"data"`
}
type AddRealMobileWhitesItem struct {
	// 订单号
	OrderId string `json:"order_id"`
	// 售后单号
	AfterSaleId string `json:"after_sale_id"`
	// 报备失败原因
	Msg string `json:"msg"`
}
type OrderOrdeReportListData struct {
	// 报备响应
	AddRealMobileWhites []AddRealMobileWhitesItem `json:"add_real_mobile_whites"`
}
