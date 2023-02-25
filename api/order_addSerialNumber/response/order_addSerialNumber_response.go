package order_addSerialNumber_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderAddSerialNumberResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderAddSerialNumberData `json:"data"`
}
type OrderAddSerialNumberData struct {
}
