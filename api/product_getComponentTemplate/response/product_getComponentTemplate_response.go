package product_getComponentTemplate_response

import (
	"doudian.com/open/sdk_golang/core"
)

type ProductGetComponentTemplateResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *ProductGetComponentTemplateData `json:"data"`
}
type Image struct {
	// 图片url
	Url string `json:"url"`
	// 图片宽度
	Width int16 `json:"width"`
	// 图片长度
	Height int16 `json:"height"`
}
type ComponentTemplateInfoListItem struct {
	// 模板ID
	TemplateId int64 `json:"template_id"`
	// 模板类型: size_info(尺码表)
	TemplateType string `json:"template_type"`
	// 组件模板子类型:clothing、undies、shoes、children_clothing
	TemplateSubType string `json:"template_sub_type"`
	// 模板名称
	TemplateName string `json:"template_name"`
	// 模板数据
	ComponentData string `json:"component_data"`
	// 尺码模板图片
	Image *Image `json:"image"`
	// 是否可共享 是-false 否-true
	Shareable bool `json:"shareable"`
	// 类目ID
	CategoryId int64 `json:"category_id"`
	// 创建时间
	CreateTime int64 `json:"create_time"`
	// 更新时间
	UpdateTime int64 `json:"update_time"`
	// 模板配置数据
	ComponentFrontData string `json:"component_front_data"`
}
type ProductGetComponentTemplateData struct {
	// 尺码模板列表
	ComponentTemplateInfoList []ComponentTemplateInfoListItem `json:"component_template_info_list"`
	// 总页数
	TotalNum int64 `json:"total_num"`
}
