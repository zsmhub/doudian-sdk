package product_datchDelComponentTemplate_request

import (
	"doudian.com/open/sdk_golang/api/product_datchDelComponentTemplate/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ProductDatchDelComponentTemplateRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ProductDatchDelComponentTemplateParam 
}
func (c *ProductDatchDelComponentTemplateRequest) GetUrlPath() string{
	return "/product/datchDelComponentTemplate"
}


func New() *ProductDatchDelComponentTemplateRequest{
	request := &ProductDatchDelComponentTemplateRequest{
		Param: &ProductDatchDelComponentTemplateParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ProductDatchDelComponentTemplateRequest) Execute(accessToken *doudian_sdk.AccessToken) (*product_datchDelComponentTemplate_response.ProductDatchDelComponentTemplateResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &product_datchDelComponentTemplate_response.ProductDatchDelComponentTemplateResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ProductDatchDelComponentTemplateRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *ProductDatchDelComponentTemplateRequest) GetParams() *ProductDatchDelComponentTemplateParam{
	return c.Param
}


type ProductDatchDelComponentTemplateParam struct {
	// 模板ID
	TemplateId *[]int64 `json:"template_id,omitempty"`
}
