package afterSale_openOutAfterSale_request

import (
	"doudian.com/open/sdk_golang/api/afterSale_openOutAfterSale/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type AfterSaleOpenOutAfterSaleRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *AfterSaleOpenOutAfterSaleParam 
}
func (c *AfterSaleOpenOutAfterSaleRequest) GetUrlPath() string{
	return "/afterSale/openOutAfterSale"
}


func New() *AfterSaleOpenOutAfterSaleRequest{
	request := &AfterSaleOpenOutAfterSaleRequest{
		Param: &AfterSaleOpenOutAfterSaleParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *AfterSaleOpenOutAfterSaleRequest) Execute(accessToken *doudian_sdk.AccessToken) (*afterSale_openOutAfterSale_response.AfterSaleOpenOutAfterSaleResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &afterSale_openOutAfterSale_response.AfterSaleOpenOutAfterSaleResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *AfterSaleOpenOutAfterSaleRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *AfterSaleOpenOutAfterSaleRequest) GetParams() *AfterSaleOpenOutAfterSaleParam{
	return c.Param
}


type AfterSaleOpenOutAfterSaleParam struct {
	// 要打开超售后入口的商品单ID
	OrderId *int64 `json:"order_id,omitempty"`
}
