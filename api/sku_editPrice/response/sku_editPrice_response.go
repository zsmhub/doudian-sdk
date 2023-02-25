package sku_editPrice_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SkuEditPriceResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SkuEditPriceData `json:"data"`
}
type SkuEditPriceData struct {
}
