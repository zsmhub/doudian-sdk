package sms_template_search_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SmsTemplateSearchResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SmsTemplateSearchData `json:"data"`
}
type TemplateSearchListItem struct {
	// 短信发送渠道，主要做资源隔离
	SmsAccount string `json:"sms_account"`
	// 短信模版内容
	TemplateContent string `json:"template_content"`
	// 短信模板id
	TemplateId string `json:"template_id"`
	// CN_NTC 国内通知短信 CN_MKT 国内营销短信（免审模版营销短信将自带退订功能，非免审模版需自行添加退订描述） CN_OTP 国内验证码
	ChannelType string `json:"channel_type"`
	// 短信模版名称
	TemplateName string `json:"template_name"`
}
type SmsTemplateSearchData struct {
	// 模板列表
	TemplateSearchList []TemplateSearchListItem `json:"template_search_list"`
	// 总条数
	Total int64 `json:"total"`
}
