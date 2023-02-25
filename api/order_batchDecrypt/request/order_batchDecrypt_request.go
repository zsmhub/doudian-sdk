package order_batchDecrypt_request

import (
	"doudian.com/open/sdk_golang/api/order_batchDecrypt/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderBatchDecryptRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderBatchDecryptParam 
}
func (c *OrderBatchDecryptRequest) GetUrlPath() string{
	return "/order/batchDecrypt"
}


func New() *OrderBatchDecryptRequest{
	request := &OrderBatchDecryptRequest{
		Param: &OrderBatchDecryptParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderBatchDecryptRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_batchDecrypt_response.OrderBatchDecryptResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_batchDecrypt_response.OrderBatchDecryptResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderBatchDecryptRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderBatchDecryptRequest) GetParams() *OrderBatchDecryptParam{
	return c.Param
}


type CipherInfosItem struct {
	// 订单号
	AuthId *string `json:"auth_id,omitempty"`
	// 待解密值
	CipherText *string `json:"cipher_text,omitempty"`
}
type OrderBatchDecryptParam struct {
	// 待解密值集合，最大支持一次解密50条。待解密的密文列表示例，入参结构{"cipher_infos":[{"auth_id”:”订单号”,”cipher_text”:”待解密值”},{“auth_id”:”订单号”,”cipher_text”:”待解密值”}]}
	CipherInfos *[]CipherInfosItem `json:"cipher_infos,omitempty"`
	// 服务商账号体系中，商户的账户ID，每个ISV下需要保证唯一，可选格式:独立生成的账户唯一标识
	AccountId string `json:"account_id,omitempty"`
	// 商户的账户ID类型；服务商账号中的主-main_account；子账号-sub_account；
	AccountType string `json:"account_type,omitempty"`
}
