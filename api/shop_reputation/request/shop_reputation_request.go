package shop_reputation_request

import (
	"doudian.com/open/sdk_golang/api/shop_reputation/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ShopReputationRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ShopReputationParam 
}
func (c *ShopReputationRequest) GetUrlPath() string{
	return "/shop/reputation"
}


func New() *ShopReputationRequest{
	request := &ShopReputationRequest{
		Param: &ShopReputationParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ShopReputationRequest) Execute(accessToken *doudian_sdk.AccessToken) (*shop_reputation_response.ShopReputationResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &shop_reputation_response.ShopReputationResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ShopReputationRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *ShopReputationRequest) GetParams() *ShopReputationParam{
	return c.Param
}


type ShopReputationParam struct {
}
