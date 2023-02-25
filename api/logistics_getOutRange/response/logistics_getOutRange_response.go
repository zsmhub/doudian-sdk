package logistics_getOutRange_response

import (
	"doudian.com/open/sdk_golang/core"
)

type LogisticsGetOutRangeResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *LogisticsGetOutRangeData `json:"data"`
}
type LogisticsGetOutRangeData struct {
	// 是否超区响应结果（超区-true；未超区-fasle）
	IsOutRange bool `json:"is_out_range"`
	// 超区原因，1、疫情管控 2、洪涝台风等自然灾害 3、特殊会议管控 4、网点经营问题 5、其他原因
	OutRangeReason string `json:"out_range_reason"`
}
