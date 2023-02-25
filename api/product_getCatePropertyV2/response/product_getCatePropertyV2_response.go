package product_getCatePropertyV2_response

import (
	"doudian.com/open/sdk_golang/core"
)

type ProductGetCatePropertyV2Response struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *ProductGetCatePropertyV2Data `json:"data"`
}
type ProductGetCatePropertyV2Data struct {
	// 返回参数列表
	Data []DataItem `json:"data"`
	// 模板类型，SPU为1，结构化为2，默认为0（无特殊需求忽略该字段即可）
	TplType int64 `json:"tpl_type"`
}
type OptionsItem struct {
	// 可选值名称
	Name string `json:"name"`
	// 可选值
	Value string `json:"value"`
	// 值的id
	ValueId int64 `json:"value_id"`
	// 属性值顺序
	Sequence int64 `json:"sequence"`
}
type DataItem struct {
	// 【已废弃】老类目id，请使用category_id字段
	Cid int64 `json:"cid"`
	// 属性名称
	PropertyName string `json:"property_name"`
	// 属性可选值列表，为空时需要手动填写
	Options []OptionsItem `json:"options"`
	// 1：创建商品时该属性字段必填 0：创建商品时该属性字段选填
	Required int64 `json:"required"`
	// 属性状态，0：有效，1：失效
	Status int64 `json:"status"`
	// 输入text、单选select、多选multi_select、时间戳timestamp、时间段timerange
	Type string `json:"type"`
	// 新版类目id（优先使用）
	CategoryId int64 `json:"category_id"`
	// 最大多选数
	MultiSelectMax int64 `json:"multi_select_max"`
	// 属性类型，0 绑定属性 1关键属性 2售卖属性 3 商品属性
	PropertyType int64 `json:"property_type"`
	// 属性id
	PropertyId int64 `json:"property_id"`
	// 属性顺序
	Sequence int64 `json:"sequence"`
	// 关系id,SPU类目使用，表示自己的上一个关键属性ID
	RelationId int64 `json:"relation_id"`
	// 商品属性是否有支持商家自定义，1=支持自定义，0=不支持自定义。 使用场景：当开发者传入自定义的属性值时需在创建或更新商品接口的属性时，需把【product_format_new】字段中传入diy_type=1并且value=0；
	DiyType int64 `json:"diy_type"`
	// 0:非重要属性，1:重要属性
	ImportantType int64 `json:"important_type"`
}
