package material_editFolder_response

import (
	"doudian.com/open/sdk_golang/core"
)

type MaterialEditFolderResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *MaterialEditFolderData `json:"data"`
}
type MaterialEditFolderData struct {
}
