package order_logisticsAdd_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderLogisticsAddResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderLogisticsAddData `json:"data"`
}
type OrderLogisticsAddData struct {
}
