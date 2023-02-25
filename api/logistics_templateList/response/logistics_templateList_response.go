package logistics_templateList_response

import (
	"doudian.com/open/sdk_golang/core"
)

type LogisticsTemplateListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *LogisticsTemplateListData `json:"data"`
}
type TemplateInfosItem struct {
	// 模版id
	TemplateId int64 `json:"template_id"`
	// 模版编码
	TemplateCode string `json:"template_code"`
	// 模版名称
	TemplateName string `json:"template_name"`
	// 模版URL
	TemplateUrl string `json:"template_url"`
	// 版本
	Version int16 `json:"version"`
	// 模版类型； 1-一联单 2-二联单
	TemplateType int16 `json:"template_type"`
	// 预览URL
	PerviewUrl string `json:"perview_url"`
}
type TemplateDataItem struct {
	// 模版信息
	TemplateInfos []TemplateInfosItem `json:"template_infos"`
	// 物流公司
	LogisticsCode string `json:"logistics_code"`
}
type LogisticsTemplateListData struct {
	// 模版数据
	TemplateData []TemplateDataItem `json:"template_data"`
}
