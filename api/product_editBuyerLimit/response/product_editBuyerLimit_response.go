package product_editBuyerLimit_response

import (
	"doudian.com/open/sdk_golang/core"
)

type ProductEditBuyerLimitResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *ProductEditBuyerLimitData `json:"data"`
}
type ProductEditBuyerLimitData struct {
}
