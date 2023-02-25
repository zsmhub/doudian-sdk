package logistics_appendSubOrder_response

import (
	"doudian.com/open/sdk_golang/core"
)

type LogisticsAppendSubOrderResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *LogisticsAppendSubOrderData `json:"data"`
}
type LogisticsAppendSubOrderData struct {
	// 运单号
	TrackNo string `json:"track_no"`
	// 子母单数量
	PackQuantity int64 `json:"pack_quantity"`
	// 新追加的子单号
	SubWaybillCodes string `json:"sub_waybill_codes"`
}
