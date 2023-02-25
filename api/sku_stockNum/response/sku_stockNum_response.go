package sku_stockNum_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SkuStockNumResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SkuStockNumData `json:"data"`
}
type ShipRuleMapItem struct {
}
type SkuStockNumData struct {
	// 商品规格id，店铺商品id下唯一，抖店平台生成
	SkuId int64 `json:"sku_id"`
	// 外部开发者自定义商品规格id，会校验再店铺下唯一性
	OutSkuId int64 `json:"out_sku_id"`
	// sku商家编码，对应抖店后台商品sku商家编码
	Code string `json:"code"`
	// sku库存类型，0-普通库存；1-区域库存；10-阶梯库存；
	SkuType int64 `json:"sku_type"`
	// 商品规格库存；如果sku_type=0，表示全国库存； 如果sku_type=1，且入参有out_warehouse_id，则表示该仓库的库存； 如果sku_type=1，且入参无out_warehouse_id，则为空；
	StockNum int64 `json:"stock_num"`
	// 占用库存
	PreholdStockNum int64 `json:"prehold_stock_num"`
	// 活动库存
	PromStockNum int64 `json:"prom_stock_num"`
	// 阶梯库存
	StepStockNum int64 `json:"step_stock_num"`
	// 阶梯占用库存
	PreholdStepStockNum int64 `json:"prehold_step_stock_num"`
	// 活动阶梯库存
	PromStepStockNum int64 `json:"prom_step_stock_num"`
	// 如果sku_type=0，为空 如果sku_type=1，则为区域仓库存映射表，key为out_warehouse_id，value为库存
	StockNumMap map[string]int64 `json:"stock_num_map"`
	// 如果sku_type=0，为空 如果sku_type=1，则为区域仓库存映射表，key为out_warehouse_id，value为占用库存
	PreholdStockMap map[string]int64 `json:"prehold_stock_map"`
	// 如果sku_type=0，为空 如果sku_type=1，则为区域仓库存映射表，key为out_warehouse_id，value为sku在仓中的发货时效
	ShipRuleMap map[string]ShipRuleMapItem `json:"ship_rule_map"`
	// 普通库存，非活动可售
	NormalStockNum int64 `json:"normal_stock_num"`
	// 渠道库存
	ChannelStockNum int64 `json:"channel_stock_num"`
}
