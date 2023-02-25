package shop_brandList_request

import (
	"doudian.com/open/sdk_golang/api/shop_brandList/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ShopBrandListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ShopBrandListParam 
}
func (c *ShopBrandListRequest) GetUrlPath() string{
	return "/shop/brandList"
}


func New() *ShopBrandListRequest{
	request := &ShopBrandListRequest{
		Param: &ShopBrandListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ShopBrandListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*shop_brandList_response.ShopBrandListResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &shop_brandList_response.ShopBrandListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ShopBrandListRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *ShopBrandListRequest) GetParams() *ShopBrandListParam{
	return c.Param
}


type ShopBrandListParam struct {
}
