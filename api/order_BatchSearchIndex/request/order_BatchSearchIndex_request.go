package order_BatchSearchIndex_request

import (
	"doudian.com/open/sdk_golang/api/order_BatchSearchIndex/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderBatchSearchIndexRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderBatchSearchIndexParam 
}
func (c *OrderBatchSearchIndexRequest) GetUrlPath() string{
	return "/order/BatchSearchIndex"
}


func New() *OrderBatchSearchIndexRequest{
	request := &OrderBatchSearchIndexRequest{
		Param: &OrderBatchSearchIndexParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderBatchSearchIndexRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_BatchSearchIndex_response.OrderBatchSearchIndexResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_BatchSearchIndex_response.OrderBatchSearchIndexResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderBatchSearchIndexRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderBatchSearchIndexRequest) GetParams() *OrderBatchSearchIndexParam{
	return c.Param
}


type PlainTextListItem struct {
	// 明文
	PlainText *string `json:"plain_text,omitempty"`
	// 加密类型；1地址加密 2姓名加密 3电话加密
	EncryptType *int16 `json:"encrypt_type,omitempty"`
}
type OrderBatchSearchIndexParam struct {
	// 明文列表
	PlainTextList *[]PlainTextListItem `json:"plain_text_list,omitempty"`
}
