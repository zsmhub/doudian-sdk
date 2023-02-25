package order_getShopAccountItem_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderGetShopAccountItemResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderGetShopAccountItemData `json:"data"`
}
type OrderGetShopAccountItemData struct {
	// 返回code 100000为成功，其他为失败
	Code string `json:"code"`
	// 返回描述
	CodeMsg string `json:"code_msg"`
	// 资金流水明细
	Data []DataItem `json:"data"`
	// 下一次请求参数start_index 的值
	NextStartIndex string `json:"next_start_index"`
	// 下一次请求参数start_time的值
	NextStartTime string `json:"next_start_time"`
	// 查询数量，请求参数size传了值得话为请求参数传的值，未传的话，值为100
	Size int32 `json:"size"`
	// 返回结果data的数量
	DataSize int32 `json:"data_size"`
	// 判断查询是否结束。0 未结束, 1 结束。未结束时，需要把next_start_index作为下一次请求的start_index,next_start_time作为下一次请求的start_time。
	IsEnd int32 `json:"is_end"`
}
type DataItem struct {
	// 动账时间
	BillTime string `json:"bill_time"`
	// 动账方向  0:入账 1:出账
	FundFlow int32 `json:"fund_flow"`
	// 动账方向描述  出账、入账
	FundFlowDesc string `json:"fund_flow_desc"`
	// 动账金额(分)
	AccountAmount int64 `json:"account_amount"`
	// 动账摘要
	AccountBillDesc string `json:"account_bill_desc"`
	// 动账摘要 0：其他 1：达人带货佣金 2：达人佣金退款 3：订单结算 4：部分结算 5：运费单结算 6：服务费返还 7：平台补贴扣回 8：退款失败退票 9：结算重复扣款调账 10：保证金退款 11：提现 12：提现退票 13：极速退款分账 14：小额打款(原因:补差价) 15：小额打款(原因:其他) 16：小额打款(原因:商品补偿) 17：小额打款(原因:运费补偿) 18：小额打款退票 19：已退款 20：运费险扣减货款 21：支付BIC服务费 22：BIC服务费退票  23: 原路退
	AccountBillDescTag int32 `json:"account_bill_desc_tag"`
	// 计费类型  0:全部 1:鲁班广告 2:精选联盟 3:值点商城 4:小店自卖 5:橙子建站 6:POI 7:抖+ 8:穿山甲 9:服务市场 10:服务市场外包客服 11:学浪
	BizType int32 `json:"biz_type"`
	// 计费类型描述
	BizTypeDesc string `json:"biz_type_desc"`
	// 订单号
	OrderId string `json:"order_id"`
	// 店铺单号
	ShopOrderId string `json:"shop_order_id"`
	// 售后编号
	AfterSaleServiceNo string `json:"after_sale_service_no"`
	// 下单时间
	BusinessOrderCreateTime string `json:"business_order_create_time"`
	// 商品ID
	ProductId string `json:"product_id"`
	// 订单实付应结(分)
	PayAmount int64 `json:"pay_amount"`
	// 实际平台补贴(分)
	PromotionAmount int64 `json:"promotion_amount"`
	// 订单退款(分)
	RefundAmount int64 `json:"refund_amount"`
	// 平台服务费(分)
	PlatformServiceFee int64 `json:"platform_service_fee"`
	// 佣金(分)
	Commission int64 `json:"commission"`
	// 渠道分成(分)
	ChannelFee int64 `json:"channel_fee"`
	// 招商服务费(分)
	ColonelServiceFee int64 `json:"colonel_service_fee"`
	// 动账账户  1: 微信 2:支付宝 3:合众支付 4:聚合支付
	AccountType int32 `json:"account_type"`
	// 动账账户
	AccountTypeDesc string `json:"account_type_desc"`
	// 实际达人补贴(分)
	AuthorCouponSubsidy int64 `json:"author_coupon_subsidy"`
	// 运费(分)
	PostAmount int64 `json:"post_amount"`
	// 动账流水号(唯一)
	AccountTradeNo string `json:"account_trade_no"`
	// 订单类型  0：普通订单  1：定金(已退款) 2：定金(尾款已支付) 3：定金(尾款未支付) 4：尾款(尾款已支付) 5：尾款(已退款)
	OrderType int32 `json:"order_type"`
	// 订单类型描述
	OrderTypeDesc string `json:"order_type_desc"`
	// 实际抖音支付补贴(分)
	ActualZtPayPromotion int64 `json:"actual_zt_pay_promotion"`
	// 实际DOU分期营销补贴(分)
	ActualZrPayPromotion int64 `json:"actual_zr_pay_promotion"`
	// 直播间站外推广(分)
	ChannelPromotionFee int64 `json:"channel_promotion_fee"`
	// 其他分成金额(分)
	OtherSharingAmount int64 `json:"other_sharing_amount"`
	// 备注
	Remark string `json:"remark"`
	// 店铺id
	ShopId string `json:"shop_id"`
}
