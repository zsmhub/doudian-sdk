package sms_template_apply_list_request

import (
	"doudian.com/open/sdk_golang/api/sms_template_apply_list/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SmsTemplateApplyListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SmsTemplateApplyListParam 
}
func (c *SmsTemplateApplyListRequest) GetUrlPath() string{
	return "/sms/template/apply/list"
}


func New() *SmsTemplateApplyListRequest{
	request := &SmsTemplateApplyListRequest{
		Param: &SmsTemplateApplyListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SmsTemplateApplyListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sms_template_apply_list_response.SmsTemplateApplyListResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sms_template_apply_list_response.SmsTemplateApplyListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SmsTemplateApplyListRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SmsTemplateApplyListRequest) GetParams() *SmsTemplateApplyListParam{
	return c.Param
}


type SmsTemplateApplyListParam struct {
	// 短信发送渠道，主要做资源隔离
	SmsAccount *string `json:"sms_account,omitempty"`
	// 短信模板内容： 英文短信：整条短信（包括签名+模板+变量中的内容）最多支持140个英文字符，超出将按140个字符截取为多条短信进行发送，费用按截取的条数收费； 非英文短信：整条短信（包括签名+模板+变量中的内容）最多支持70字符，超出将按70个字符截取为多条短信进行发送，费用按截取的条数收费；
	TemplateContent string `json:"template_content,omitempty"`
	// 短信模板申请单id
	SmsTemplateApplyId string `json:"sms_template_apply_id,omitempty"`
	// 每页数量，默认10
	Size int64 `json:"size,omitempty"`
	// 页码，默认0
	Page int64 `json:"page,omitempty"`
	// 审核状态： 1:审核中 2:未通过 3:已开通 4:已关闭
	Status int64 `json:"status,omitempty"`
}
