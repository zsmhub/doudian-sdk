package afterSale_returnGoodsToWareHouseSuccess_request

import (
	"doudian.com/open/sdk_golang/api/afterSale_returnGoodsToWareHouseSuccess/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type AfterSaleReturnGoodsToWareHouseSuccessRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *AfterSaleReturnGoodsToWareHouseSuccessParam 
}
func (c *AfterSaleReturnGoodsToWareHouseSuccessRequest) GetUrlPath() string{
	return "/afterSale/returnGoodsToWareHouseSuccess"
}


func New() *AfterSaleReturnGoodsToWareHouseSuccessRequest{
	request := &AfterSaleReturnGoodsToWareHouseSuccessRequest{
		Param: &AfterSaleReturnGoodsToWareHouseSuccessParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *AfterSaleReturnGoodsToWareHouseSuccessRequest) Execute(accessToken *doudian_sdk.AccessToken) (*afterSale_returnGoodsToWareHouseSuccess_response.AfterSaleReturnGoodsToWareHouseSuccessResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &afterSale_returnGoodsToWareHouseSuccess_response.AfterSaleReturnGoodsToWareHouseSuccessResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *AfterSaleReturnGoodsToWareHouseSuccessRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *AfterSaleReturnGoodsToWareHouseSuccessRequest) GetParams() *AfterSaleReturnGoodsToWareHouseSuccessParam{
	return c.Param
}


type AfterSaleReturnGoodsToWareHouseSuccessParam struct {
	// 售后单Id
	AftersaleId *string `json:"aftersale_id,omitempty"`
	// 商家确认退货入仓时间，Unix时间戳，时间为秒
	OpTime *int64 `json:"op_time,omitempty"`
	// 用户退货物流单号
	TrackingNo string `json:"tracking_no,omitempty"`
	// 物流公司代号
	LogisticsCompanyCode string `json:"logistics_company_code,omitempty"`
}
