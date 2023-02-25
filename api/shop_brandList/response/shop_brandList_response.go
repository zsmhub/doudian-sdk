package shop_brandList_response

import (
	"doudian.com/open/sdk_golang/core"
)

type ShopBrandListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data []DataItem `json:"data"`
}
type DataItem struct {
	// 品牌ID
	Id int64 `json:"id"`
	// 品牌中文名
	BrandChineseName string `json:"brand_chinese_name"`
	// 品牌英文名
	BrandEnglishName string `json:"brand_english_name"`
	// 商标注册号
	BrandRegNum string `json:"brand_reg_num"`
}
type ShopBrandListData struct {
	// 返回值
	Data []DataItem `json:"data"`
}
