package sku_editCode_request

import (
	"doudian.com/open/sdk_golang/api/sku_editCode/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SkuEditCodeRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SkuEditCodeParam 
}
func (c *SkuEditCodeRequest) GetUrlPath() string{
	return "/sku/editCode"
}


func New() *SkuEditCodeRequest{
	request := &SkuEditCodeRequest{
		Param: &SkuEditCodeParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SkuEditCodeRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sku_editCode_response.SkuEditCodeResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sku_editCode_response.SkuEditCodeResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SkuEditCodeRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SkuEditCodeRequest) GetParams() *SkuEditCodeParam{
	return c.Param
}


type SkuEditCodeParam struct {
	// 编码
	Code string `json:"code,omitempty"`
	// skuid
	SkuId int64 `json:"sku_id,omitempty"`
	// 外部skuid
	OutSkuId int64 `json:"out_sku_id,omitempty"`
	// 商品id
	ProductId int64 `json:"product_id,omitempty"`
	// 外部商品id
	OutProductId int64 `json:"out_product_id,omitempty"`
}
