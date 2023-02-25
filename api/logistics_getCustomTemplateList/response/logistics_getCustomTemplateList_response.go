package logistics_getCustomTemplateList_response

import (
	"doudian.com/open/sdk_golang/core"
)

type LogisticsGetCustomTemplateListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *LogisticsGetCustomTemplateListData `json:"data"`
}
type CustomTemplateInfosItem struct {
	// 自定义区模板code
	CustomTemplateCode string `json:"custom_template_code"`
	// 自定义区模板名称
	CustomTemplateName string `json:"custom_template_name"`
	// 父模板code(查询标准模板API中返回的template_code)
	ParentTemplateCode string `json:"parent_template_code"`
	// 自定义区模板url(URL资源的内容为xml格式的报文)
	CustomTemplateUrl string `json:"custom_template_url"`
	// customTemplateKeyList（打印项中字段列表）
	CustomTemplateKeyList []string `json:"custom_template_key_list"`
	// 自定义区模板id
	CustomTemplateId int64 `json:"custom_template_id"`
	// 父模板id
	ParentTemplateId int64 `json:"parent_template_id"`
}
type CustomTemplateDataItem struct {
	// 物流服务商编码
	LogisticsCode string `json:"logistics_code"`
	// 用户使用的模板数据
	CustomTemplateInfos []CustomTemplateInfosItem `json:"custom_template_infos"`
}
type LogisticsGetCustomTemplateListData struct {
	// 商家所有快递自定义模板的数据列表
	CustomTemplateData []CustomTemplateDataItem `json:"custom_template_data"`
}
