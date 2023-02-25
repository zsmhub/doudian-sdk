package material_recoverFolder_response

import (
	"doudian.com/open/sdk_golang/core"
)

type MaterialRecoverFolderResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *MaterialRecoverFolderData `json:"data"`
}
type MaterialRecoverFolderData struct {
	// 操作成功的文件夹列表
	SuccessIds []string `json:"success_ids"`
	// 操作失败的文件夹及失败的详情
	FailedMap map[int64]FailedMapItem `json:"failed_map"`
}
type FailedMapItem struct {
	// 操作失败的错误码
	Code int32 `json:"code"`
	// 操作失败的原因
	Msg string `json:"msg"`
}
