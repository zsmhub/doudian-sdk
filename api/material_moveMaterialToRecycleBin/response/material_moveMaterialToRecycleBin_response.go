package material_moveMaterialToRecycleBin_response

import (
	"doudian.com/open/sdk_golang/core"
)

type MaterialMoveMaterialToRecycleBinResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *MaterialMoveMaterialToRecycleBinData `json:"data"`
}
type MaterialMoveMaterialToRecycleBinData struct {
	// 成功的素材id列表
	SuccessIds []string `json:"success_ids"`
	// 失败素材列表
	FailedMap map[string]FailedMapItem `json:"failed_map"`
}
type FailedMapItem struct {
	// 错误码
	Code int32 `json:"code"`
	// 错误原因
	Msg string `json:"msg"`
}
