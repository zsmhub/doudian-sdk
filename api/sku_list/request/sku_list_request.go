package sku_list_request

import (
	"doudian.com/open/sdk_golang/api/sku_list/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SkuListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SkuListParam 
}
func (c *SkuListRequest) GetUrlPath() string{
	return "/sku/list"
}


func New() *SkuListRequest{
	request := &SkuListRequest{
		Param: &SkuListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SkuListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sku_list_response.SkuListResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sku_list_response.SkuListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SkuListRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SkuListRequest) GetParams() *SkuListParam{
	return c.Param
}


type SkuListParam struct {
	// 商品ID；抖店系统生成。
	ProductId int64 `json:"product_id,omitempty"`
	// 外部商品ID；商家创建商品时自定义。
	OutProductId int64 `json:"out_product_id,omitempty"`
}
