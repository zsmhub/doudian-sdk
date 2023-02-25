package sku_detail_request

import (
	"doudian.com/open/sdk_golang/api/sku_detail/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SkuDetailRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SkuDetailParam 
}
func (c *SkuDetailRequest) GetUrlPath() string{
	return "/sku/detail"
}


func New() *SkuDetailRequest{
	request := &SkuDetailRequest{
		Param: &SkuDetailParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SkuDetailRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sku_detail_response.SkuDetailResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sku_detail_response.SkuDetailResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SkuDetailRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SkuDetailRequest) GetParams() *SkuDetailParam{
	return c.Param
}


type SkuDetailParam struct {
	// sku id
	SkuId int64 `json:"sku_id,omitempty"`
	// sku code
	Code string `json:"code,omitempty"`
}
