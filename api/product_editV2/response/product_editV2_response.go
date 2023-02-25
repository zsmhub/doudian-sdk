package product_editV2_response

import (
	"doudian.com/open/sdk_golang/core"
)

type ProductEditV2Response struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *ProductEditV2Data `json:"data"`
}
type ProductEditV2Data struct {
}
