package product_CategoryDimList_request

import (
	"doudian.com/open/sdk_golang/api/product_CategoryDimList/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ProductCategoryDimListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ProductCategoryDimListParam 
}
func (c *ProductCategoryDimListRequest) GetUrlPath() string{
	return "/product/CategoryDimList"
}


func New() *ProductCategoryDimListRequest{
	request := &ProductCategoryDimListRequest{
		Param: &ProductCategoryDimListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ProductCategoryDimListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*product_CategoryDimList_response.ProductCategoryDimListResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &product_CategoryDimList_response.ProductCategoryDimListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ProductCategoryDimListRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *ProductCategoryDimListRequest) GetParams() *ProductCategoryDimListParam{
	return c.Param
}


type ProductCategoryDimListParam struct {
	// 类目层级
	Level int32 `json:"level,omitempty"`
	// 类目名, 模糊搜索
	Name string `json:"name,omitempty"`
	// 上级类目ID，一级类目上级类目ID为0，可选
	ParentId int64 `json:"parent_id,omitempty"`
}
