package order_batchSensitive_request

import (
	"doudian.com/open/sdk_golang/api/order_batchSensitive/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderBatchSensitiveRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderBatchSensitiveParam 
}
func (c *OrderBatchSensitiveRequest) GetUrlPath() string{
	return "/order/batchSensitive"
}


func New() *OrderBatchSensitiveRequest{
	request := &OrderBatchSensitiveRequest{
		Param: &OrderBatchSensitiveParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderBatchSensitiveRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_batchSensitive_response.OrderBatchSensitiveResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_batchSensitive_response.OrderBatchSensitiveResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderBatchSensitiveRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderBatchSensitiveRequest) GetParams() *OrderBatchSensitiveParam{
	return c.Param
}


type CipherInfosItem struct {
	// 业务标识，value为抖音订单号
	AuthId *string `json:"auth_id,omitempty"`
	// 密文
	CipherText *string `json:"cipher_text,omitempty"`
}
type OrderBatchSensitiveParam struct {
	// 待脱敏的密文列表，每次调用不超过50条
	CipherInfos *[]CipherInfosItem `json:"cipher_infos,omitempty"`
}
