package product_getCatePropertyV2_request

import (
	"doudian.com/open/sdk_golang/api/product_getCatePropertyV2/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ProductGetCatePropertyV2Request struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ProductGetCatePropertyV2Param 
}
func (c *ProductGetCatePropertyV2Request) GetUrlPath() string{
	return "/product/getCatePropertyV2"
}


func New() *ProductGetCatePropertyV2Request{
	request := &ProductGetCatePropertyV2Request{
		Param: &ProductGetCatePropertyV2Param{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ProductGetCatePropertyV2Request) Execute(accessToken *doudian_sdk.AccessToken) (*product_getCatePropertyV2_response.ProductGetCatePropertyV2Response, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &product_getCatePropertyV2_response.ProductGetCatePropertyV2Response{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ProductGetCatePropertyV2Request) GetParamObject() interface{}{
	return c.Param
}


func (c *ProductGetCatePropertyV2Request) GetParams() *ProductGetCatePropertyV2Param{
	return c.Param
}


type ProductGetCatePropertyV2Param struct {
	// 叶子类目id 1、传category_leaf_id ，则不需要传first_cid、second_cid、third_cid这三个字段 2、如果没传category_leaf_id，走之前的逻辑，需要传first_cid、second_cid、third_cid这三个字段
	CategoryLeafId *int64 `json:"category_leaf_id,omitempty"`
}
