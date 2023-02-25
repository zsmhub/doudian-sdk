package sku_syncStockBatchMultiProducts_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SkuSyncStockBatchMultiProductsResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SkuSyncStockBatchMultiProductsData `json:"data"`
}
type ResultsItem struct {
	// 商品规格ID
	SkuId int64 `json:"sku_id"`
	// 仓ID
	WarehouseId int64 `json:"warehouse_id"`
	// 现货库存数
	StockNum int64 `json:"stock_num"`
	// 阶梯库存数
	StepStockNum int64 `json:"step_stock_num"`
	// 更改的现货库存数
	ChangeNum int64 `json:"change_num"`
	// 更改的阶梯库存数
	StepChangeNum int64 `json:"step_change_num"`
	// 状态码
	StatusCode int64 `json:"status_code"`
	// 状态信息
	StatusMessage string `json:"status_message"`
}
type SkuSyncStockBatchMultiProductsData struct {
	// 结果列表
	Results []ResultsItem `json:"results"`
}
