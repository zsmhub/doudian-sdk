package afterSale_CancelSendGoodsSuccess_request

import (
	"doudian.com/open/sdk_golang/api/afterSale_CancelSendGoodsSuccess/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type AfterSaleCancelSendGoodsSuccessRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *AfterSaleCancelSendGoodsSuccessParam 
}
func (c *AfterSaleCancelSendGoodsSuccessRequest) GetUrlPath() string{
	return "/afterSale/CancelSendGoodsSuccess"
}


func New() *AfterSaleCancelSendGoodsSuccessRequest{
	request := &AfterSaleCancelSendGoodsSuccessRequest{
		Param: &AfterSaleCancelSendGoodsSuccessParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *AfterSaleCancelSendGoodsSuccessRequest) Execute(accessToken *doudian_sdk.AccessToken) (*afterSale_CancelSendGoodsSuccess_response.AfterSaleCancelSendGoodsSuccessResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &afterSale_CancelSendGoodsSuccess_response.AfterSaleCancelSendGoodsSuccessResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *AfterSaleCancelSendGoodsSuccessRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *AfterSaleCancelSendGoodsSuccessRequest) GetParams() *AfterSaleCancelSendGoodsSuccessParam{
	return c.Param
}


type AfterSaleCancelSendGoodsSuccessParam struct {
	// 售后单ID
	AftersaleId *string `json:"aftersale_id,omitempty"`
	// unix时间戳，单位为秒
	OpTime *int64 `json:"op_time,omitempty"`
}
