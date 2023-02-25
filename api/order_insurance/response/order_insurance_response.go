package order_insurance_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderInsuranceResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderInsuranceData `json:"data"`
}
type OrderInsuranceData struct {
	// 保单详情列表
	PolicyInfoList []PolicyInfoListItem `json:"policy_info_list"`
}
type GoodsInfoListItem struct {
	// 商品名称
	ProductName string `json:"product_name"`
	// 商品id
	ProductId int64 `json:"product_id"`
	// 类目
	CategoryId string `json:"category_id"`
	// url
	ShowPageUrl string `json:"show_page_url"`
	// 数量
	Amount int32 `json:"amount"`
}
type PolicyInfoListItem struct {
	// 保单号
	InsPolicyNo string `json:"ins_policy_no"`
	// 预计退保费用，单位分
	ApproximateReturnFee int64 `json:"approximate_return_fee"`
	// 总保费，单位分
	TotalPremium int64 `json:"total_premium"`
	// 用户支付的保费，单位分
	UserPremium int64 `json:"user_premium"`
	// 商家支付的保费，单位分
	MerchantPremium int64 `json:"merchant_premium"`
	// 平台承担的保费，单位分
	PlatformPremium int64 `json:"platform_premium"`
	// 保单状态 0: 初始化 1: 待生效 2: 保障中 3: 赔付完成 4: 保单失效 5: 已经取消
	Status int32 `json:"status"`
	// 申述状态 0: 初始化 1: 申诉处理中 2: 申诉成功 3: 申诉失败
	AppealStatus int32 `json:"appeal_status"`
	// 理赔状态 0: 初始化 1: 理赔中 2: 理赔成功 3: 理赔失败
	ClaimStatus int32 `json:"claim_status"`
	// 保险公司名称
	CompanyName string `json:"company_name"`
	// 只有在保单状态为赔付失败的时候并允许申诉 true，其他情况 false
	IsAllowAppeal bool `json:"is_allow_appeal"`
	// 商品列表
	GoodsInfoList []GoodsInfoListItem `json:"goods_info_list"`
	// 理赔或申述失败原因
	FailReason string `json:"fail_reason"`
	// 保险客服电话
	InsuranceHotline string `json:"insurance_hotline"`
	// 保费状态 1：待扣减 2：已扣减 3：已退还 4：已垫付 5：已关闭
	PremiumStatus int32 `json:"premium_status"`
	// 出保时间
	InsEnsuredTimeBegin string `json:"ins_ensured_time_begin"`
	// 保险过期时间
	InsEnsuredTimeEnd string `json:"ins_ensured_time_end"`
}
