package order_updateOrderAmount_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderUpdateOrderAmountResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderUpdateOrderAmountData `json:"data"`
}
type OrderUpdateOrderAmountData struct {
}
