package order_ordeReportList_request

import (
	"doudian.com/open/sdk_golang/api/order_ordeReportList/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderOrdeReportListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderOrdeReportListParam 
}
func (c *OrderOrdeReportListRequest) GetUrlPath() string{
	return "/order/ordeReportList"
}


func New() *OrderOrdeReportListRequest{
	request := &OrderOrdeReportListRequest{
		Param: &OrderOrdeReportListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderOrdeReportListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_ordeReportList_response.OrderOrdeReportListResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_ordeReportList_response.OrderOrdeReportListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderOrdeReportListRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderOrdeReportListRequest) GetParams() *OrderOrdeReportListParam{
	return c.Param
}


type AddRealMobileWhitesItem struct {
	// 订单号，订单号和售后单号只需选择一个传入即可
	OrderId string `json:"order_id,omitempty"`
	// 售后单号，订单号和售后单号只需选择一个传入即可
	AfterSaleId string `json:"after_sale_id,omitempty"`
	// 报备原因编码：1-无法发货，2-无法处理售后，3-无法处理客诉，4-无法处理物流包裹，5-其他原因
	ReasonNo *int64 `json:"reason_no,omitempty"`
	// 报备备注
	Remark string `json:"remark,omitempty"`
}
type OrderOrdeReportListParam struct {
	// 报备请求
	AddRealMobileWhites *[]AddRealMobileWhitesItem `json:"add_real_mobile_whites,omitempty"`
}
