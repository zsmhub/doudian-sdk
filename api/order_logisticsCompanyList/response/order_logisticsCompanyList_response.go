package order_logisticsCompanyList_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderLogisticsCompanyListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data []DataItem `json:"data"`
}
type OrderLogisticsCompanyListData struct {
	// 物流公司信息
	Data []DataItem `json:"data"`
}
type DataItem struct {
	// 物流公司id
	Id int64 `json:"id"`
	// 物流公司名字
	Name string `json:"name"`
	// 物流公司code
	Code string `json:"code"`
}
