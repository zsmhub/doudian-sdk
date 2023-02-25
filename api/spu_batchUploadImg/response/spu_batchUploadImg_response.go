package spu_batchUploadImg_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SpuBatchUploadImgResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SpuBatchUploadImgData `json:"data"`
}
type SpuBatchUploadImgData struct {
	// 成功列表
	SuccessList []SuccessListItem `json:"success_list"`
	// 失败列表
	FailedList []FailedListItem `json:"failed_list"`
}
type SuccessListItem struct {
	// 图片名称
	Name string `json:"name"`
	// 入参URL
	OriginUrl string `json:"origin_url"`
	// 抖店SPU专用URL
	SpuUrl string `json:"spu_url"`
}
type FailedListItem struct {
	// 图片名
	Name string `json:"name"`
	// 入参URL
	OriginUrl string `json:"origin_url"`
	// 失败错误码
	ErrCode int32 `json:"err_code"`
	// 失败原因
	ErrMsg string `json:"err_msg"`
}
