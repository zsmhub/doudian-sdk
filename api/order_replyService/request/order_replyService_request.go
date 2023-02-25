package order_replyService_request

import (
	"doudian.com/open/sdk_golang/api/order_replyService/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderReplyServiceRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderReplyServiceParam 
}
func (c *OrderReplyServiceRequest) GetUrlPath() string{
	return "/order/replyService"
}


func New() *OrderReplyServiceRequest{
	request := &OrderReplyServiceRequest{
		Param: &OrderReplyServiceParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderReplyServiceRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_replyService_response.OrderReplyServiceResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_replyService_response.OrderReplyServiceResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderReplyServiceRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderReplyServiceRequest) GetParams() *OrderReplyServiceParam{
	return c.Param
}


type OrderReplyServiceParam struct {
	// 服务请求列表中获取的id
	Id *int64 `json:"id,omitempty"`
	// 回复内容
	Reply *string `json:"reply,omitempty"`
	// 回复凭证，通过/order/serviceDetail获取是否当前服务单</br>是否必须上传凭证。多张图片用竖线分开。不超过4张
	Evidence string `json:"evidence,omitempty"`
}
