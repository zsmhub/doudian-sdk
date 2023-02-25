package antispam_orderSend_response

import (
	"doudian.com/open/sdk_golang/core"
)

type AntispamOrderSendResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *AntispamOrderSendData `json:"data"`
}
type Decision struct {
	// 决策
	Decision string `json:"decision"`
	// 决策详情
	DecisionDetail string `json:"decision_detail"`
	// 提示信息
	HitStatus string `json:"hit_status"`
}
type AntispamOrderSendData struct {
	// 决策
	Decision *Decision `json:"decision"`
}
