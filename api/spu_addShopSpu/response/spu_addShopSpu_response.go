package spu_addShopSpu_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SpuAddShopSpuResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SpuAddShopSpuData `json:"data"`
}
type SpuAddShopSpuData struct {
	// spuId
	SpuId string `json:"spu_id"`
}
