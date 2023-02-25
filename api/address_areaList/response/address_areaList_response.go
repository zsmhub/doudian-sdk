package address_areaList_response

import (
	"doudian.com/open/sdk_golang/core"
)

type AddressAreaListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data []DataItem `json:"data"`
}
type DataItem struct {
	AreaId int64 `json:"area_id"`
	Area string `json:"area"`
	FatherId int64 `json:"father_id"`
}
type AddressAreaListData struct {
	Data []DataItem `json:"data"`
}
