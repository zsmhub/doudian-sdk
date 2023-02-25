package material_uploadVideoAsync_response

import (
	"doudian.com/open/sdk_golang/core"
)

type MaterialUploadVideoAsyncResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *MaterialUploadVideoAsyncData `json:"data"`
}
type MaterialUploadVideoAsyncData struct {
	// 素材id
	MaterialId string `json:"material_id"`
	// 素材所在文件夹id，0-素材中心的根目录
	FolderId string `json:"folder_id"`
	// 是否是新建，false-重复请求 true-新建
	IsNew bool `json:"is_new"`
	// 素材审核状态 0-下载中 1-等待审核 2-审核中 3-通过 4-拒绝
	AuditStatus int32 `json:"audit_status"`
}
