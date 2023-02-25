package product_qualityDetail_request

import (
	"doudian.com/open/sdk_golang/api/product_qualityDetail/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ProductQualityDetailRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ProductQualityDetailParam 
}
func (c *ProductQualityDetailRequest) GetUrlPath() string{
	return "/product/qualityDetail"
}


func New() *ProductQualityDetailRequest{
	request := &ProductQualityDetailRequest{
		Param: &ProductQualityDetailParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ProductQualityDetailRequest) Execute(accessToken *doudian_sdk.AccessToken) (*product_qualityDetail_response.ProductQualityDetailResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &product_qualityDetail_response.ProductQualityDetailResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ProductQualityDetailRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *ProductQualityDetailRequest) GetParams() *ProductQualityDetailParam{
	return c.Param
}


type ProductQualityDetailParam struct {
	// 商品ID
	ProductId *int64 `json:"product_id,omitempty"`
}
