package order_serviceDetail_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderServiceDetailResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderServiceDetailData `json:"data"`
}
type OrderServiceDetailData struct {
	// 详情
	Detail *Detail `json:"detail"`
	// 日志
	Logs []LogsItem `json:"logs"`
}
type Detail struct {
	// 服务请求ID 
	Id int64 `json:"id"`
	// 订单号
	OrderId int64 `json:"order_id"`
	// 操作状态，增加审核中状态码
	OperateStatus int32 `json:"operate_status"`
	// 服务单详情
	Detail string `json:"detail"`
	// 商家答复内容
	Reply string `json:"reply"`
	// 服务单创建时间 
	CreateTime string `json:"create_time"`
	// 服务单类型，枚举
	ServiceType int64 `json:"service_type"`
	// 最新回复时间
	ReplyTime string `json:"reply_time"`
	// 操作状态含义，增加审核中状态
	OperateStatusDesc string `json:"operate_status_desc"`
	// 店铺id
	ShopId int64 `json:"shop_id"`
	// 是否为被审核驳回的服务单
	IsReject int64 `json:"is_reject"`
	// 驳回内容
	RejectDetail string `json:"reject_detail"`
	// 驳回时间
	RejectTime string `json:"reject_time"`
	// 驳回时间
	RejectTime string `json:"reject_time"`
	// 凭证demo url
	ProofDemo string `json:"proof_demo"`
	// 是否必须上传凭证
	EvidenceRequired int64 `json:"evidence_required"`
	// 图片url 数组json
	ImgList []string `json:"img_list"`
	// 超时时间
	ExpirationTime string `json:"expiration_time"`
	// 关闭时间
	CloseTime string `json:"close_time"`
	// 关闭原因 
	CloseDetail string `json:"close_detail"`
	// 首次回复时间
	FirstReplyTime string `json:"first_reply_time"`
	// 服务单更新时间时间
	UpdateTime string `json:"update_time"`
}
type LogsItem struct {
	// 内容
	Content string `json:"content"`
	// 图片列表
	ImgList []string `json:"img_list"`
	// 日志类型
	ServiceLogType int32 `json:"service_log_type"`
	// 平台客服
	OperateName string `json:"operate_name"`
	// 服务详情
	ServiceDetail string `json:"service_detail"`
	// 回复内容
	ReplyDetail string `json:"reply_detail"`
	// 关闭原因 
	CloseDetail string `json:"close_detail"`
	// 驳回内容
	RejectDetail string `json:"reject_detail"`
	// 24小时
	DealTime string `json:"deal_time"`
}
