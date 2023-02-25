package rights_info_response

import (
	"doudian.com/open/sdk_golang/core"
)

type RightsInfoResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *RightsInfoData `json:"data"`
}
type RightsInfoData struct {
	// 权益到期时间
	ExpireTime string `json:"expire_time"`
	// 0-试用版，1-正式版
	RightsType int32 `json:"rights_type"`
	// 规格类型，0-版本，其他待定
	SpecType int32 `json:"spec_type"`
	// 规格名称
	SpecVal string `json:"spec_val"`
}
