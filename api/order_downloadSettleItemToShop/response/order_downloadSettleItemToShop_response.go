package order_downloadSettleItemToShop_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderDownloadSettleItemToShopResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderDownloadSettleItemToShopData `json:"data"`
}
type OrderDownloadSettleItemToShopData struct {
	// 返回的download_id
	DownloadId string `json:"download_id"`
	// 状态码
	Code string `json:"code"`
	// 状态信息
	CodeMsg string `json:"code_msg"`
}
