package address_getProvince_response

import (
	"doudian.com/open/sdk_golang/core"
)

type AddressGetProvinceResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data []DataItem `json:"data"`
}
type DataItem struct {
	// 省份id
	ProvinceId int64 `json:"province_id"`
	// 省份
	Province string `json:"province"`
}
type AddressGetProvinceData struct {
	// data
	Data []DataItem `json:"data"`
}
