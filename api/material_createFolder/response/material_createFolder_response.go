package material_createFolder_response

import (
	"doudian.com/open/sdk_golang/core"
)

type MaterialCreateFolderResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *MaterialCreateFolderData `json:"data"`
}
type MaterialCreateFolderData struct {
	// 文件夹id，全局唯一
	FolderId string `json:"folder_id"`
	// 父文件夹id，全局唯一
	ParentFolderId string `json:"parent_folder_id"`
	// 文件夹名称
	Name string `json:"name"`
	// 文件夹类型。0-用户自建 1-默认 2-系统文件夹
	Type int32 `json:"type"`
}
