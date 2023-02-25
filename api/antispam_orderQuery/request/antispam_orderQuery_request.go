package antispam_orderQuery_request

import (
	"doudian.com/open/sdk_golang/api/antispam_orderQuery/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type AntispamOrderQueryRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *AntispamOrderQueryParam 
}
func (c *AntispamOrderQueryRequest) GetUrlPath() string{
	return "/antispam/orderQuery"
}


func New() *AntispamOrderQueryRequest{
	request := &AntispamOrderQueryRequest{
		Param: &AntispamOrderQueryParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *AntispamOrderQueryRequest) Execute(accessToken *doudian_sdk.AccessToken) (*antispam_orderQuery_response.AntispamOrderQueryResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &antispam_orderQuery_response.AntispamOrderQueryResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *AntispamOrderQueryRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *AntispamOrderQueryRequest) GetParams() *AntispamOrderQueryParam{
	return c.Param
}


type User struct {
	// 用户类型
	UidType int32 `json:"uid_type,omitempty"`
	// 用户 ID
	UserId int64 `json:"user_id,omitempty"`
}
type AntispamOrderQueryParam struct {
	// 事件时间
	EventTime *int64 `json:"event_time,omitempty"`
	// 用户
	User *User `json:"user,omitempty"`
	// 上报参数
	Params *string `json:"params,omitempty"`
}
