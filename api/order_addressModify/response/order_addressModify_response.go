package order_addressModify_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderAddressModifyResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderAddressModifyData `json:"data"`
}
type OrderAddressModifyData struct {
}
