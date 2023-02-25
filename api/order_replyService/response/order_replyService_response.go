package order_replyService_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderReplyServiceResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderReplyServiceData `json:"data"`
}
type OrderReplyServiceData struct {
}
