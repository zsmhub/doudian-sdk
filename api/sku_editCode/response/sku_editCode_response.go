package sku_editCode_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SkuEditCodeResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SkuEditCodeData `json:"data"`
}
type SkuEditCodeData struct {
}
