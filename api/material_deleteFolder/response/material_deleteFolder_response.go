package material_deleteFolder_response

import (
	"doudian.com/open/sdk_golang/core"
)

type MaterialDeleteFolderResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *MaterialDeleteFolderData `json:"data"`
}
type MaterialDeleteFolderData struct {
	// 操作成功的文件夹id列表
	SuccessIds []string `json:"success_ids"`
	// 操作失败的的文件夹及错误详情
	FailedMap map[string]FailedMapItem `json:"failed_map"`
}
type FailedMapItem struct {
	// 错误码
	Code int32 `json:"code"`
	// 错误码的描述字段
	Msg string `json:"msg"`
}
