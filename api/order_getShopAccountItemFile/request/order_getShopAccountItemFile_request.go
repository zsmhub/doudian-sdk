package order_getShopAccountItemFile_request

import (
	"doudian.com/open/sdk_golang/api/order_getShopAccountItemFile/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderGetShopAccountItemFileRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderGetShopAccountItemFileParam 
}
func (c *OrderGetShopAccountItemFileRequest) GetUrlPath() string{
	return "/order/getShopAccountItemFile"
}


func New() *OrderGetShopAccountItemFileRequest{
	request := &OrderGetShopAccountItemFileRequest{
		Param: &OrderGetShopAccountItemFileParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderGetShopAccountItemFileRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_getShopAccountItemFile_response.OrderGetShopAccountItemFileResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_getShopAccountItemFile_response.OrderGetShopAccountItemFileResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderGetShopAccountItemFileRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderGetShopAccountItemFileRequest) GetParams() *OrderGetShopAccountItemFileParam{
	return c.Param
}


type OrderGetShopAccountItemFileParam struct {
	// 开始账单日期(闭区间)
	StartDate string `json:"start_date,omitempty"`
	// 结束账单日期(闭区间)
	EndDate string `json:"end_date,omitempty"`
}
