package open_openId_replace_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OpenOpenIdReplaceResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OpenOpenIdReplaceData `json:"data"`
}
type OpenOpenIdReplaceData struct {
	// 一个map, key为老openId, val为新openId
	OpenIdOldToNewMap map[string]string `json:"open_id_old_to_new_map"`
}
