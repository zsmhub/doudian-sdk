package sku_syncStock_request

import (
	"doudian.com/open/sdk_golang/api/sku_syncStock/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SkuSyncStockRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SkuSyncStockParam 
}
func (c *SkuSyncStockRequest) GetUrlPath() string{
	return "/sku/syncStock"
}


func New() *SkuSyncStockRequest{
	request := &SkuSyncStockRequest{
		Param: &SkuSyncStockParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SkuSyncStockRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sku_syncStock_response.SkuSyncStockResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sku_syncStock_response.SkuSyncStockResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SkuSyncStockRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SkuSyncStockRequest) GetParams() *SkuSyncStockParam{
	return c.Param
}


type SkuSyncStockParam struct {
	// sku编码
	Code string `json:"code,omitempty"`
	// 规格id；抖店系统生成，商品id下唯一。
	SkuId int64 `json:"sku_id,omitempty"`
	// 外部skuid，开发者在创建或更新商品时设置。
	OutSkuId int64 `json:"out_sku_id,omitempty"`
	// 商品ID，抖店系统生成，店铺下唯一；长度19位。
	ProductId int64 `json:"product_id,omitempty"`
	// 外部商品id；开发者在创建或更新商品时设置的。
	OutProductId int64 `json:"out_product_id,omitempty"`
	// 外部仓库id编码，开发者自定义；如传入传值表示更新区域库存，不传默认更新普通库存。可以使用【/warehouse/list】接口响应参数【out_warehouse_id】字段获取。超市小时达店铺类型，需传入store_id字段，store_id可以使用【/shop/batchCreateStore】创建成功后获取
	OutWarehouseId string `json:"out_warehouse_id,omitempty"`
	// 供应商ID
	SupplierId string `json:"supplier_id,omitempty"`
	// 库存更新方式；true-增量更新，同时idempotent_id字段必填；false-全量更新；默认为false
	Incremental bool `json:"incremental,omitempty"`
	// 幂等ID，当incremental=true时该参数必传
	IdempotentId string `json:"idempotent_id,omitempty"`
	// 库存值；可以设置为0；
	StockNum *int64 `json:"stock_num,omitempty"`
	// 阶梯库存
	StepStockNum int64 `json:"step_stock_num,omitempty"`
}
