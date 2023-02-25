package order_review_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderReviewResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderReviewData `json:"data"`
}
type OrderReviewData struct {
}
