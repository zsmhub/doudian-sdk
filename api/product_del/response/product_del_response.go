package product_del_response

import (
	"doudian.com/open/sdk_golang/core"
)

type ProductDelResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *ProductDelData `json:"data"`
}
type ProductDelData struct {
}
