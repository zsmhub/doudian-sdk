package product_detail_request

import (
	"doudian.com/open/sdk_golang/api/product_detail/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ProductDetailRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ProductDetailParam 
}
func (c *ProductDetailRequest) GetUrlPath() string{
	return "/product/detail"
}


func New() *ProductDetailRequest{
	request := &ProductDetailRequest{
		Param: &ProductDetailParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ProductDetailRequest) Execute(accessToken *doudian_sdk.AccessToken) (*product_detail_response.ProductDetailResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &product_detail_response.ProductDetailResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ProductDetailRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *ProductDetailRequest) GetParams() *ProductDetailParam{
	return c.Param
}


type ProductDetailParam struct {
	// 商品ID，抖店系统生成，店铺下唯一；长度19位。
	ProductId string `json:"product_id,omitempty"`
	// 外部商家编码，商家自定义字段
	OutProductId string `json:"out_product_id,omitempty"`
	// true：读取草稿数据；false：读取线上数据；不传默认为false
	ShowDraft string `json:"show_draft,omitempty"`
}
