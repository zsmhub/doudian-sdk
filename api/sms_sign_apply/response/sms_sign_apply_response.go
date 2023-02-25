package sms_sign_apply_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SmsSignApplyResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SmsSignApplyData `json:"data"`
}
type SmsSignApplyData struct {
	// 短信发送渠道，主要做资源隔离
	SmsAccount string `json:"sms_account"`
	// 短信签名申请单id
	SmsSignApplyId string `json:"sms_sign_apply_id"`
	// 是否成功 0表示成功
	Code int64 `json:"code"`
	// 说明
	Message string `json:"message"`
	// 签名
	Sign string `json:"sign"`
	// 审核状态： 1:审核中 2:未通过 3:已开通 4:已关闭 5:免审核（店铺名和签名完全一致时返回）
	Status int64 `json:"status"`
}
