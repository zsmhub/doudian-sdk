package sms_batchSend_request

import (
	"doudian.com/open/sdk_golang/api/sms_batchSend/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SmsBatchSendRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SmsBatchSendParam 
}
func (c *SmsBatchSendRequest) GetUrlPath() string{
	return "/sms/batchSend"
}


func New() *SmsBatchSendRequest{
	request := &SmsBatchSendRequest{
		Param: &SmsBatchSendParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SmsBatchSendRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sms_batchSend_response.SmsBatchSendResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sms_batchSend_response.SmsBatchSendResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SmsBatchSendRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SmsBatchSendRequest) GetParams() *SmsBatchSendParam{
	return c.Param
}


type SmsMessageListItem struct {
	// 既支持手机号明文，又支持手机号密文
	PostTel *string `json:"post_tel,omitempty"`
	// 参数
	TemplateParam *string `json:"template_param,omitempty"`
}
type SmsBatchSendParam struct {
	// 短信发送渠道，主要做资源隔离
	SmsAccount *string `json:"sms_account,omitempty"`
	// 短信签名
	Sign *string `json:"sign,omitempty"`
	// 短信列表
	SmsMessageList *[]SmsMessageListItem `json:"sms_message_list,omitempty"`
	// 短信模板id
	TemplateId *string `json:"template_id,omitempty"`
	// 透传字段，回执的时候原样返回给调用方，最大长度512字符
	Tag string `json:"tag,omitempty"`
	// 用户自定义扩展码，仅当允许自定义扩展码的时候生效
	UserExtCode string `json:"user_ext_code,omitempty"`
}
