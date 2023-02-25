package antispam_orderQuery_response

import (
	"doudian.com/open/sdk_golang/core"
)

type AntispamOrderQueryResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *AntispamOrderQueryData `json:"data"`
}
type AntispamOrderQueryData struct {
	// 决策
	Decision *Decision `json:"decision"`
}
type Decision struct {
	// 决策
	Decision string `json:"decision"`
	// 决策详情
	DecisionDetail string `json:"decision_detail"`
	// 提示信息
	HitStatus string `json:"hit_status"`
}
