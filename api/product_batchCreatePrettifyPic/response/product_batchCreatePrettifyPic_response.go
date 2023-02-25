package product_batchCreatePrettifyPic_response

import (
	"doudian.com/open/sdk_golang/core"
)

type ProductBatchCreatePrettifyPicResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *ProductBatchCreatePrettifyPicData `json:"data"`
}
type PrettifyPicListItem struct {
	// 图片url
	ImgUrl string `json:"img_url"`
	// 图片宽度
	Width int64 `json:"width"`
	// 图片高度
	Height int64 `json:"height"`
	// 错误内容
	ErrorMsg string `json:"error_msg"`
	// 是否正确生成
	IsSuccess bool `json:"is_success"`
}
type ProductBatchCreatePrettifyPicData struct {
	// 返回内容，与入参顺序对应
	PrettifyPicList []PrettifyPicListItem `json:"prettify_pic_list"`
}
