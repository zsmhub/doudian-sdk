package open_binaryupload_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OpenBinaryuploadResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OpenBinaryuploadData `json:"data"`
}
type OpenBinaryuploadData struct {
	// 上传二进制文件的uri
	StoreUri string `json:"store_uri"`
}
