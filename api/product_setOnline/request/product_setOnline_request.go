package product_setOnline_request

import (
	"doudian.com/open/sdk_golang/api/product_setOnline/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ProductSetOnlineRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ProductSetOnlineParam 
}
func (c *ProductSetOnlineRequest) GetUrlPath() string{
	return "/product/setOnline"
}


func New() *ProductSetOnlineRequest{
	request := &ProductSetOnlineRequest{
		Param: &ProductSetOnlineParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ProductSetOnlineRequest) Execute(accessToken *doudian_sdk.AccessToken) (*product_setOnline_response.ProductSetOnlineResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &product_setOnline_response.ProductSetOnlineResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ProductSetOnlineRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *ProductSetOnlineRequest) GetParams() *ProductSetOnlineParam{
	return c.Param
}


type ProductSetOnlineParam struct {
	// 商品id
	ProductId *int64 `json:"product_id,omitempty"`
}
