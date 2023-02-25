package address_create_response

import (
	"doudian.com/open/sdk_golang/core"
)

type AddressCreateResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *AddressCreateData `json:"data"`
}
type AddressCreateData struct {
	// 新建地址ID
	AddressId int64 `json:"address_id"`
}
