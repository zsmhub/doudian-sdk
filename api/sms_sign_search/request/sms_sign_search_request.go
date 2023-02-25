package sms_sign_search_request

import (
	"doudian.com/open/sdk_golang/api/sms_sign_search/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SmsSignSearchRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SmsSignSearchParam 
}
func (c *SmsSignSearchRequest) GetUrlPath() string{
	return "/sms/sign/search"
}


func New() *SmsSignSearchRequest{
	request := &SmsSignSearchRequest{
		Param: &SmsSignSearchParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SmsSignSearchRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sms_sign_search_response.SmsSignSearchResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sms_sign_search_response.SmsSignSearchResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SmsSignSearchRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SmsSignSearchRequest) GetParams() *SmsSignSearchParam{
	return c.Param
}


type SmsSignSearchParam struct {
	// 短信发送渠道，主要做资源隔离
	SmsAccount *string `json:"sms_account,omitempty"`
	// 模糊搜索串
	Like string `json:"like,omitempty"`
	// 每页大小，默认10
	Size int64 `json:"size,omitempty"`
	// 页码，默认0
	Page int64 `json:"page,omitempty"`
}
