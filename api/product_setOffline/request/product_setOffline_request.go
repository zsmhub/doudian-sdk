package product_setOffline_request

import (
	"doudian.com/open/sdk_golang/api/product_setOffline/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ProductSetOfflineRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ProductSetOfflineParam 
}
func (c *ProductSetOfflineRequest) GetUrlPath() string{
	return "/product/setOffline"
}


func New() *ProductSetOfflineRequest{
	request := &ProductSetOfflineRequest{
		Param: &ProductSetOfflineParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ProductSetOfflineRequest) Execute(accessToken *doudian_sdk.AccessToken) (*product_setOffline_response.ProductSetOfflineResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &product_setOffline_response.ProductSetOfflineResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ProductSetOfflineRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *ProductSetOfflineRequest) GetParams() *ProductSetOfflineParam{
	return c.Param
}


type ProductSetOfflineParam struct {
	// 商品ID
	ProductId *int64 `json:"product_id,omitempty"`
}
