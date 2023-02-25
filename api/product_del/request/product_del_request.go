package product_del_request

import (
	"doudian.com/open/sdk_golang/api/product_del/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ProductDelRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ProductDelParam 
}
func (c *ProductDelRequest) GetUrlPath() string{
	return "/product/del"
}


func New() *ProductDelRequest{
	request := &ProductDelRequest{
		Param: &ProductDelParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ProductDelRequest) Execute(accessToken *doudian_sdk.AccessToken) (*product_del_response.ProductDelResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &product_del_response.ProductDelResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ProductDelRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *ProductDelRequest) GetParams() *ProductDelParam{
	return c.Param
}


type ProductDelParam struct {
	// 商品ID
	ProductId int64 `json:"product_id,omitempty"`
	// 外部商品ID
	OutProductId int64 `json:"out_product_id,omitempty"`
	// 是否彻底删除
	DeleteForever bool `json:"delete_forever,omitempty"`
}
