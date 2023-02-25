package product_getRecommendName_response

import (
	"doudian.com/open/sdk_golang/core"
)

type ProductGetRecommendNameResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *ProductGetRecommendNameData `json:"data"`
}
type RecommendNameItem struct {
	// 推荐场景
	RecommendScene string `json:"recommend_scene"`
	// 推荐结果
	RecommendValue string `json:"recommend_value"`
}
type ProductGetRecommendNameData struct {
	// 商品标题推荐结果，当返回值为空时，使用【/product/getProductUpdateRule】接口查询类目id查看recommend_name_rule.satisfy_prefix是否=true
	RecommendName []RecommendNameItem `json:"recommend_name"`
}
