package freightTemplate_create_response

import (
	"doudian.com/open/sdk_golang/core"
)

type FreightTemplateCreateResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *FreightTemplateCreateData `json:"data"`
}
type FreightTemplateCreateData struct {
	// 创建的模板的id
	TemplateId int64 `json:"template_id"`
}
