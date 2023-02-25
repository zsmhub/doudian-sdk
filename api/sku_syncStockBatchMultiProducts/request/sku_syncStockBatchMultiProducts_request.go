package sku_syncStockBatchMultiProducts_request

import (
	"doudian.com/open/sdk_golang/api/sku_syncStockBatchMultiProducts/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SkuSyncStockBatchMultiProductsRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SkuSyncStockBatchMultiProductsParam 
}
func (c *SkuSyncStockBatchMultiProductsRequest) GetUrlPath() string{
	return "/sku/syncStockBatchMultiProducts"
}


func New() *SkuSyncStockBatchMultiProductsRequest{
	request := &SkuSyncStockBatchMultiProductsRequest{
		Param: &SkuSyncStockBatchMultiProductsParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SkuSyncStockBatchMultiProductsRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sku_syncStockBatchMultiProducts_response.SkuSyncStockBatchMultiProductsResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sku_syncStockBatchMultiProducts_response.SkuSyncStockBatchMultiProductsResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SkuSyncStockBatchMultiProductsRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SkuSyncStockBatchMultiProductsRequest) GetParams() *SkuSyncStockBatchMultiProductsParam{
	return c.Param
}


type ItemsItem struct {
	// 商品ID
	ProductId *int64 `json:"product_id,omitempty"`
	// 商品规格ID；一次请求最大支持50个sku更新
	SkuId *int64 `json:"sku_id,omitempty"`
	// 外部仓库id编码，开发者自定义；如传入传值表示更新区域库存，不传默认更新普通库存。可以使用【/warehouse/list】接口响应参数【out_warehouse_id】字段获取。超市小时达店铺类型，需传入store_id字段，store_id可以使用【/shop/batchCreateStore】创建成功后获取
	OutWarehouseId string `json:"out_warehouse_id,omitempty"`
	// 现货库存数
	StockNum int64 `json:"stock_num,omitempty"`
	// 阶梯库存数
	StepStockNum int64 `json:"step_stock_num,omitempty"`
}
type SkuSyncStockBatchMultiProductsParam struct {
	// 请求列表
	Items *[]ItemsItem `json:"items,omitempty"`
	// 幂等ID
	IdempotentId *string `json:"idempotent_id,omitempty"`
	// 库存更新方式，必填字段；true-增量更新，同时stock_num和step_stock_num字段可以为正数或负数；false-全量更新，同时stock_num和step_stock_num字段只能为正数数
	Incremental *bool `json:"incremental,omitempty"`
	// 请求来源，开发者自定义；
	Source string `json:"source,omitempty"`
}
