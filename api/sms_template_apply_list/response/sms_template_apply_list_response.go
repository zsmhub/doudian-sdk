package sms_template_apply_list_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SmsTemplateApplyListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SmsTemplateApplyListData `json:"data"`
}
type TemplateApplyListItem struct {
	// 短信发送渠道，主要做资源隔离
	SmsAccount string `json:"sms_account"`
	// 短信模板申请单id
	SmsTemplateApplyId string `json:"sms_template_apply_id"`
	// 短信模版名称
	TemplateName string `json:"template_name"`
	// 短信模版内容
	TemplateContent string `json:"template_content"`
	// CN_NTC 国内通知短信 CN_MKT 国内营销短信（营销短信将自带退订功能） CN_OTP 国内验证码
	ChannelType string `json:"channel_type"`
	// 审核状态： 1:审核中 2:未通过 3:已开通 4:已关闭
	Status int64 `json:"status"`
	// 状态说明
	StatusRemark string `json:"status_remark"`
}
type SmsTemplateApplyListData struct {
	// 短信模板申请单列表
	TemplateApplyList []TemplateApplyListItem `json:"template_apply_list"`
	// 总数
	Total int64 `json:"total"`
}
