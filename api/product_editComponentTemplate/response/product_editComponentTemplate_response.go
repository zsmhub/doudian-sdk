package product_editComponentTemplate_response

import (
	"doudian.com/open/sdk_golang/core"
)

type ProductEditComponentTemplateResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *ProductEditComponentTemplateData `json:"data"`
}
type ProductEditComponentTemplateData struct {
}
