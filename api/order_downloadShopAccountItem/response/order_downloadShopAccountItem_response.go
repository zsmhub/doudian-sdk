package order_downloadShopAccountItem_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderDownloadShopAccountItemResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderDownloadShopAccountItemData `json:"data"`
}
type OrderDownloadShopAccountItemData struct {
	// 返回code 100000为成功，其他为失败
	Code string `json:"code"`
	// 返回描述
	CodeMsg string `json:"code_msg"`
	// 下载id
	DownloadId string `json:"download_id"`
}
