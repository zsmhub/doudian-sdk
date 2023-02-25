package antispam_userLogin_response

import (
	"doudian.com/open/sdk_golang/core"
)

type AntispamUserLoginResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *AntispamUserLoginData `json:"data"`
}
type Decision struct {
	// 动作
	Decision string `json:"decision"`
	// 详情
	DecisionDetail string `json:"decision_detail"`
	// 状态
	HitStatus string `json:"hit_status"`
}
type AntispamUserLoginData struct {
	// 决议
	Decision *Decision `json:"decision"`
}
