package spu_getSpuRule_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SpuGetSpuRuleResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SpuGetSpuRuleData `json:"data"`
}
type SpuActualImagesRule struct {
	// 是否必填
	IsRequired int64 `json:"is_required"`
	// 最大数量
	MaxNum int64 `json:"max_num"`
	// 最小数量
	MinNum int64 `json:"min_num"`
}
type SpuGetSpuRuleData struct {
	// SPU属性规则
	SpuPropertyRule []SpuPropertyRuleItem `json:"spu_property_rule"`
	// SPU图片规则
	SpuImagesRule *SpuImagesRule `json:"spu_images_rule"`
	// SPU实物图规则（实物图：证明SPU存在的图片，如版权页、包装图）
	SpuActualImagesRule *SpuActualImagesRule `json:"spu_actual_images_rule"`
	// 0-不管控，商家在该类目下发布商品时，不强制要求命中SPU。 - 1-弱管控，商家在该类目下发布商品时，强制要求命中SPU（即关键属性必须一致），但绑定属性不强制要求一致。 - 2-强管控，商家在该类目下发布商品时，强制要求命中SPU，且绑定属性必须一致。
	ControlType int64 `json:"control_type"`
	// - true 表示商家在该类目基于SPU发布商品，发商品的管控策略见管控类型。 - false 表示商家在该类目不需要基于SPU发布商品。不用发布SPU，直接发布商品即可。
	SupportSpuProduct bool `json:"support_spu_product"`
	// true-可以发布SPU，false-不能发布SPU
	SupportCreateSpu bool `json:"support_create_spu"`
}
type ValuesItem struct {
	// 属性值ID，对于输入框类型，填0即可
	ValueId int64 `json:"value_id"`
	// 属性值，属性值的中文
	ValueName string `json:"value_name"`
}
type SpuPropertyRuleItem struct {
	// 属性名
	PropertyName string `json:"property_name"`
	// 属性ID
	PropertyId int64 `json:"property_id"`
	// 是否必填，0:非必填，1:必填
	IsRequired int64 `json:"is_required"`
	// 输入类型 1单选 2多选, 3输入框，4时间戳，5时间区间，6成分属性
	ValueType string `json:"value_type"`
	// 属性类型，0 绑定属性 1关键属性 2销售属性
	PropertyType int64 `json:"property_type"`
	// 是否支持枚举可输入
	DiyType int64 `json:"diy_type"`
	// 最大可选数量，多选时需要用
	MaxSelectNum int64 `json:"max_select_num"`
	// 属性值
	Values []ValuesItem `json:"values"`
}
type SpuImagesRule struct {
	// 是否必填
	IsRequired int64 `json:"is_required"`
	// 最大数量
	MaxNum int64 `json:"max_num"`
	// 最小数量
	MinNum int64 `json:"min_num"`
}
