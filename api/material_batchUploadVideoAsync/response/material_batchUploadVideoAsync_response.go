package material_batchUploadVideoAsync_response

import (
	"doudian.com/open/sdk_golang/core"
)

type MaterialBatchUploadVideoAsyncResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *MaterialBatchUploadVideoAsyncData `json:"data"`
}
type MaterialBatchUploadVideoAsyncData struct {
	// 成功上传的素材，key为参数中的request_id
	SuccessMap map[string]SuccessMapItem `json:"success_map"`
	// 失败的素材，key为参数中的request_id
	FailedMap map[string]FailedMapItem `json:"failed_map"`
}
type SuccessMapItem struct {
	// 素材id，素材中心对素材的唯一编号
	MaterialId string `json:"MaterialId"`
	// 素材名称，长度限制为50个字符
	Name string `json:"Name"`
	// 文件夹id，“0”表示将素材上传到了素材中心根目录
	FolderId string `json:"FolderId"`
	// 上传素材时，传入参数url的取值
	OriginUrl string `json:"OriginUrl"`
	// 审核状态 0-下载中 1-待审核 2-审核中 3-审核通过 4-审核拒绝
	AuditStatus int32 `json:"AuditStatus"`
	// 是否为新建的素材，true-表示新建
	IsNew bool `json:"IsNew"`
}
type FailedMapItem struct {
	// 错误码
	ErrCode int32 `json:"ErrCode"`
	// 错误描述
	ErrMsg string `json:"ErrMsg"`
}
