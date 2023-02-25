package spu_getSpu_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SpuGetSpuResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SpuGetSpuData `json:"data"`
}
type PropertyValuesItem struct {
	// 属性值ID
	ValueId int64 `json:"value_id"`
	// 属性值
	ValueName string `json:"value_name"`
}
type PropertyInfosItem struct {
	// 属性ID
	PropertyId int64 `json:"property_id"`
	// 属性值
	PropertyValues []PropertyValuesItem `json:"property_values"`
	// 属性名
	PropertyName string `json:"property_name"`
	// 属性类型，0 绑定属性 1关键属性 2销售属性
	PropertyType int64 `json:"property_type"`
}
type SalePropertiesItem struct {
	// 属性ID
	PropertyId int64 `json:"property_id"`
	// 属性名
	PropertyName string `json:"property_name"`
	// 属性值ID
	ValueId int64 `json:"value_id"`
	// 属性值
	ValueName string `json:"value_name"`
}
type CspusItem struct {
	// CSPU ID
	CspuId int64 `json:"cspu_id"`
	// SPU ID
	SpuId int64 `json:"spu_id"`
	// 销售属性
	SaleProperties []SalePropertiesItem `json:"sale_properties"`
}
type SpuGetSpuData struct {
	// SPU ID
	SpuId int64 `json:"spu_id"`
	// SPU名
	SpuName string `json:"spu_name"`
	// 类目ID
	CategoryLeafId int64 `json:"category_leaf_id"`
	// 属性信息
	PropertyInfos []PropertyInfosItem `json:"property_infos"`
	// CSPU信息
	Cspus []CspusItem `json:"cspus"`
	// 1:上线，2:下线，3:审核中，4:审核不通过
	Status int64 `json:"status"`
}
