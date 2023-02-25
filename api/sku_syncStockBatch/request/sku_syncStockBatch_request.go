package sku_syncStockBatch_request

import (
	"doudian.com/open/sdk_golang/api/sku_syncStockBatch/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SkuSyncStockBatchRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SkuSyncStockBatchParam 
}
func (c *SkuSyncStockBatchRequest) GetUrlPath() string{
	return "/sku/syncStockBatch"
}


func New() *SkuSyncStockBatchRequest{
	request := &SkuSyncStockBatchRequest{
		Param: &SkuSyncStockBatchParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SkuSyncStockBatchRequest) Execute(accessToken *doudian_sdk.AccessToken) (*sku_syncStockBatch_response.SkuSyncStockBatchResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &sku_syncStockBatch_response.SkuSyncStockBatchResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SkuSyncStockBatchRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SkuSyncStockBatchRequest) GetParams() *SkuSyncStockBatchParam{
	return c.Param
}


type StockMapItem struct {
	// 仓id，超市小时达子品传storeId
	OutWarehouseId *string `json:"out_warehouse_id,omitempty"`
	// 库存数量
	StockNum *int64 `json:"stock_num,omitempty"`
}
type SkuSyncListItem struct {
	// skuid
	SkuId *int64 `json:"sku_id,omitempty"`
	// 库存类型：0普通库存，1区域库存,10阶梯库存。 超市小时达子品传区域库存
	SkuType *int64 `json:"sku_type,omitempty"`
	// 普通库存数量
	StockNum int64 `json:"stock_num,omitempty"`
	// 阶梯库存数量
	StepStockNum int64 `json:"step_stock_num,omitempty"`
	// 区域库存配置
	StockMap []StockMapItem `json:"stock_map,omitempty"`
	// 供应商ID
	SupplierId string `json:"supplier_id,omitempty"`
}
type SkuSyncStockBatchParam struct {
	// 商品ID，整型格式
	ProductId *int64 `json:"product_id,omitempty"`
	// 幂等ID，仅incremental=true时有用
	IdempotentId string `json:"idempotent_id,omitempty"`
	// true表示增量库存，false表示全量库存，默认为false
	Incremental *bool `json:"incremental,omitempty"`
	// 需要批量提交的skuid及对应的库存和仓；接口规则全部成功或全部失败，例：批量更新10个skuid库存，其中一个skuid信息不正确，这样整个请求都会失败，10个skuid都未更新成功。
	SkuSyncList *[]SkuSyncListItem `json:"sku_sync_list,omitempty"`
}
