package order_addressConfirm_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderAddressConfirmResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderAddressConfirmData `json:"data"`
}
type OrderAddressConfirmData struct {
}
