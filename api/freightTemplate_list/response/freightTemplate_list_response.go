package freightTemplate_list_response

import (
	"doudian.com/open/sdk_golang/core"
)

type FreightTemplateListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *FreightTemplateListData `json:"data"`
}
type Template struct {
	// 运费模板id，可用于商品发布接口使用
	Id int64 `json:"id"`
	// 运费模板名称
	TemplateName string `json:"template_name"`
	// 发货省份id
	ProductProvince string `json:"product_province"`
	// 发货城市id
	ProductCity string `json:"product_city"`
	// 计价方式-1.按重量计价 2.按数量计价
	CalculateType int64 `json:"calculate_type"`
	// 快递方式-1.快递 目前仅支持1
	TransferType int64 `json:"transfer_type"`
	// 模板类型-0:阶梯计价 1:固定运费 2:卖家包邮 3:货到付款
	RuleType int64 `json:"rule_type"`
	// 固定运费金额(单位:分) 固定运费模板必填 1-9900之间的整数
	FixedAmount int64 `json:"fixed_amount"`
}
type ListItem struct {
	// 运费模版
	Template *Template `json:"template"`
}
type FreightTemplateListData struct {
	// 运费模版列表
	List []ListItem `json:"List"`
	// 总数
	Count int64 `json:"Count"`
}
