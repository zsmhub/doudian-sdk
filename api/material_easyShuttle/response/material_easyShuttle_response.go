package material_easyShuttle_response

import (
	"doudian.com/open/sdk_golang/core"
)

type MaterialEasyShuttleResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *MaterialEasyShuttleData `json:"data"`
}
type MaterialEasyShuttleData struct {
}
