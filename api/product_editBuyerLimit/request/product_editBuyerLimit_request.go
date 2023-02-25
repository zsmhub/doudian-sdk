package product_editBuyerLimit_request

import (
	"doudian.com/open/sdk_golang/api/product_editBuyerLimit/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ProductEditBuyerLimitRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ProductEditBuyerLimitParam 
}
func (c *ProductEditBuyerLimitRequest) GetUrlPath() string{
	return "/product/editBuyerLimit"
}


func New() *ProductEditBuyerLimitRequest{
	request := &ProductEditBuyerLimitRequest{
		Param: &ProductEditBuyerLimitParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ProductEditBuyerLimitRequest) Execute(accessToken *doudian_sdk.AccessToken) (*product_editBuyerLimit_response.ProductEditBuyerLimitResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &product_editBuyerLimit_response.ProductEditBuyerLimitResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ProductEditBuyerLimitRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *ProductEditBuyerLimitRequest) GetParams() *ProductEditBuyerLimitParam{
	return c.Param
}


type ProductEditBuyerLimitParam struct {
	// 商品ID
	ProductId *int64 `json:"product_id,omitempty"`
	// 每个用户每次下单限购件数
	MaximumPerOrder int64 `json:"maximum_per_order,omitempty"`
	// 每个用户累计限购件数
	LimitPerBuyer int64 `json:"limit_per_buyer,omitempty"`
	// 每个用户每次下单至少购买的件数 
	MinimumPerOrder int64 `json:"minimum_per_order,omitempty"`
}
