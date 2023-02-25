package product_editComponentTemplate_request

import (
	"doudian.com/open/sdk_golang/api/product_editComponentTemplate/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ProductEditComponentTemplateRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ProductEditComponentTemplateParam 
}
func (c *ProductEditComponentTemplateRequest) GetUrlPath() string{
	return "/product/editComponentTemplate"
}


func New() *ProductEditComponentTemplateRequest{
	request := &ProductEditComponentTemplateRequest{
		Param: &ProductEditComponentTemplateParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ProductEditComponentTemplateRequest) Execute(accessToken *doudian_sdk.AccessToken) (*product_editComponentTemplate_response.ProductEditComponentTemplateResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &product_editComponentTemplate_response.ProductEditComponentTemplateResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ProductEditComponentTemplateRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *ProductEditComponentTemplateRequest) GetParams() *ProductEditComponentTemplateParam{
	return c.Param
}


type ProductEditComponentTemplateParam struct {
	// 模板ID
	TemplateId *int64 `json:"template_id,omitempty"`
	// 模板名称
	TemplateName string `json:"template_name,omitempty"`
	// 模板数据json
	ComponentData string `json:"component_data,omitempty"`
	// 是否为公有模板(多个商品可共用)
	Shareable bool `json:"shareable,omitempty"`
}
