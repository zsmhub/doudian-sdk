package logistics_getShopKey_request

import (
	"doudian.com/open/sdk_golang/api/logistics_getShopKey/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type LogisticsGetShopKeyRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *LogisticsGetShopKeyParam 
}
func (c *LogisticsGetShopKeyRequest) GetUrlPath() string{
	return "/logistics/getShopKey"
}


func New() *LogisticsGetShopKeyRequest{
	request := &LogisticsGetShopKeyRequest{
		Param: &LogisticsGetShopKeyParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *LogisticsGetShopKeyRequest) Execute(accessToken *doudian_sdk.AccessToken) (*logistics_getShopKey_response.LogisticsGetShopKeyResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &logistics_getShopKey_response.LogisticsGetShopKeyResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *LogisticsGetShopKeyRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *LogisticsGetShopKeyRequest) GetParams() *LogisticsGetShopKeyParam{
	return c.Param
}


type LogisticsGetShopKeyParam struct {
	// 打印密文
	CipherText string `json:"cipher_text,omitempty"`
	// 设备信息
	DeviceInfo string `json:"deviceInfo,omitempty"`
}
