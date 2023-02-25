package afterSale_List_request

import (
	"doudian.com/open/sdk_golang/api/afterSale_List/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type AfterSaleListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *AfterSaleListParam 
}
func (c *AfterSaleListRequest) GetUrlPath() string{
	return "/afterSale/List"
}


func New() *AfterSaleListRequest{
	request := &AfterSaleListRequest{
		Param: &AfterSaleListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *AfterSaleListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*afterSale_List_response.AfterSaleListResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &afterSale_List_response.AfterSaleListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *AfterSaleListRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *AfterSaleListRequest) GetParams() *AfterSaleListParam{
	return c.Param
}


type AfterSaleListParam struct {
	// 父订单号
	OrderId string `json:"order_id,omitempty"`
	// 售后类型；0-退货退款；1-已发货仅退款；2-未发货仅退款；3-换货；6-价保；7-补寄；8-维修
	AftersaleType int64 `json:"aftersale_type,omitempty"`
	// 已废弃，推荐使用standard_aftersale_status字段。售后状态，枚举为6(待商家同意),7(待买家退货),8(待商家发货),11(待商家二次同意),12(售后成功),13(换货待买家收货),14(换货成功),27(商家一次拒绝),28(售后失败),29(商家二次拒绝)
	AftersaleStatus int64 `json:"aftersale_status,omitempty"`
	// 售后理由；1-七天无理由退货；2-非七天无理由退货；
	Reason int64 `json:"reason,omitempty"`
	// 退货物流状态，枚举为1(全部),2(已发货),3(未发货)
	LogisticsStatus int64 `json:"logistics_status,omitempty"`
	// 付款方式，枚举为1(全部), 2(货到付款),3(线上付款)
	PayType int64 `json:"pay_type,omitempty"`
	// 退款类型，枚举为0(原路退款),1(线下退款),2(备用金),3(保证金),4(无需退款)
	RefundType int64 `json:"refund_type,omitempty"`
	// 仲裁状态，枚举为0(未介入),1(客服处理中),2(仲裁结束-支持买家),3(仲裁结束-支持卖家),4(待商家举证),5(待与买家协商),6(仲裁结束),255(取消)
	ArbitrateStatus []int64 `json:"arbitrate_status,omitempty"`
	// 插旗信息：0：灰 1：紫 2: 青 3：绿 4： 橙 5： 红
	OrderFlag []int64 `json:"order_flag,omitempty"`
	// 申请时间开始，单位为秒
	StartTime int64 `json:"start_time,omitempty"`
	// 申请时间结束，单位为秒
	EndTime int64 `json:"end_time,omitempty"`
	// 金额下限，单位为分
	AmountStart int64 `json:"amount_start,omitempty"`
	// 金额上限，单位为分
	AmountEnd int64 `json:"amount_end,omitempty"`
	// 风控标签，枚举为-1(退货正常),1(退货异常)
	RiskFlag int64 `json:"risk_flag,omitempty"`
	// 排序方式，优先级按照列表顺序从前往后依次减小，写法为"<字段名称> <排序方式>"，字段名称目前支持"status_deadline"（逾期时间）、"apply_time"（申请时间）和 "update_time"（更新时间），排序方式目前支持"asc"（升序）和"desc"（降序）。按照"逾期时间"升序排列，会优先返回临近逾期时间的数据。
	OrderBy []string `json:"order_by,omitempty"`
	// 页数，从0开始
	Page *int64 `json:"page,omitempty"`
	// 每页数量，最多100个
	Size *int64 `json:"size,omitempty"`
	// 售后单号
	AftersaleId string `json:"aftersale_id,omitempty"`
	// 售后状态；6-待商家同意；7-待买家退货；8-待商家发货；11-待商家二次同意；12-售后成功；13-换货\补寄\维修待买家收货；14-换货\补寄\维修成功；27-商家一次拒绝；28-售后失败；29-商家二次拒绝；支持传多种状态，使用英文“,”分隔
	StandardAftersaleStatus []int64 `json:"standard_aftersale_status,omitempty"`
	// 是否展示特殊售后
	NeedSpecialType bool `json:"need_special_type,omitempty"`
	// 更新时间开始，单位为秒
	UpdateStartTime int64 `json:"update_start_time,omitempty"`
	// 更新时间结束，单位为秒
	UpdateEndTime int64 `json:"update_end_time,omitempty"`
	// 正向物流单号
	OrderLogisticsTrackingNo []string `json:"order_logistics_tracking_no,omitempty"`
	// 正向物流状态（仅支持拒签场景下的状态筛选，状态更新有一定时延。1：买家已拒签；2：买家已签收；3：快递退回中，运往商家，包含快递拦截成功；4：商家已签收）
	OrderLogisticsState []int64 `json:"order_logistics_state,omitempty"`
	// 是否拒签后退款（1：已同意拒签, 2：未同意拒签）
	AgreeRefuseSign []int64 `json:"agree_refuse_sign,omitempty"`
	// 售后子类型；8001-以换代修。
	AftersaleSubType int64 `json:"aftersale_sub_type,omitempty"`
}
