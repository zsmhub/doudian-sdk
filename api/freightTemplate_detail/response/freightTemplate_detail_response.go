package freightTemplate_detail_response

import (
	"doudian.com/open/sdk_golang/core"
)

type FreightTemplateDetailResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *Data `json:"data"`
}
type ProvinceInfosItem struct {
	// 地址id，第一级是省份、第二级是城市、第三级是区、第四级是街道
	Id string `json:"id"`
	// 下一级地址信息
	Children []ChildrenItem_5 `json:"children"`
}
type ColumnsItem struct {
	// 首重(单位:kg) 按重量计价必填 0.1-999.9之间的小数，小数点后一位
	FirstWeight float64 `json:"first_weight"`
	// 首重价格(单位:元) 按重量计价必填 0.00-30.00之间的小数，小数点后两位
	FirstWeightPrice float64 `json:"first_weight_price"`
	// 首件数量(单位:个) 按数量计价必填 1-999的整数
	FirstNum int64 `json:"first_num"`
	// 首件价格(单位:元)按数量计价必填 0.00-30.00之间的小数，小数点后两位
	FirstNumPrice float64 `json:"first_num_price"`
	// 续重(单位:kg) 按重量计价必填 0.1-999.9之间的小数，小数点后一位
	AddWeight float64 `json:"add_weight"`
	// 续重价格(单位:元) 按重量计价必填 0.00-30.00之间的小数，小数点后两位
	AddWeightPrice float64 `json:"add_weight_price"`
	// 续件(单位:个) 按数量计价必填 1-999的整数
	AddNum int64 `json:"add_num"`
	// 续件价格(单位:元) 按数量计价必填 0.00-30.00之间的小数，小数点后两位
	AddNumPrice float64 `json:"add_num_price"`
	// 是否默认计价方式(1:是；0:不是)
	IsDefault int64 `json:"is_default"`
	// 是否限运规则
	IsLimited bool `json:"is_limited"`
	// 是否包邮规则
	IsOverFree bool `json:"is_over_free"`
	// 满xx重量包邮(单位:kg)0.1-10.0之间的小数，小数点后一位
	OverWeight float64 `json:"over_weight"`
	// 满xx金额包邮(单位:分)10-99900的整数
	OverAmount int64 `json:"over_amount"`
	// 满xx件包邮 1-10之间的整数
	OverNum int64 `json:"over_num"`
	// 当前规则生效的地址，统一以List<Struct>结构返回，该结构为嵌套结构。对应的json格式为[{"id":"32","children":[{"id":"320500","children":[{"id":"320508","children":[{"id":"320508014"},{"id":"320508004"}]}]}]}] 注意：返回的为最新的四级地址版本（地址存储升级变更的可能，以最新的返回）
	ProvinceInfos []ProvinceInfosItem `json:"province_infos"`
}
type Data struct {
	// 模板信息
	Template *Template `json:"template"`
	// 规则
	Columns []ColumnsItem `json:"columns"`
}
type FreightTemplateDetailData struct {
	// 模板详情
	Data *Data `json:"data"`
}
type Template struct {
	// 模板id
	Id int64 `json:"id"`
	// 模板名称
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
type ChildrenItem struct {
	// 地址id，第一级是省份、第二级是城市、第三级是区、第四级是街道
	Id string `json:"id"`
}
type ChildrenItem_6 struct {
	// 地址id，第一级是省份、第二级是城市、第三级是区、第四级是街道
	Id string `json:"id"`
	// 下一级地址信息
	Children []ChildrenItem `json:"children"`
}
type ChildrenItem_5 struct {
	// 地址id，第一级是省份、第二级是城市、第三级是区、第四级是街道
	Id string `json:"id"`
	// 下一级地址信息
	Children []ChildrenItem_6 `json:"children"`
}
