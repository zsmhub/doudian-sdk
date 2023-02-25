package logistics_newCreateOrder_response

import (
	"doudian.com/open/sdk_golang/core"
)

type LogisticsNewCreateOrderResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *LogisticsNewCreateOrderData `json:"data"`
}
type ErrInfosItem struct {
	// 订单号
	OrderId string `json:"order_id"`
	// 包裹id
	PackId string `json:"pack_id"`
	// 错误码
	ErrCode string `json:"err_code"`
	// 错误信息
	ErrMsg string `json:"err_msg"`
	// 2；详情请看文档映射表
	OrderChannel string `json:"order_channel"`
}
type LogisticsNewCreateOrderData struct {
	// 电子面单信息列表
	EbillInfos []EbillInfosItem `json:"ebill_infos"`
	// 错误信息列表
	ErrInfos []ErrInfosItem `json:"err_infos"`
}
type EbillInfosItem struct {
	// 面单对应的物流商code
	Company string `json:"company"`
	// 成功的单子的提示信息
	HintMsg string `json:"hint_msg"`
	// 订单号
	OrderId string `json:"order_id"`
	// 包裹id
	PackId string `json:"pack_id"`
	// 运单号
	TrackNo string `json:"track_no"`
	// 分拣码（三段码）
	SortCode string `json:"sort_code"`
	// 集包地代码
	PackageCenterCode string `json:"package_center_code"`
	// 集包名称
	PackageCenterName string `json:"package_center_name"`
	// 大头笔编码
	ShortAddressCode string `json:"short_address_code"`
	// 大头笔名称
	ShortAddressName string `json:"short_address_name"`
	// 额外打印信息（众邮、京东、丰网使用），具体请参考[抖音电商电子面单对接文档](https://op.jinritemai.com/docs/guide-docs/33/338)附录4
	ExtraResp string `json:"extra_resp"`
	// 子母件列表英文逗号分隔
	SubWaybillCodes string `json:"sub_waybill_codes"`
	// 2；详情请看文档映射表
	OrderChannel string `json:"order_channel"`
	// 快递商侧系统生成的寄件码
	ShippingCode string `json:"shipping_code"`
}
