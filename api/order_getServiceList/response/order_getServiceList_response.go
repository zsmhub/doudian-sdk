package order_getServiceList_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderGetServiceListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderGetServiceListData `json:"data"`
}
type ListItem struct {
	// 服务请求ID
	Id int64 `json:"id"`
	// 订单号
	OrderId int64 `json:"order_id"`
	// 操作状态
	OperateStatus int32 `json:"operate_status"`
	// 服务单详情
	Detail string `json:"detail"`
	// 回复内容
	Reply string `json:"reply"`
	// 服务创建时间
	CreateTime string `json:"create_time"`
	// 服务类型
	ServiceType int64 `json:"service_type"`
	// 回复时间
	ReplyTime string `json:"reply_time"`
	// 操作状态含义
	OperateStatusDesc string `json:"operate_status_desc"`
	// 店铺id
	ShopId int64 `json:"shop_id"`
	// 服务更新时间时间
	UpdateTime string `json:"update_time"`
}
type OrderGetServiceListData struct {
	// 总数
	Total int64 `json:"total"`
	// 列表
	List []ListItem `json:"list"`
}
