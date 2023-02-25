package order_downloadToShop_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderDownloadToShopResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderDownloadToShopData `json:"data"`
}
type OrderDownloadToShopData struct {
	// 状态码    100000-成功   100002-下载记录不存在   100015-文件还未生成   100016-文件已经失效   100025-文件生成失败
	Code string `json:"code"`
	// 状态说明
	CodeMsg string `json:"code_msg"`
	// 生成的下载链接
	Url string `json:"url"`
}
