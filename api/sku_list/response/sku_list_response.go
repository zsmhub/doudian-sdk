package sku_list_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SkuListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data []DataItem `json:"data"`
}
type SkuListData struct {
	// sku列表
	Data []DataItem `json:"data"`
}
type CustomsReportInfo struct {
	// 海关代码
	HsCode string `json:"hs_code"`
	// 法定第一计量数量
	FirstMeasureQty float64 `json:"first_measure_qty"`
	// 法定第二计量数量
	SecondMeasureQty float64 `json:"second_measure_qty"`
	// 法定第一计量单位
	FirstMeasureUnit string `json:"first_measure_unit"`
	// 法定第二计量单位
	SecondMeasureUnit string `json:"second_measure_unit"`
	// 售卖单位
	Unit string `json:"unit"`
	// 品名
	ReportName string `json:"report_name"`
	// 品牌
	ReportBrandName string `json:"report_brand_name"`
	// 用途
	Usage string `json:"usage"`
	// 规格型号
	GModel string `json:"g_model"`
	// 条形码
	BarCode string `json:"bar_code"`
}
type ShipRuleMapItem struct {
	// 0 现货模式  1 全款预售模式
	PresellType int64 `json:"presell_type"`
	// 发货延迟时间；0表示当天发货；
	DelayDay int64 `json:"delay_day"`
	// 全款预售截止时间  presell_type = 1时有用
	PresellEndTime int64 `json:"presell_end_time"`
}
type SellPropertiesItem struct {
	// 规格值id
	ValueSpecDetailId int64 `json:"value_spec_detail_id"`
	// 规格值名称
	ValueName string `json:"value_name"`
}
type DataItem struct {
	// 商品sku_id;抖店系统生成。
	Id int64 `json:"id"`
	// app_key
	OpenUserId int64 `json:"open_user_id"`
	// 外部的skuId；商家自定义；未设置返回为0。
	OutSkuId int64 `json:"out_sku_id"`
	// 商品ID；抖店系统生成。
	ProductId int64 `json:"product_id"`
	// 第一级子规格
	SpecDetailId1 int64 `json:"spec_detail_id1"`
	// 第二级子规格
	SpecDetailId2 int64 `json:"spec_detail_id2"`
	// 第三级子规格
	SpecDetailId3 int64 `json:"spec_detail_id3"`
	// 第一级子规格名
	SpecDetailName1 string `json:"spec_detail_name1"`
	// 第二级子规格名
	SpecDetailName2 string `json:"spec_detail_name2"`
	// 第三级子规格名
	SpecDetailName3 string `json:"spec_detail_name3"`
	// 海关申报要素
	CustomsReportInfo *CustomsReportInfo `json:"customs_report_info"`
	// 商品价格，单位：分
	Price int64 `json:"price"`
	// 结算价格，【已废弃】价格字段是price
	SettlementPrice int64 `json:"settlement_price"`
	// 规格ID
	SpecId int64 `json:"spec_id"`
	// 创建时间，时间戳：单位秒；
	CreateTime int64 `json:"create_time"`
	// sku外部编码
	Code string `json:"code"`
	// 库存类型；0-普通；1-区域库存；10-阶梯库存；
	SkuType int64 `json:"sku_type"`
	// 供应商ID
	SupplierId string `json:"supplier_id"`
	// sku_type=0时，表示普通库存数量 ;sku_type=1时，使用stock_map，表示区域库存数量
	StockNum int64 `json:"stock_num"`
	// sku_type=0时，表示预占库存数量； sku_type=1时，表示区域库存数量，使用prehold_stock_map
	PreholdStockNum int64 `json:"prehold_stock_num"`
	// 活动库存
	PromStockNum int64 `json:"prom_stock_num"`
	// 阶梯库存
	StepStockNum int64 `json:"step_stock_num"`
	// 预占阶梯库存
	PreholdStepStockNum int64 `json:"prehold_step_stock_num"`
	// 活动阶梯库存
	PromStepStockNum int64 `json:"prom_step_stock_num"`
	// 如果sku_type=0，为空 如果sku_type=1，则为区域仓库存映射表，key为out_warehouse_id（区域仓库id使用【/warehouse/list】查询），value为库存
	StockMap map[string]int64 `json:"stock_map"`
	// 如果sku_type=0，为空 如果sku_type=1，则为区域仓库存映射表，key为out_warehouse_id（区域仓库id使用【/warehouse/list】查询），value为占用库存
	PreholdStockMap map[string]int64 `json:"prehold_stock_map"`
	// 商品 ID字符串类型
	ProductIdStr string `json:"product_id_str"`
	// 如果sku_type=0，为空 如果sku_type=1，则为区域仓库存映射表，key为out_warehouse_id（区域仓库id使用【/warehouse/list】查询），value为sku 在对应仓中的发货时效
	ShipRuleMap map[string]ShipRuleMapItem `json:"ship_rule_map"`
	// 海南免税：是否套装，0 否，1 是
	IsSuit int32 `json:"is_suit"`
	// 海南免税：套装数量
	SuitNum int64 `json:"suit_num"`
	// 海南免税：净含量
	Volume int64 `json:"volume"`
	// 库存模型V2新增 普通库存 非活动可售
	NormalStockNum int64 `json:"normal_stock_num"`
	// 库存模型V2新增 渠道库存
	ChannelStockNum int64 `json:"channel_stock_num"`
	// 销售属性，代替spec_detail_id123、spec_detail_name123
	SellProperties []SellPropertiesItem `json:"sell_properties"`
}
