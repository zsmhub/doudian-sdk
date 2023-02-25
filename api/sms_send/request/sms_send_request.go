package sms_send_request

import (
	"doudian.com/open/sdk_golang/api/sms_send/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SmsSendRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SmsSendParam 
}
func (c *SmsSendRequest) GetUrlPath() string{
	return "/sms/send"
}


func New() *SmsSendRequest{
	request := &SmsSendRequest{
		Param: &SmsSendParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SmsSendRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sms_send_response.SmsSendResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sms_send_response.SmsSendResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SmsSendRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SmsSendRequest) GetParams() *SmsSendParam{
	return c.Param
}


type SmsSendParam struct {
	// 短信发送渠道，主要做资源隔离
	SmsAccount *string `json:"sms_account,omitempty"`
	// 签名
	Sign *string `json:"sign,omitempty"`
	// 短信模版id
	TemplateId *string `json:"template_id,omitempty"`
	// 短信模板占位符要替换的值
	TemplateParam *string `json:"template_param,omitempty"`
	// 透传字段，回执的时候原样返回给调用方，最大长度512字符
	Tag string `json:"tag,omitempty"`
	// 既支持手机号明文，又支持手机号密文
	PostTel *string `json:"post_tel,omitempty"`
	// 用户自定义扩展码，仅当允许自定义扩展码的时候生效
	UserExtCode string `json:"user_ext_code,omitempty"`
}
