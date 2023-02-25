package product_createComponentTemplateV2_request

import (
	"doudian.com/open/sdk_golang/api/product_createComponentTemplateV2/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ProductCreateComponentTemplateV2Request struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ProductCreateComponentTemplateV2Param 
}
func (c *ProductCreateComponentTemplateV2Request) GetUrlPath() string{
	return "/product/createComponentTemplateV2"
}


func New() *ProductCreateComponentTemplateV2Request{
	request := &ProductCreateComponentTemplateV2Request{
		Param: &ProductCreateComponentTemplateV2Param{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ProductCreateComponentTemplateV2Request) Execute(accessToken *doudian_sdk.AccessToken) (*product_createComponentTemplateV2_response.ProductCreateComponentTemplateV2Response, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &product_createComponentTemplateV2_response.ProductCreateComponentTemplateV2Response{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ProductCreateComponentTemplateV2Request) GetParamObject() interface{}{
	return c.Param
}


func (c *ProductCreateComponentTemplateV2Request) GetParams() *ProductCreateComponentTemplateV2Param{
	return c.Param
}


type ProductCreateComponentTemplateV2Param struct {
	// 模板类型：尺码模板
	TemplateType *string `json:"template_type,omitempty"`
	// 模板子类型: clothing(服装)、undies(内衣)、shoes(鞋靴类)、children_clothing(童装)
	TemplateSubType *string `json:"template_sub_type,omitempty"`
	// 模板名称
	TemplateName *string `json:"template_name,omitempty"`
	// 商品组件数据 json，表格行列顺序以selectedSize和selectedSpecs的顺序为准
	ComponentFrontData *string `json:"component_front_data,omitempty"`
	// 是否设置为公有模板(多个商品可共用)。true-是，false-不是；不传默认fasle
	Shareable bool `json:"shareable,omitempty"`
	// 类目id，用来确定模板类型
	CategoryId int64 `json:"category_id,omitempty"`
}
