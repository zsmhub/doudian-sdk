package product_setOnline_response

import (
	"doudian.com/open/sdk_golang/core"
)

type ProductSetOnlineResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *ProductSetOnlineData `json:"data"`
}
type ProductSetOnlineData struct {
}
