package material_moveFolderToRecycleBin_response

import (
	"doudian.com/open/sdk_golang/core"
)

type MaterialMoveFolderToRecycleBinResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *MaterialMoveFolderToRecycleBinData `json:"data"`
}
type FailedMapItem struct {
	// 失败code
	Code int32 `json:"code"`
	// 失败原因
	Msg string `json:"msg"`
}
type MaterialMoveFolderToRecycleBinData struct {
	// 操作成功的id list
	SuccessIds []string `json:"success_ids"`
	// 操作失败的id map
	FailedMap map[int64]FailedMapItem `json:"failed_map"`
}
