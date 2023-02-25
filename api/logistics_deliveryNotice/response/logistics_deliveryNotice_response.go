package logistics_deliveryNotice_response

import (
	"doudian.com/open/sdk_golang/core"
)

type LogisticsDeliveryNoticeResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *LogisticsDeliveryNoticeData `json:"data"`
}
type LogisticsDeliveryNoticeData struct {
	// 是否成功
	Result bool `json:"result"`
}
