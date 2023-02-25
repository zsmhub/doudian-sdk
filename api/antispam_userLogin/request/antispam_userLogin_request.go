package antispam_userLogin_request

import (
	"doudian.com/open/sdk_golang/api/antispam_userLogin/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type AntispamUserLoginRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *AntispamUserLoginParam 
}
func (c *AntispamUserLoginRequest) GetUrlPath() string{
	return "/antispam/userLogin"
}


func New() *AntispamUserLoginRequest{
	request := &AntispamUserLoginRequest{
		Param: &AntispamUserLoginParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *AntispamUserLoginRequest) Execute(accessToken *doudian_sdk.AccessToken) (*antispam_userLogin_response.AntispamUserLoginResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &antispam_userLogin_response.AntispamUserLoginResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *AntispamUserLoginRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *AntispamUserLoginRequest) GetParams() *AntispamUserLoginParam{
	return c.Param
}


type User struct {
	// 用户类型
	UidType int32 `json:"uid_type,omitempty"`
	// 用户 ID
	UserId int64 `json:"user_id,omitempty"`
}
type AntispamUserLoginParam struct {
	// json 字符串
	Params *string `json:"params,omitempty"`
	// 事件发生的时间
	EventTime *int64 `json:"event_time,omitempty"`
	// 用户
	User *User `json:"user,omitempty"`
}
