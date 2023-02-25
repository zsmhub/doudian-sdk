package product_getComponentTemplate_request

import (
	"doudian.com/open/sdk_golang/api/product_getComponentTemplate/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ProductGetComponentTemplateRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ProductGetComponentTemplateParam 
}
func (c *ProductGetComponentTemplateRequest) GetUrlPath() string{
	return "/product/getComponentTemplate"
}


func New() *ProductGetComponentTemplateRequest{
	request := &ProductGetComponentTemplateRequest{
		Param: &ProductGetComponentTemplateParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ProductGetComponentTemplateRequest) Execute(accessToken *doudian_sdk.AccessToken) (*product_getComponentTemplate_response.ProductGetComponentTemplateResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &product_getComponentTemplate_response.ProductGetComponentTemplateResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ProductGetComponentTemplateRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *ProductGetComponentTemplateRequest) GetParams() *ProductGetComponentTemplateParam{
	return c.Param
}


type ProductGetComponentTemplateParam struct {
	// 模板类型: size_info(尺码表)
	TemplateType string `json:"template_type,omitempty"`
	// 组件模板子类型；clothing -饰；undies-内衣；shoes 鞋靴；children_clothing-童装
	TemplateSubType string `json:"template_sub_type,omitempty"`
	// 模板ID
	TemplateId int64 `json:"template_id,omitempty"`
	// 是否设置为公有模板(多个商品可共用)。true-是，false-不是；不传默认true
	Shareable bool `json:"shareable,omitempty"`
	// 页码，从0开始，最大支持100
	PageNum int64 `json:"page_num,omitempty"`
	// 每页条数，最大支持20；page_size和template_id二选一入参
	PageSize int64 `json:"page_size,omitempty"`
}
