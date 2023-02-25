package sms_template_apply_request

import (
	"doudian.com/open/sdk_golang/api/sms_template_apply/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SmsTemplateApplyRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SmsTemplateApplyParam 
}
func (c *SmsTemplateApplyRequest) GetUrlPath() string{
	return "/sms/template/apply"
}


func New() *SmsTemplateApplyRequest{
	request := &SmsTemplateApplyRequest{
		Param: &SmsTemplateApplyParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SmsTemplateApplyRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sms_template_apply_response.SmsTemplateApplyResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sms_template_apply_response.SmsTemplateApplyResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SmsTemplateApplyRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SmsTemplateApplyRequest) GetParams() *SmsTemplateApplyParam{
	return c.Param
}


type SmsTemplateApplyParam struct {
	// 短信发送渠道，主要做资源隔离
	SmsAccount *string `json:"sms_account,omitempty"`
	// CN_NTC 国内通知短信 CN_MKT 国内营销短信（营销短信将自带退订功能） CN_OTP 国内验证码
	TemplateType *string `json:"template_type,omitempty"`
	// 短信模板名称
	TemplateName *string `json:"template_name,omitempty"`
	// 短信模板内容： 英文短信：整条短信（包括签名+模板+变量中的内容）最多支持140个英文字符，超出将按140个字符截取为多条短信进行发送，费用按截取的条数收费； 非英文短信：整条短信（包括签名+模板+变量中的内容）最多支持70字符，超出将按70个字符截取为多条短信进行发送，费用按截取的条数收费；
	TemplateContent *string `json:"template_content,omitempty"`
}
