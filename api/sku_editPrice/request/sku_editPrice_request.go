package sku_editPrice_request

import (
	"doudian.com/open/sdk_golang/api/sku_editPrice/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SkuEditPriceRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SkuEditPriceParam 
}
func (c *SkuEditPriceRequest) GetUrlPath() string{
	return "/sku/editPrice"
}


func New() *SkuEditPriceRequest{
	request := &SkuEditPriceRequest{
		Param: &SkuEditPriceParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SkuEditPriceRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sku_editPrice_response.SkuEditPriceResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sku_editPrice_response.SkuEditPriceResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SkuEditPriceRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SkuEditPriceRequest) GetParams() *SkuEditPriceParam{
	return c.Param
}


type SkuEditPriceParam struct {
	// 售价 (单位 分)
	Price *int64 `json:"price,omitempty"`
	// sku编码
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
