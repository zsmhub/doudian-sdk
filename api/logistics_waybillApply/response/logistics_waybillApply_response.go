package logistics_waybillApply_response

import (
	"doudian.com/open/sdk_golang/core"
)

type LogisticsWaybillApplyResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *LogisticsWaybillApplyData `json:"data"`
}
type WaybillInfosItem struct {
	// 订单号
	OrderId string `json:"order_id"`
	// 运单号
	TrackNo string `json:"track_no"`
	// 加密的面单数据
	PrintData string `json:"print_data"`
	// 签名信息
	Sign string `json:"sign"`
}
type ErrInfosItem struct {
	// 运单号
	TrackNo string `json:"track_no"`
	// 订单号
	OrderId string `json:"order_id"`
	// 错误码
	ErrCode int32 `json:"err_code"`
	// 错误信息
	ErrMsg string `json:"err_msg"`
}
type LogisticsWaybillApplyData struct {
	// 正常返回结构体
	WaybillInfos []WaybillInfosItem `json:"waybill_infos"`
	// 错误返回结构体
	ErrInfos []ErrInfosItem `json:"err_infos"`
}
