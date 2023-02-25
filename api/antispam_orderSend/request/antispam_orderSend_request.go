package antispam_orderSend_request

import (
	"doudian.com/open/sdk_golang/api/antispam_orderSend/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type AntispamOrderSendRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *AntispamOrderSendParam 
}
func (c *AntispamOrderSendRequest) GetUrlPath() string{
	return "/antispam/orderSend"
}


func New() *AntispamOrderSendRequest{
	request := &AntispamOrderSendRequest{
		Param: &AntispamOrderSendParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *AntispamOrderSendRequest) Execute(accessToken *doudian_sdk.AccessToken) (*antispam_orderSend_response.AntispamOrderSendResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &antispam_orderSend_response.AntispamOrderSendResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *AntispamOrderSendRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *AntispamOrderSendRequest) GetParams() *AntispamOrderSendParam{
	return c.Param
}


type User struct {
	// 用户类型
	UidType int32 `json:"uid_type,omitempty"`
	// 用户 ID
	UserId int64 `json:"user_id,omitempty"`
}
type AntispamOrderSendParam struct {
	// 事件时间
	EventTime *int64 `json:"event_time,omitempty"`
	// 用户
	User *User `json:"user,omitempty"`
	// 可变参数
	Params *string `json:"params,omitempty"`
}
