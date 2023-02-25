package sms_public_template_request

import (
	"doudian.com/open/sdk_golang/api/sms_public_template/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SmsPublicTemplateRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SmsPublicTemplateParam 
}
func (c *SmsPublicTemplateRequest) GetUrlPath() string{
	return "/sms/public/template"
}


func New() *SmsPublicTemplateRequest{
	request := &SmsPublicTemplateRequest{
		Param: &SmsPublicTemplateParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SmsPublicTemplateRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sms_public_template_response.SmsPublicTemplateResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sms_public_template_response.SmsPublicTemplateResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SmsPublicTemplateRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SmsPublicTemplateRequest) GetParams() *SmsPublicTemplateParam{
	return c.Param
}


type SmsPublicTemplateParam struct {
	// 每页数据大小
	Size int64 `json:"size,omitempty"`
	// 第几页，从0开始
	Page int64 `json:"page,omitempty"`
	// 模版id
	TemplateId string `json:"template_id,omitempty"`
}
