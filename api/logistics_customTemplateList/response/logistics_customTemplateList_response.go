package logistics_customTemplateList_response

import (
	"doudian.com/open/sdk_golang/core"
)

type LogisticsCustomTemplateListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *LogisticsCustomTemplateListData `json:"data"`
}
type CustomTemplateInfosItem struct {
	// 自定义模板code
	CustomTemplateCode string `json:"custom_template_code"`
	// 自定义区模板url(URL资源的内容为xml格式的报文)
	CustomTemplateUrl string `json:"custom_template_url"`
	// 自定义区模板名称
	CustomTemplateName string `json:"custom_template_name"`
	// 父模板code(查询标准模板API中返回的template_code)
	ParentTemplateCode string `json:"parent_template_code"`
}
type CustomTemplateDataItem struct {
	// 物流商编码
	LogisticsCode string `json:"logistics_code"`
	// 用户使用的模板数据
	CustomTemplateInfos []CustomTemplateInfosItem `json:"custom_template_infos"`
}
type LogisticsCustomTemplateListData struct {
	// 自定义模板的数据列表
	CustomTemplateData []CustomTemplateDataItem `json:"custom_template_data"`
}
