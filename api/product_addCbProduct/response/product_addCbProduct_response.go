package product_addCbProduct_response

import (
	"doudian.com/open/sdk_golang/core"
)

type ProductAddCbProductResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *ProductAddCbProductData `json:"data"`
}
type SkuItem struct {
	// 1
	SkuId int64 `json:"sku_id"`
	// 1
	OutSkuId int64 `json:"out_sku_id"`
	// 1
	OuterSkuId string `json:"outer_sku_id"`
	// 1
	Code string `json:"code"`
	// 1
	SpecDetailId1 int64 `json:"spec_detail_id1"`
	// 1
	SpecDetailId2 int64 `json:"spec_detail_id2"`
	// 1
	SpecDetailId3 int64 `json:"spec_detail_id3"`
}
type ProductAddCbProductData struct {
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
