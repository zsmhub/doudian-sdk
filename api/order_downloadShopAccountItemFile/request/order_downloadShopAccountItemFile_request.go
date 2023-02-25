package order_downloadShopAccountItemFile_request

import (
	"doudian.com/open/sdk_golang/api/order_downloadShopAccountItemFile/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderDownloadShopAccountItemFileRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderDownloadShopAccountItemFileParam 
}
func (c *OrderDownloadShopAccountItemFileRequest) GetUrlPath() string{
	return "/order/downloadShopAccountItemFile"
}


func New() *OrderDownloadShopAccountItemFileRequest{
	request := &OrderDownloadShopAccountItemFileRequest{
		Param: &OrderDownloadShopAccountItemFileParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderDownloadShopAccountItemFileRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_downloadShopAccountItemFile_response.OrderDownloadShopAccountItemFileResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_downloadShopAccountItemFile_response.OrderDownloadShopAccountItemFileResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderDownloadShopAccountItemFileRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderDownloadShopAccountItemFileRequest) GetParams() *OrderDownloadShopAccountItemFileParam{
	return c.Param
}


type OrderDownloadShopAccountItemFileParam struct {
	// 下载id
	DownloadId *string `json:"download_id,omitempty"`
}
