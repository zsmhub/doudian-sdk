package sms_template_revoke_request

import (
	"doudian.com/open/sdk_golang/api/sms_template_revoke/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SmsTemplateRevokeRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SmsTemplateRevokeParam 
}
func (c *SmsTemplateRevokeRequest) GetUrlPath() string{
	return "/sms/template/revoke"
}


func New() *SmsTemplateRevokeRequest{
	request := &SmsTemplateRevokeRequest{
		Param: &SmsTemplateRevokeParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SmsTemplateRevokeRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sms_template_revoke_response.SmsTemplateRevokeResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sms_template_revoke_response.SmsTemplateRevokeResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SmsTemplateRevokeRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SmsTemplateRevokeRequest) GetParams() *SmsTemplateRevokeParam{
	return c.Param
}


type SmsTemplateRevokeParam struct {
	// 短信发送渠道，主要做资源隔离
	SmsAccount *string `json:"sms_account,omitempty"`
	// 短信模板申请单id
	SmsTemplateApplyId *string `json:"sms_template_apply_id,omitempty"`
}
