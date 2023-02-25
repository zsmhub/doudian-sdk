package order_getSearchIndex_request

import (
	"doudian.com/open/sdk_golang/api/order_getSearchIndex/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderGetSearchIndexRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderGetSearchIndexParam 
}
func (c *OrderGetSearchIndexRequest) GetUrlPath() string{
	return "/order/getSearchIndex"
}


func New() *OrderGetSearchIndexRequest{
	request := &OrderGetSearchIndexRequest{
		Param: &OrderGetSearchIndexParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderGetSearchIndexRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_getSearchIndex_response.OrderGetSearchIndexResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_getSearchIndex_response.OrderGetSearchIndexResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderGetSearchIndexRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderGetSearchIndexRequest) GetParams() *OrderGetSearchIndexParam{
	return c.Param
}


type OrderGetSearchIndexParam struct {
	// 电话号码
	PlainText *string `json:"plain_text,omitempty"`
	// 加密类型；1地址加密 2姓名加密 3电话加密
	SensitiveType *int16 `json:"sensitive_type,omitempty"`
}
