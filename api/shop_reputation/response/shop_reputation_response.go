package shop_reputation_response

import (
	"doudian.com/open/sdk_golang/core"
)

type ShopReputationResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *ShopReputationData `json:"data"`
}
type ShopReputationData struct {
	// 店铺 id
	ShopId int64 `json:"shop_id"`
	// 店铺名称
	ShopName string `json:"shop_name"`
	// 商家体验分
	ShopScore string `json:"shop_score"`
	// 商品体验分
	ProductScore string `json:"product_score"`
	// 物流体验分
	LogisticsScore string `json:"logistics_score"`
	// 商家服务分
	ServiceScore string `json:"service_score"`
}
