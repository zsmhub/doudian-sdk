package order_getSettleBillDetailV2_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderGetSettleBillDetailV2Response struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderGetSettleBillDetailV2Data `json:"data"`
}
type DataItem struct {
	// 店铺单号（主订单号）
	ShopOrderId string `json:"shop_order_id"`
	// SKU单（子订单号）
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
	// 结算状态 0:已结算 1:已退款
	TradeType string `json:"trade_type"`
	// 结算账户类型 1:微信（升级前）, 2:微信，3:支付宝, 4:合众支付 5:聚合账户
	PayType string `json:"pay_type"`
	// 结算单号,每一条订单流水明细结算单号唯一
	RequestNo string `json:"request_no"`
	// 业务类型 1:鲁班广告, 2:精选联盟 ,3:值点商城, 4:小店自卖
	FlowType string `json:"flow_type"`
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
	// 渠道分成(好好学习分成)(分)
	GoodLearnChannelFee int64 `json:"good_learn_channel_fee"`
	// 招商服务费(分)
	ColonelServiceFee int64 `json:"colonel_service_fee"`
	// 退款扣佣金(退佣失败垫付金额)(分)
	ShopRefundLoss int64 `json:"shop_refund_loss"`
	// 结算状态描述
	TradeTypeDesc string `json:"trade_type_desc"`
	// 结算账户类型 微信（升级前）,微信，支付宝,聚合账户,合众支付
	PayTypeDesc string `json:"pay_type_desc"`
	// 业务类型 广告,频道,联盟,免费
	FlowTypeDesc string `json:"flow_type_desc"`
	// 结算前退款(分)
	RefundBeforeSettlement int64 `json:"refund_before_settlement"`
	// 实际抖音支付补贴金额(分)
	ActualZtPayPromotion int64 `json:"actual_zt_pay_promotion"`
	// 实际DOU分期支付补贴金额(分)
	ActualZrPayPromotion int64 `json:"actual_zr_pay_promotion"`
	// 抖音支付支付补贴(分)
	ZtPayPromotion int64 `json:"zt_pay_promotion"`
	// DOU分期支付补贴(分)
	ZrPayPromotion int64 `json:"zr_pay_promotion"`
	// 直播间站外推广(外部渠道推广费)(分)
	ChannelPromotionFee int64 `json:"channel_promotion_fee"`
	// 其他分成金额(分)
	OtherSharingAmount int64 `json:"other_sharing_amount"`
	// 备注
	Remark string `json:"remark"`
}
type OrderGetSettleBillDetailV2Data struct {
	// 返回code 100000为成功，其他为失败
	Code string `json:"code"`
	// 返回信息描述，失败状态下会有失败描述
	CodeMsg string `json:"code_msg"`
	// 订单流水明细列表
	Data []DataItem `json:"data"`
	// 请求的size
	Size int64 `json:"size"`
	// 下一次查询start_index
	NextStartIndex string `json:"next_start_index"`
	// 下一次查询start_time
	NextStartTime string `json:"next_start_time"`
	// 结果data的大小
	DataSize int64 `json:"data_size"`
	// 判断查询是否结束。0 未结束, 1 结束。未结束时，需要把next_start_index作为下一次请求的start_index,next_start_time作为下一次请求的start_time。
	IsEnd int32 `json:"is_end"`
}
