package order_batchEncrypt_request

import (
	"doudian.com/open/sdk_golang/api/order_batchEncrypt/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderBatchEncryptRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderBatchEncryptParam 
}
func (c *OrderBatchEncryptRequest) GetUrlPath() string{
	return "/order/batchEncrypt"
}


func New() *OrderBatchEncryptRequest{
	request := &OrderBatchEncryptRequest{
		Param: &OrderBatchEncryptParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderBatchEncryptRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_batchEncrypt_response.OrderBatchEncryptResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_batchEncrypt_response.OrderBatchEncryptResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderBatchEncryptRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderBatchEncryptRequest) GetParams() *OrderBatchEncryptParam{
	return c.Param
}


type BatchEncryptListItem struct {
	// 明文
	PlainText *string `json:"plain_text,omitempty"`
	// 业务标识，value为抖音订单号；若是三方订单，可用三方订单号作为标识或自定义标识
	AuthId *string `json:"auth_id,omitempty"`
	// 是否支持密文索引
	IsSupportIndex *bool `json:"is_support_index,omitempty"`
	// 加密类型；1地址加密 2姓名加密 3电话加密
	SensitiveType *int16 `json:"sensitive_type,omitempty"`
}
type OrderBatchEncryptParam struct {
	// 待加密列表
	BatchEncryptList *[]BatchEncryptListItem `json:"batch_encrypt_list,omitempty"`
	// 加密场景：OrderCode代表订单 AftersaleCode代表售后单
	SensitiveScene string `json:"sensitive_scene,omitempty"`
}
