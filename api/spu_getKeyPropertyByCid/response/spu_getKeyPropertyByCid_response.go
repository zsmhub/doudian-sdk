package spu_getKeyPropertyByCid_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SpuGetKeyPropertyByCidResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SpuGetKeyPropertyByCidData `json:"data"`
}
type PropertyInfoItem struct {
	// 关联属性值id，例如手机类目，型号关联属性值id就是品牌的品牌id
	RelValueId int64 `json:"rel_value_id"`
	// 属性id
	PropertyId int64 `json:"property_id"`
	// 属性名
	PropertyName string `json:"property_name"`
	// 属性值id
	ValueId int64 `json:"value_id"`
	// 属性值名称
	ValueName string `json:"value_name"`
}
type BrandInfoItem struct {
	// 关联属性值id,没有为0
	RelValueId int64 `json:"rel_value_id"`
	// 属性id
	PropertyId int64 `json:"property_id"`
	// 属性名
	PropertyName string `json:"property_name"`
	// 属性值id
	ValueId int64 `json:"value_id"`
	// 属性值名
	ValueName string `json:"value_name"`
}
type SpuGetKeyPropertyByCidData struct {
	// 属性信息
	PropertyInfo []PropertyInfoItem `json:"property_info"`
	// 品牌信息（品牌单独处理）
	BrandInfo []BrandInfoItem `json:"brand_info"`
	// 总数
	Total int64 `json:"total"`
}
