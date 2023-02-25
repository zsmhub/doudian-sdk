package order_updatePostAmount_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderUpdatePostAmountResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderUpdatePostAmountData `json:"data"`
}
type OrderUpdatePostAmountData struct {
}
