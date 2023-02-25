package order_addOrderRemark_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderAddOrderRemarkResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderAddOrderRemarkData `json:"data"`
}
type OrderAddOrderRemarkData struct {
}
