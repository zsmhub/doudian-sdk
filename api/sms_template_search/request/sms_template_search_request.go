package sms_template_search_request

import (
	"doudian.com/open/sdk_golang/api/sms_template_search/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SmsTemplateSearchRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SmsTemplateSearchParam 
}
func (c *SmsTemplateSearchRequest) GetUrlPath() string{
	return "/sms/template/search"
}


func New() *SmsTemplateSearchRequest{
	request := &SmsTemplateSearchRequest{
		Param: &SmsTemplateSearchParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SmsTemplateSearchRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sms_template_search_response.SmsTemplateSearchResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sms_template_search_response.SmsTemplateSearchResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SmsTemplateSearchRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SmsTemplateSearchRequest) GetParams() *SmsTemplateSearchParam{
	return c.Param
}


type SmsTemplateSearchParam struct {
	// 短信发送渠道，主要做资源隔离
	SmsAccount *string `json:"sms_account,omitempty"`
	// 短信模板id
	TemplateId string `json:"template_id,omitempty"`
	// 短信模版内容
	TemplateContent string `json:"template_content,omitempty"`
	// 页码，默认0
	Page int64 `json:"page,omitempty"`
	// 每页大小，默认10
	Size int64 `json:"size,omitempty"`
	// 短信模版名称
	TemplateName string `json:"template_name,omitempty"`
}
