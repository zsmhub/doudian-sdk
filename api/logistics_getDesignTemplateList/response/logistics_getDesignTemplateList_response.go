package logistics_getDesignTemplateList_response

import (
	"doudian.com/open/sdk_golang/core"
)

type LogisticsGetDesignTemplateListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *LogisticsGetDesignTemplateListData `json:"data"`
}
type DesignTemplateDataItem struct {
	// 自定义模板code
	DesignTemplateCode string `json:"design_template_code"`
	// 自定义模板名称
	DesignTemplateName string `json:"design_template_name"`
	// 自定义模板url
	DesignTemplateUrl string `json:"design_template_url"`
	// 打印项中字段列表
	DesignTemplateKeyList []string `json:"design_template_key_list"`
}
type LogisticsGetDesignTemplateListData struct {
	// 已发布的自定义模板列表
	DesignTemplateData []DesignTemplateDataItem `json:"design_template_data"`
}
