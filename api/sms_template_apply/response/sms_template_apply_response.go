package sms_template_apply_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SmsTemplateApplyResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SmsTemplateApplyData `json:"data"`
}
type SmsTemplateApplyData struct {
	// 短信发送渠道，主要做资源隔离
	SmsAccount string `json:"sms_account"`
	// 短信模板内容： 英文短信：最多支持140个英文字符，超出将按140个字符截取为多条短信进行发送，费用按截取的条数收费； 非英文短信：最多支持140个英文字符，超出将按140个字符截取为多条短信进行发送，费用按截取的条数收费；
	TemplateContent string `json:"template_content"`
	// 短信模版名称
	TemplateName string `json:"template_name"`
	// 模版id
	SmsTemplateId string `json:"sms_template_id"`
	// 短信模板申请单id
	SmsTemplateApplyId string `json:"sms_template_apply_id"`
	// 是否成功 0表示成功
	Code int64 `json:"code"`
	// 说明
	Message string `json:"message"`
}
