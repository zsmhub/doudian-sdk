package logistics_cancelOrder_response

import (
	"doudian.com/open/sdk_golang/core"
)

type LogisticsCancelOrderResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *LogisticsCancelOrderData `json:"data"`
}
type LogisticsCancelOrderData struct {
	// 取消状态
	CancelResult *CancelResult `json:"cancel_result"`
}
type CancelResult struct {
	// true：取消成功 false 取消失败
	Success bool `json:"success"`
}
