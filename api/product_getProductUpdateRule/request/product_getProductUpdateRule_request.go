package product_getProductUpdateRule_request

import (
	"doudian.com/open/sdk_golang/api/product_getProductUpdateRule/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ProductGetProductUpdateRuleRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ProductGetProductUpdateRuleParam 
}
func (c *ProductGetProductUpdateRuleRequest) GetUrlPath() string{
	return "/product/getProductUpdateRule"
}


func New() *ProductGetProductUpdateRuleRequest{
	request := &ProductGetProductUpdateRuleRequest{
		Param: &ProductGetProductUpdateRuleParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ProductGetProductUpdateRuleRequest) Execute(accessToken *doudian_sdk.AccessToken) (*product_getProductUpdateRule_response.ProductGetProductUpdateRuleResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &product_getProductUpdateRule_response.ProductGetProductUpdateRuleResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ProductGetProductUpdateRuleRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *ProductGetProductUpdateRuleRequest) GetParams() *ProductGetProductUpdateRuleParam{
	return c.Param
}


type ProductGetProductUpdateRuleParam struct {
	// 类目id
	CategoryId *int64 `json:"category_id,omitempty"`
	// 闪购定制参数，普通发品忽略
	Senses []int32 `json:"senses,omitempty"`
	// 品牌id
	StandardBrandId int64 `json:"standard_brand_id,omitempty"`
	// spu_id
	SpuId int64 `json:"spu_id,omitempty"`
}
