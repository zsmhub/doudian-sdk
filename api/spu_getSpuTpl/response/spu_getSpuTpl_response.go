package spu_getSpuTpl_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SpuGetSpuTplResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SpuGetSpuTplData `json:"data"`
}
type OptionsItem struct {
	// 属性值名称
	Name string `json:"name"`
	// 属性值id
	ValueId int64 `json:"value_id"`
	// 属性值顺序
	Sequence int64 `json:"sequence"`
}
type ProductFormatItem struct {
	// 属性名
	Name string `json:"name"`
	// 选项
	Options []OptionsItem `json:"options"`
	// 是否必填 1是
	Require int64 `json:"require"`
	// 类型 text：可输入，select：单选，multi_select: 多选
	Type string `json:"type"`
	// 类目id
	CategoryId int64 `json:"category_id"`
	// 多选最大选项
	MultiSelectMax int64 `json:"multi_select_max"`
	// 属性类型 0 绑定类型。1关键属性
	PropertyType int64 `json:"property_type"`
	// 属性id
	PropertyId int64 `json:"property_id"`
	// 顺序
	Sequence int64 `json:"sequence"`
}
type SpuGetSpuTplData struct {
	// 属性信息
	ProductFormat []ProductFormatItem `json:"product_format"`
}
