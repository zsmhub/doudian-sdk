package material_deleteMaterial_response

import (
	"doudian.com/open/sdk_golang/core"
)

type MaterialDeleteMaterialResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *MaterialDeleteMaterialData `json:"data"`
}
type FailedMapItem struct {
	// 失败code
	Code int32 `json:"code"`
	// 失败原因
	Msg string `json:"msg"`
}
type MaterialDeleteMaterialData struct {
	// 操作成功的素材id
	SuccessIds []string `json:"success_ids"`
	// 操作失败的素材id和原因
	FailedMap map[string]FailedMapItem `json:"failed_map"`
}
