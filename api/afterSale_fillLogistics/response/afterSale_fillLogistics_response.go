package afterSale_fillLogistics_response

import (
	"doudian.com/open/sdk_golang/core"
)

type AfterSaleFillLogisticsResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *AfterSaleFillLogisticsData `json:"data"`
}
type AfterSaleFillLogisticsData struct {
}
