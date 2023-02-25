package afterSale_timeExtend_response

import (
	"doudian.com/open/sdk_golang/core"
)

type AfterSaleTimeExtendResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *AfterSaleTimeExtendData `json:"data"`
}
type AfterSaleTimeExtendData struct {
}
