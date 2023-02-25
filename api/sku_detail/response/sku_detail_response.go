package sku_detail_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SkuDetailResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SkuDetailData `json:"data"`
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
type SkuDetailData struct {
	// sku id
	Id int64 `json:"id"`
	// app_key
	OpenUserId int64 `json:"open_user_id"`
	// 外部的skuId
	OutSkuId int64 `json:"out_sku_id"`
	// 商品ID
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
	// 价格
	Price int64 `json:"price"`
	// 结算价格
	SettlementPrice int64 `json:"settlement_price"`
	// 规格ID
	SpecId int64 `json:"spec_id"`
	// 创建时间
	CreateTime int64 `json:"create_time"`
	// sku外部编码
	Code string `json:"code"`
	// 0-普通 1-区域库存
	SkuType int64 `json:"sku_type"`
	// 供应商ID
	SupplierId string `json:"supplier_id"`
	// sku_type=0时，表示库存数量 sku_type=1时，使用stock_map
	StockNum int64 `json:"stock_num"`
	// sku_type=0时，表示预占库存数量 sku_type=1时，使用prehold_stock_map
	PreholdStockNum int64 `json:"prehold_stock_num"`
	// 活动库存
	PromStockNum int64 `json:"prom_stock_num"`
	// 阶梯库存
	StepStockNum int64 `json:"step_stock_num"`
	// 预占阶梯库存
	PreholdStepStockNum int64 `json:"prehold_step_stock_num"`
	// 活动阶梯库存
	PromStepStockNum int64 `json:"prom_step_stock_num"`
	// 商品 ID 字符串
	ProductIdStr string `json:"product_id_str"`
	// 是否套装，0 否，1 是
	IsSuit int32 `json:"is_suit"`
	// 套装内数量
	SuitNum int64 `json:"suit_num"`
	// 容量单位ml，酒类专有
	Volume int64 `json:"volume"`
}
