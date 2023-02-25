package brand_convert_response

import (
	"doudian.com/open/sdk_golang/core"
)

type BrandConvertResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *BrandConvertData `json:"data"`
}
type BrandConvertData struct {
	// 品牌id，对应商品发布接口standard_brand_id字段
	BrandId int64 `json:"brand_id"`
}
