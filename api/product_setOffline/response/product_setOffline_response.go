package product_setOffline_response

import (
	"doudian.com/open/sdk_golang/core"
)

type ProductSetOfflineResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *ProductSetOfflineData `json:"data"`
}
type ProductSetOfflineData struct {
}
