package afterSale_fillLogistics_request

import (
	"doudian.com/open/sdk_golang/api/afterSale_fillLogistics/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type AfterSaleFillLogisticsRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *AfterSaleFillLogisticsParam 
}
func (c *AfterSaleFillLogisticsRequest) GetUrlPath() string{
	return "/afterSale/fillLogistics"
}


func New() *AfterSaleFillLogisticsRequest{
	request := &AfterSaleFillLogisticsRequest{
		Param: &AfterSaleFillLogisticsParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *AfterSaleFillLogisticsRequest) Execute(accessToken *doudian_sdk.AccessToken) (*afterSale_fillLogistics_response.AfterSaleFillLogisticsResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &afterSale_fillLogistics_response.AfterSaleFillLogisticsResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *AfterSaleFillLogisticsRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *AfterSaleFillLogisticsRequest) GetParams() *AfterSaleFillLogisticsParam{
	return c.Param
}


type AfterSaleFillLogisticsParam struct {
	// 售后单ID
	AftersaleId *int64 `json:"aftersale_id,omitempty"`
	// 发货类型；适用场景： send_type=1：用于补寄商家发货 send_type=3：超市预约上门取货；退货退款和换货场景下商家帮买家填写退货物流信息； send_type=4：维修场景下商家帮买家填写退货物流信息；
	SendType *int32 `json:"send_type,omitempty"`
	// 物流公司编号pick_up_type 2:线下取货;3:用户退回，无需物流公司，可传 -
	CompanyCode *string `json:"company_code,omitempty"`
	// 物流单号pick_up_type 2:线下取货;3:用户退回，无需快递单号，可传 -
	TrackingNo *string `json:"tracking_no,omitempty"`
	// 预约上门取货时间戳，单位：秒（目前抖超小时达店铺使用）
	BookTimeBegin int64 `json:"book_time_begin,omitempty"`
	// 预约上门取货时间戳，单位：秒（目前抖超小时达店铺使用）
	BookTimeEnd int64 `json:"book_time_end,omitempty"`
	// 门店ID
	StoreId int64 `json:"store_id,omitempty"`
	// 1:自行配送;2:线下取货;3:用户退回 ，不传默认自行配送；适用于send_type=3超市预约上门取货场景
	PickUpType int32 `json:"pick_up_type,omitempty"`
}
