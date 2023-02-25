package order_downloadToShop_request

import (
	"doudian.com/open/sdk_golang/api/order_downloadToShop/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderDownloadToShopRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderDownloadToShopParam 
}
func (c *OrderDownloadToShopRequest) GetUrlPath() string{
	return "/order/downloadToShop"
}


func New() *OrderDownloadToShopRequest{
	request := &OrderDownloadToShopRequest{
		Param: &OrderDownloadToShopParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderDownloadToShopRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_downloadToShop_response.OrderDownloadToShopResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_downloadToShop_response.OrderDownloadToShopResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderDownloadToShopRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderDownloadToShopRequest) GetParams() *OrderDownloadToShopParam{
	return c.Param
}


type OrderDownloadToShopParam struct {
	// 生成的download_id
	DownloadId *string `json:"download_id,omitempty"`
}
