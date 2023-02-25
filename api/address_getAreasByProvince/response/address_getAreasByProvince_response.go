package address_getAreasByProvince_response

import (
	"doudian.com/open/sdk_golang/core"
)

type AddressGetAreasByProvinceResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data []DataItem `json:"data"`
}
type SubDistrictsItem struct {
	// 镇或者街道ID
	Code int64 `json:"code"`
	// 父ID
	FatherCode int64 `json:"father_code"`
	// 地址等级
	Level string `json:"level"`
	// 镇或者街道名称
	Name string `json:"name"`
}
type SubDistrictsItem_4 struct {
	// 区县ID
	Code int64 `json:"code"`
	// 父ID
	FatherCode int64 `json:"father_code"`
	// 区县名称
	Name string `json:"name"`
	// 地址等级
	Level string `json:"level"`
	// 响应结果
	SubDistricts []SubDistrictsItem `json:"sub_districts"`
}
type SubDistrictsItem_3 struct {
	// 市ID
	Code int64 `json:"code"`
	// 父ID
	FatherCode int64 `json:"father_code"`
	// 市名
	Name string `json:"name"`
	// 地址等级
	Level string `json:"level"`
	// 响应结果
	SubDistricts []SubDistrictsItem_4 `json:"sub_districts"`
}
type DataItem struct {
	// 省ID
	Code int64 `json:"code"`
	// 父ID
	FatherCode int64 `json:"father_code"`
	// 省名称
	Name string `json:"name"`
	// 地址等级
	Level string `json:"level"`
	// 响应结果
	SubDistricts []SubDistrictsItem_3 `json:"sub_districts"`
}
type AddressGetAreasByProvinceData struct {
	// 响应结果
	Data []DataItem `json:"data"`
}
