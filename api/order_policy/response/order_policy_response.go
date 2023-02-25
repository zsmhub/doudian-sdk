package order_policy_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderPolicyResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderPolicyData `json:"data"`
}
type ClaimInfoListItem struct {
	// 理赔状态
	Status int32 `json:"status"`
	// 预计理赔金额：分
	Amount int64 `json:"amount"`
	// 实际理赔金额
	Premium int64 `json:"premium"`
	// 理赔时间
	ClaimTime int64 `json:"claim_time"`
	// 理赔单号
	InsClaimNo string `json:"ins_claim_no"`
	// 申请原因
	ClaimMsg string `json:"claim_msg"`
	// 拒绝原因
	RefusedMsg string `json:"refused_msg"`
	// 聚合后的理赔状态：绿植险有用
	AggClaimStatus int64 `json:"agg_claim_status"`
	// 理赔次数
	ClaimAppliedTimes int32 `json:"claim_applied_times"`
}
type OrderPolicyData struct {
	// 保单详情列表
	PolicyInfo *PolicyInfo `json:"policy_info"`
	// 理赔列表
	ClaimInfoList []ClaimInfoListItem `json:"claim_info_list"`
}
type GoodsInfoListItem struct {
	// 商品名称
	Name string `json:"name"`
	// 商品id
	ProductId int64 `json:"product_id"`
	// 类目
	CategoryId string `json:"category_id"`
	// url
	ShowPageUrl string `json:"show_page_url"`
	// 数量
	Count int32 `json:"count"`
}
type PolicyInfo struct {
	// 保单号
	InsPolicyNo string `json:"ins_policy_no"`
	// 预计退保费用，单位分
	Amount int64 `json:"amount"`
	// 总保费，单位分
	Premium int64 `json:"premium"`
	// 用户支付的保费，单位分
	UserPremium int64 `json:"user_premium"`
	// 商家支付的保费，单位分
	MerchantPremium int64 `json:"merchant_premium"`
	// 平台承担的保费，单位分
	PlatformPremium int64 `json:"platform_premium"`
	// 保单状态 0: 初始化 1: 待生效 2: 保障中 3: 赔付完成 4: 保单失效 5: 已经取消
	Status int32 `json:"status"`
	// 理赔状态 0: 初始化 1: 理赔中 2: 理赔成功 3: 理赔失败  绿植险该字段没用
	ClaimStatus int32 `json:"claim_status"`
	// 申述状态 0: 初始化 1: 申诉处理中 2: 申诉成功 3: 申诉失败  绿植险该字段没用
	AppealStatus int32 `json:"appeal_status"`
	// 商品列表
	GoodsInfoList []GoodsInfoListItem `json:"goods_info_list"`
	// 出保时间
	InsEnsuredTimeBegin int64 `json:"ins_ensured_time_begin"`
	// 保险过期时间
	InsEnsuredTimeEnd int64 `json:"ins_ensured_time_end"`
	// 只有在保单状态为赔付失败的时候并允许申诉 true，其他情况 false
	IsAllowAppeal bool `json:"is_allow_appeal"`
	// 理赔或申述失败原因
	RefusedMsg string `json:"refused_msg"`
	// 保险客服电话
	InsHotline string `json:"ins_hotline"`
	// 投保人  2:商家  3:平台
	PayerType int32 `json:"payer_type"`
}
