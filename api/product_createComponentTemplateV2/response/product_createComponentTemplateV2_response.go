package product_createComponentTemplateV2_response

import (
	"doudian.com/open/sdk_golang/core"
)

type ProductCreateComponentTemplateV2Response struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *ProductCreateComponentTemplateV2Data `json:"data"`
}
type ProductCreateComponentTemplateV2Data struct {
	// 模板ID
	TemplateId int64 `json:"template_id"`
}
