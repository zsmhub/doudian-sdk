package logistics_updateOrder_response

import (
	"doudian.com/open/sdk_golang/core"
)

type LogisticsUpdateOrderResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *LogisticsUpdateOrderData `json:"data"`
}
type LogisticsUpdateOrderData struct {
	// 运单号
	TrackNo string `json:"track_no"`
	// 分拣码（三段码）
	SortCode string `json:"sort_code"`
	// 集包地代码
	PackageCenterCode string `json:"package_center_code"`
	// 集包名称
	PackageCenterName string `json:"package_center_name"`
	// 大头笔编码
	ShortAddressCode string `json:"short_address_code"`
	// 大头笔名称
	ShortAddressName string `json:"short_address_name"`
	// 扩展字段
	ExtraResp string `json:"extra_resp"`
}
