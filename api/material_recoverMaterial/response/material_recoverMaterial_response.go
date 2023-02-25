package material_recoverMaterial_response

import (
	"doudian.com/open/sdk_golang/core"
)

type MaterialRecoverMaterialResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *MaterialRecoverMaterialData `json:"data"`
}
type FailedMapItem struct {
	// 错误码
	Code int32 `json:"code"`
	// 错误码描述
	Msg string `json:"msg"`
}
type MaterialRecoverMaterialData struct {
	// 成功操作的素材id列表
	SuccessIds []string `json:"success_ids"`
	// 失败素材列表及失败原因
	FailedMap map[string]FailedMapItem `json:"failed_map"`
}
