package material_editMaterial_response

import (
	"doudian.com/open/sdk_golang/core"
)

type MaterialEditMaterialResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *MaterialEditMaterialData `json:"data"`
}
type MaterialEditMaterialData struct {
}
