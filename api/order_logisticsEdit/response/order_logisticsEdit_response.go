package order_logisticsEdit_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderLogisticsEditResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderLogisticsEditData `json:"data"`
}
type OrderLogisticsEditData struct {
}
