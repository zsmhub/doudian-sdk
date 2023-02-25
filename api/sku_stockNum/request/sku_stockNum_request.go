package sku_stockNum_request

import (
	"doudian.com/open/sdk_golang/api/sku_stockNum/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SkuStockNumRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SkuStockNumParam 
}
func (c *SkuStockNumRequest) GetUrlPath() string{
	return "/sku/stockNum"
}


func New() *SkuStockNumRequest{
	request := &SkuStockNumRequest{
		Param: &SkuStockNumParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SkuStockNumRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sku_stockNum_response.SkuStockNumResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sku_stockNum_response.SkuStockNumResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SkuStockNumRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SkuStockNumRequest) GetParams() *SkuStockNumParam{
	return c.Param
}


type SkuStockNumParam struct {
	// 商品规格id，店铺商品id下唯一，抖店平台生成
	SkuId int64 `json:"sku_id,omitempty"`
	// sku商家编码，对应抖店后台商品sku商家编码。外部开发者自定义商品规格id，会校验在店铺下唯一性
	Code string `json:"code,omitempty"`
	// 仓库编码（供应商编码）
	OutWarehouseId string `json:"out_warehouse_id,omitempty"`
}
