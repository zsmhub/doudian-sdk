package spu_getSpuInfoBySpuId_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SpuGetSpuInfoBySpuIdResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SpuGetSpuInfoBySpuIdData `json:"data"`
}
type SpuPropertyValueInfoItem struct {
	// 属性值id
	ValueId int64 `json:"value_id"`
	// 属性值名称
	ValueName string `json:"value_name"`
	// 属性值别名
	ValueAlias string `json:"value_alias"`
}
type SpuPropertyInfosItem struct {
	// 属性id
	PropertyId int64 `json:"property_id"`
	// 属性名
	PropertyName string `json:"property_name"`
	// 属性类型，0 绑定属性 1关键属性 2售卖属性 3 商品属性
	Type int64 `json:"type"`
	// 属性别名
	PropertyAlias string `json:"property_alias"`
	// 属性值信息
	SpuPropertyValueInfo []SpuPropertyValueInfoItem `json:"spu_property_value_info"`
}
type SpuInfo struct {
	// spuName
	SpuName string `json:"spu_name"`
	// 产品编码
	UpcCode string `json:"upc_code"`
	// 类目id
	CategoryId int64 `json:"category_id"`
	// 品牌id
	BrandId int64 `json:"brand_id"`
	// 在线状态 1-生效（默认） 2-失效
	Status int64 `json:"status"`
	// 更新时间
	UpdateTime string `json:"update_time"`
	// 创建时间
	CreateTime string `json:"create_time"`
	// spuId
	SpuId string `json:"spu_id"`
	// 审核状态 -1-删除 0-待审核，1-审核中，2-审核不通过 ，3-审核通过4-撤销
	OpStatus int64 `json:"op_status"`
	// 审核时间
	AuditTime string `json:"audit_time"`
}
type SpuGetSpuInfoBySpuIdData struct {
	// spu属性信息
	SpuPropertyInfos []SpuPropertyInfosItem `json:"spu_property_infos"`
	// spu信息
	SpuInfo *SpuInfo `json:"spu_info"`
}
