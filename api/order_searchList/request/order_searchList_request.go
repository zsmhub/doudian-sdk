package order_searchList_request

import (
	"doudian.com/open/sdk_golang/api/order_searchList/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderSearchListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderSearchListParam 
}
func (c *OrderSearchListRequest) GetUrlPath() string{
	return "/order/searchList"
}


func New() *OrderSearchListRequest{
	request := &OrderSearchListRequest{
		Param: &OrderSearchListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderSearchListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_searchList_response.OrderSearchListResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_searchList_response.OrderSearchListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderSearchListRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderSearchListRequest) GetParams() *OrderSearchListParam{
	return c.Param
}


type CombineStatusItem struct {
	// 订单状态，","分隔多个状态；1-待支付；101-部分发货；105-已支付；2-待发货；3-已发货；4-已取消；5-已完成
	OrderStatus string `json:"order_status,omitempty"`
	// 【已废弃，不推荐使用】主流程状态，","分隔多个状态；
	MainStatus string `json:"main_status,omitempty"`
}
type OrderSearchListParam struct {
	// 商品，number型代表商品ID，其它代表商品名称
	Product string `json:"product,omitempty"`
	// 【下单端】 0、站外   1、火山   2、抖音   3、头条   4、西瓜   5、微信   6、值点app  7、头条lite   8、懂车帝  9、皮皮虾   11、抖音极速版   12、TikTok   13、musically   14、穿山甲   15、火山极速版   16、服务市场   26、番茄小说   27、UG教育营销电商平台   28、Jumanji   29、电商SDK
	BType int64 `json:"b_type,omitempty"`
	// 售后状态：all-全部，in_aftersale-售后中，refund-退款中，refund_success-退款成功，refund_fail-退款失败，exchange_success-换货成功 aftersale_close-售后关闭
	AfterSaleStatusDesc string `json:"after_sale_status_desc,omitempty"`
	// 物流单号
	TrackingNo string `json:"tracking_no,omitempty"`
	// 预售类型：0-普通订单；1-全款预售；2-定金预售；3-定金找货；
	PresellType int64 `json:"presell_type,omitempty"`
	// 订单类型 0、普通订单 2、虚拟商品订单 4、电子券（poi核销） 5、三方核销
	OrderType int64 `json:"order_type,omitempty"`
	// 下单时间：开始，秒级时间戳
	CreateTimeStart int64 `json:"create_time_start,omitempty"`
	// 下单时间：截止，秒级时间戳
	CreateTimeEnd int64 `json:"create_time_end,omitempty"`
	// 异常订单，1-异常取消，2-风控审核中
	AbnormalOrder int64 `json:"abnormal_order,omitempty"`
	// 交易类型：0-普通；1-拼团；2-定金预售；3-订金找货；4-拍卖；5-0元单；6-回收；7-寄卖；10-寄样；11-0元抽奖(超级福袋)；12-达人买样；13-普通定制；16-大众竞拍；18-小时达；102-定金预售的赠品单；103-收款；
	TradeType int64 `json:"trade_type,omitempty"`
	// 状态组合查询，直接输入状态码（只支持一个元素）
	CombineStatus []CombineStatusItem `json:"combine_status,omitempty"`
	// 更新时间：开始
	UpdateTimeStart int64 `json:"update_time_start,omitempty"`
	// 更新时间：截止
	UpdateTimeEnd int64 `json:"update_time_end,omitempty"`
	// 单页大小，限制100以内
	Size *int64 `json:"size,omitempty"`
	// 页码，0页开始
	Page *int64 `json:"page,omitempty"`
	// 排序条件(create_time 订单创建时间；update_time 订单更新时间；默认create_time；)
	OrderBy string `json:"order_by,omitempty"`
	// 排序类型，小到大或大到小，默认大到小
	OrderAsc bool `json:"order_asc,omitempty"`
	// 履约状态；如小时达未接单"no_accept"
	FulfilStatus string `json:"fulfil_status,omitempty"`
	IsSearchable bool `json:"is_searchable,omitempty,omitempty"`
}
