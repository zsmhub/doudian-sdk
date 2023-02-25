package afterSale_timeExtend_request

import (
	"doudian.com/open/sdk_golang/api/afterSale_timeExtend/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type AfterSaleTimeExtendRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *AfterSaleTimeExtendParam 
}
func (c *AfterSaleTimeExtendRequest) GetUrlPath() string{
	return "/afterSale/timeExtend"
}


func New() *AfterSaleTimeExtendRequest{
	request := &AfterSaleTimeExtendRequest{
		Param: &AfterSaleTimeExtendParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *AfterSaleTimeExtendRequest) Execute(accessToken *doudian_sdk.AccessToken) (*afterSale_timeExtend_response.AfterSaleTimeExtendResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &afterSale_timeExtend_response.AfterSaleTimeExtendResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *AfterSaleTimeExtendRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *AfterSaleTimeExtendRequest) GetParams() *AfterSaleTimeExtendParam{
	return c.Param
}


type AfterSaleTimeExtendParam struct {
	// 售后单号
	AftersaleId *int64 `json:"aftersale_id,omitempty"`
}
