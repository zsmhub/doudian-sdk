package sms_sign_apply_list_request

import (
	"doudian.com/open/sdk_golang/api/sms_sign_apply_list/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SmsSignApplyListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SmsSignApplyListParam 
}
func (c *SmsSignApplyListRequest) GetUrlPath() string{
	return "/sms/sign/apply/list"
}


func New() *SmsSignApplyListRequest{
	request := &SmsSignApplyListRequest{
		Param: &SmsSignApplyListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SmsSignApplyListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sms_sign_apply_list_response.SmsSignApplyListResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sms_sign_apply_list_response.SmsSignApplyListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SmsSignApplyListRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SmsSignApplyListRequest) GetParams() *SmsSignApplyListParam{
	return c.Param
}


type SmsSignApplyListParam struct {
	// 短信发送渠道，主要做资源隔离
	SmsAccount *string `json:"sms_account,omitempty"`
	// 搜索
	Like string `json:"like,omitempty"`
	// 页码，默认0
	Page int64 `json:"page,omitempty"`
	// 每页大小，默认10
	Size int64 `json:"size,omitempty"`
	// 申请单id
	SmsSignApplyId string `json:"sms_sign_apply_id,omitempty"`
}
