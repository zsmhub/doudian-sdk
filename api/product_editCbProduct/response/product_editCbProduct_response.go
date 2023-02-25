package product_editCbProduct_response

import (
	"doudian.com/open/sdk_golang/core"
)

type ProductEditCbProductResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *ProductEditCbProductData `json:"data"`
}
type ProductEditCbProductData struct {
}
