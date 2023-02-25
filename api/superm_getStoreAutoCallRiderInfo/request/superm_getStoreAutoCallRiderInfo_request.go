package superm_getStoreAutoCallRiderInfo_request

import (
	"doudian.com/open/sdk_golang/api/superm_getStoreAutoCallRiderInfo/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SupermGetStoreAutoCallRiderInfoRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SupermGetStoreAutoCallRiderInfoParam 
}
func (c *SupermGetStoreAutoCallRiderInfoRequest) GetUrlPath() string{
	return "/superm/getStoreAutoCallRiderInfo"
}


func New() *SupermGetStoreAutoCallRiderInfoRequest{
	request := &SupermGetStoreAutoCallRiderInfoRequest{
		Param: &SupermGetStoreAutoCallRiderInfoParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SupermGetStoreAutoCallRiderInfoRequest) Execute(accessToken *doudian_sdk.AccessToken) (*superm_getStoreAutoCallRiderInfo_response.SupermGetStoreAutoCallRiderInfoResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &superm_getStoreAutoCallRiderInfo_response.SupermGetStoreAutoCallRiderInfoResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SupermGetStoreAutoCallRiderInfoRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SupermGetStoreAutoCallRiderInfoRequest) GetParams() *SupermGetStoreAutoCallRiderInfoParam{
	return c.Param
}


type SupermGetStoreAutoCallRiderInfoParam struct {
	// 商家的门店编号
	StoreId *int64 `json:"store_id,omitempty"`
}
