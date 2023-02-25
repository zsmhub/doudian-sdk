package order_getSettleBillDetail_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderGetSettleBillDetailResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderGetSettleBillDetailData `json:"data"`
}
type DataItem struct {
	// 店铺单号
	ShopOrderId string `json:"shop_order_id"`
	// 订单号
	OrderId string `json:"order_id"`
	// 下单时间
	OrderTime string `json:"order_time"`
	// 店铺id
	ShopId int64 `json:"shop_id"`
	// 结算时间
	SettleTime string `json:"settle_time"`
	// 商品ID
	ProductId string `json:"product_id"`
	// 商品数量
	GoodsCount int32 `json:"goods_count"`
	// 结算状态描述
	TradeTypeDesc string `json:"trade_type_desc"`
	// 结算账户类型 微信（升级前）,微信，支付宝,聚合账户,合众支付
	PayTypeDesc string `json:"pay_type_desc"`
	// 结算单号
	RequestNo string `json:"request_no"`
	// 业务类型  广告,频道,联盟,免费
	FlowTypeDesc string `json:"flow_type_desc"`
	// 阶段单号
	PhaseOrderNo string `json:"phase_order_no"`
	// 有多少个阶段，比如：预售业务单有2个阶段，阶段1：支付定金阶段，阶段2：支付尾款
	PhaseCnt int32 `json:"phase_cnt"`
	// 当前在第几阶段（从1开始），比如：预售定金支付为阶段1，尾款支付为阶段2
	PhaseId int32 `json:"phase_id"`
	// 总收入(分)
	TotalIncome int64 `json:"total_income"`
	// 总支出(分)
	TotalOutcome int64 `json:"total_outcome"`
	// 收益(分)
	Profit int64 `json:"profit"`
	// 实际结算金额(分)
	SettleAmount int64 `json:"settle_amount"`
	// 实际补贴金额(分)
	ActualSubsidyAmount int64 `json:"actual_subsidy_amount"`
	// 订单总价(分)
	TotalAmount int64 `json:"total_amount"`
	// 商品总价(分)
	TotalGoodsAmount int64 `json:"total_goods_amount"`
	// 运费(分)
	PostAmount int64 `json:"post_amount"`
	// 订单实付(分)
	RealPayAmount int64 `json:"real_pay_amount"`
	// 订单实付应结(分)
	SettledPayAmount int64 `json:"settled_pay_amount"`
	// 平台补贴(分)
	PlatformCoupon int64 `json:"platform_coupon"`
	// 达人券补贴(分)
	AuthorCoupon int64 `json:"author_coupon"`
	// 支付补贴(分)
	PayPromotion int64 `json:"pay_promotion"`
	// 实际平台补贴(分)
	ActualPlatformCoupon int64 `json:"actual_platform_coupon"`
	// 实际达人券补贴(分)
	ActualAuthorCoupon int64 `json:"actual_author_coupon"`
	// 实际支付补贴(分)
	ActualPayPromotion int64 `json:"actual_pay_promotion"`
	// 店铺优惠(分)
	ShopCoupon int64 `json:"shop_coupon"`
	// 平台服务费(分)
	PlatformServiceFee int64 `json:"platform_service_fee"`
	// 订单退款(分)
	Refund int64 `json:"refund"`
	// 佣金(分)
	Commission int64 `json:"commission"`
	// 好好学习分成(分)
	GoodLearnChannelFee int64 `json:"good_learn_channel_fee"`
	// 招商服务费(分)
	ColonelServiceFee int64 `json:"colonel_service_fee"`
	// 退佣失败垫付金额(分)
	ShopRefundLoss int64 `json:"shop_refund_loss"`
	// 结算状态  0:已结算   1:已退款
	TradeType string `json:"trade_type"`
	// 结算账户类型   1:微信（升级前）, 2:微信，3:支付宝, 4:合众支付 5:聚合账户
	PayType string `json:"pay_type"`
	// 业务类型  1:鲁班广告, 2:值点商城, 3:精选联盟  4:小店自卖
	FlowType string `json:"flow_type"`
}
type OrderGetSettleBillDetailData struct {
	// 返回code   100000为成功，其他为失败
	Code string `json:"code"`
	// 返回信息描述，成功时返回SUCCESS，其他状态下会有失败描述
	CodeMsg string `json:"code_msg"`
	// 账单明细列表
	Data []DataItem `json:"data"`
	// 总数
	TotalCnt int64 `json:"total_cnt"`
	// 页数
	Page int64 `json:"page"`
	// 每页数量
	Size int64 `json:"size"`
}
