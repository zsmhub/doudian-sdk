package afterSale_operate_response

import (
	"doudian.com/open/sdk_golang/core"
)

type AfterSaleOperateResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *AfterSaleOperateData `json:"data"`
}
type ItemsItem struct {
	// 售后单号
	AftersaleId int64 `json:"aftersale_id"`
	// 操作结果码
	StatusCode int64 `json:"status_code"`
	// 操作结果描述
	StatusMsg string `json:"status_msg"`
}
type AfterSaleOperateData struct {
	// 审核结果
	Items []ItemsItem `json:"items"`
}
