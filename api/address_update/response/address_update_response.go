package address_update_response

import (
	"doudian.com/open/sdk_golang/core"
)

type AddressUpdateResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *AddressUpdateData `json:"data"`
}
type AddressUpdateData struct {
}
