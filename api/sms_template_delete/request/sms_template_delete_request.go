package sms_template_delete_request

import (
	"doudian.com/open/sdk_golang/api/sms_template_delete/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SmsTemplateDeleteRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SmsTemplateDeleteParam 
}
func (c *SmsTemplateDeleteRequest) GetUrlPath() string{
	return "/sms/template/delete"
}


func New() *SmsTemplateDeleteRequest{
	request := &SmsTemplateDeleteRequest{
		Param: &SmsTemplateDeleteParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SmsTemplateDeleteRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sms_template_delete_response.SmsTemplateDeleteResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sms_template_delete_response.SmsTemplateDeleteResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SmsTemplateDeleteRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SmsTemplateDeleteRequest) GetParams() *SmsTemplateDeleteParam{
	return c.Param
}


type SmsTemplateDeleteParam struct {
	// 短信发送渠道，主要做资源隔离
	SmsAccount *string `json:"sms_account,omitempty"`
	// 短信模板id
	TemplateId *string `json:"template_id,omitempty"`
}
