package spu_createSpu_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SpuCreateSpuResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SpuCreateSpuData `json:"data"`
}
type SpuCreateSpuData struct {
	// SPU的ID
	SpuId string `json:"spu_id"`
}
