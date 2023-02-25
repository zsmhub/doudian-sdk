package afterSale_OpenAfterSaleChannel_request

import (
	"doudian.com/open/sdk_golang/api/afterSale_OpenAfterSaleChannel/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type AfterSaleOpenAfterSaleChannelRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *AfterSaleOpenAfterSaleChannelParam 
}
func (c *AfterSaleOpenAfterSaleChannelRequest) GetUrlPath() string{
	return "/afterSale/OpenAfterSaleChannel"
}


func New() *AfterSaleOpenAfterSaleChannelRequest{
	request := &AfterSaleOpenAfterSaleChannelRequest{
		Param: &AfterSaleOpenAfterSaleChannelParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *AfterSaleOpenAfterSaleChannelRequest) Execute(accessToken *doudian_sdk.AccessToken) (*afterSale_OpenAfterSaleChannel_response.AfterSaleOpenAfterSaleChannelResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &afterSale_OpenAfterSaleChannel_response.AfterSaleOpenAfterSaleChannelResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *AfterSaleOpenAfterSaleChannelRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *AfterSaleOpenAfterSaleChannelRequest) GetParams() *AfterSaleOpenAfterSaleChannelParam{
	return c.Param
}


type AfterSaleOpenAfterSaleChannelParam struct {
	// 订单ID
	OrderId *int64 `json:"order_id,omitempty"`
}
