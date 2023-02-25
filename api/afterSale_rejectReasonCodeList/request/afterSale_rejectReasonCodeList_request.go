package afterSale_rejectReasonCodeList_request

import (
	"doudian.com/open/sdk_golang/api/afterSale_rejectReasonCodeList/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type AfterSaleRejectReasonCodeListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *AfterSaleRejectReasonCodeListParam 
}
func (c *AfterSaleRejectReasonCodeListRequest) GetUrlPath() string{
	return "/afterSale/rejectReasonCodeList"
}


func New() *AfterSaleRejectReasonCodeListRequest{
	request := &AfterSaleRejectReasonCodeListRequest{
		Param: &AfterSaleRejectReasonCodeListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *AfterSaleRejectReasonCodeListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*afterSale_rejectReasonCodeList_response.AfterSaleRejectReasonCodeListResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &afterSale_rejectReasonCodeList_response.AfterSaleRejectReasonCodeListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *AfterSaleRejectReasonCodeListRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *AfterSaleRejectReasonCodeListRequest) GetParams() *AfterSaleRejectReasonCodeListParam{
	return c.Param
}


type AfterSaleRejectReasonCodeListParam struct {
	// 传入售后单id时，返回该笔售后单可选择的拒绝原因列表；仅支持已发货的售后查询，备货中未发货订单返回为空。
	AftersaleId int64 `json:"aftersale_id,omitempty"`
}
