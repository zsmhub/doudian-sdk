package address_provinceList_response

import (
	"doudian.com/open/sdk_golang/core"
)

type AddressProvinceListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data []DataItem `json:"data"`
}
type AddressProvinceListData struct {
	Data []DataItem `json:"data"`
}
type DataItem struct {
	ProvinceId int64 `json:"province_id"`
	Province string `json:"province"`
	FatherId int64 `json:"father_id"`
}
