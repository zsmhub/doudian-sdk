package product_addV2_response

import (
	"doudian.com/open/sdk_golang/core"
)

type ProductAddV2Response struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *ProductAddV2Data `json:"data"`
}
type SkuItem struct {
	// sku_id
	SkuId int64 `json:"sku_id"`
	// 外部sku id
	OutSkuId int64 `json:"out_sku_id"`
	// 字符串形式外部sku id
	OuterSkuId string `json:"outer_sku_id"`
	// sku code
	Code string `json:"code"`
	// 子规格id
	SpecDetailId1 int64 `json:"spec_detail_id1"`
	// 子规格id
	SpecDetailId2 int64 `json:"spec_detail_id2"`
	// 子规格id
	SpecDetailId3 int64 `json:"spec_detail_id3"`
	// 所有的规格值id，代替spec_detail_id123
	SpecDetailIds []int64 `json:"spec_detail_ids"`
}
type ProductAddV2Data struct {
	// product_id
	ProductId int64 `json:"product_id"`
	// out_product_id
	OutProductId int64 `json:"out_product_id"`
	// outer_product_id
	OuterProductId string `json:"outer_product_id"`
	// create_time
	CreateTime string `json:"create_time"`
	// 1
	Sku []SkuItem `json:"sku"`
}
