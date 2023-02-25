package sms_sign_apply_list_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SmsSignApplyListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SmsSignApplyListData `json:"data"`
}
type ApplyListItem struct {
	// 申请单id
	SmsSignApplyId string `json:"sms_sign_apply_id"`
	// 签名
	Sign string `json:"sign"`
	// 短信发送渠道，主要做资源隔离
	SmsAccount string `json:"sms_account"`
	// 审核状态： 1:审核中 2:未通过 3:已开通 4:已关闭 5:免审核（店铺名和签名完全一致时返回）
	Status int64 `json:"status"`
	// 说明
	StatusRemark string `json:"status_remark"`
}
type SmsSignApplyListData struct {
	// 总数
	Total int64 `json:"total"`
	// 短信签名申请单列表
	ApplyList []ApplyListItem `json:"apply_list"`
}
