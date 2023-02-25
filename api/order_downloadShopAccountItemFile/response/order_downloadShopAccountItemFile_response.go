package order_downloadShopAccountItemFile_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderDownloadShopAccountItemFileResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderDownloadShopAccountItemFileData `json:"data"`
}
type OrderDownloadShopAccountItemFileData struct {
	// 文件url
	Url string `json:"url"`
}
