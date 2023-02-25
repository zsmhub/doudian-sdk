package order_getSettleBillDetailV3_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderGetSettleBillDetailV3Response struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderGetSettleBillDetailV3Data `json:"data"`
}
type DataItem struct {
	// 结算时间
	SettleTime string `json:"settle_time"`
	// 结算单号
	RequestNo string `json:"request_no"`
	// 订单号
	ShopOrderId string `json:"shop_order_id"`
	// 子订单号
	OrderId string `json:"order_id"`
	// 商家实收（分）
	SettleAmount int64 `json:"settle_amount"`
	// 货款结算对应的账户类型： “聚合账户”“微信”“支付宝”“微信升级前”“合众支付”等
	PayTypeDesc string `json:"pay_type_desc"`
	// 结算单类型 0 ：已结算 1 ：结算后退款-原路退回 2： 保证金退款-支出退回 3： 结算后退款-非原路退回
	TradeType int32 `json:"trade_type"`
	// 是否包含结算前退款 0：不包含 1：包含
	IsContainsRefundBeforeSettle int32 `json:"is_contains_refund_before_settle"`
	// 下单时间
	OrderTime string `json:"order_time"`
	// 商品id
	ProductId string `json:"product_id"`
	// 商品数量
	GoodsCount int32 `json:"goods_count"`
	// 业务类型: 鲁班广告、商城、精选联盟、小店自卖等
	FlowTypeDesc string `json:"flow_type_desc"`
	// 订单类型：普通订单、尾款(尾款已支付)、尾款(已退款)、定金(已退款)、定金(尾款已支付)、定金(尾款未支付)
	OrderType string `json:"order_type"`
	// 订单总价（分）
	TotalAmount int64 `json:"total_amount"`
	// 商品总价（分）
	TotalGoodsAmount int64 `json:"total_goods_amount"`
	// 运费（分）
	PostAmount int64 `json:"post_amount"`
	// 店铺券（分）
	ShopCoupon int64 `json:"shop_coupon"`
	// 结算前退款金额（分） （结算前退货+运费-店铺券）
	RefundBeforeSettle int64 `json:"refund_before_settle"`
	// 平台补贴（分）
	PlatformCoupon int64 `json:"platform_coupon"`
	// 达人补贴（分）
	AuthorCoupon int64 `json:"author_coupon"`
	// 抖音支付补贴（分）
	ZtPayPromotion int64 `json:"zt_pay_promotion"`
	// DOU分期营销补贴（分）
	ZrPayPromotion int64 `json:"zr_pay_promotion"`
	// 用户实付（分）
	RealPayAmount int64 `json:"real_pay_amount"`
	// 收入合计（分）
	TotalIncome int64 `json:"total_income"`
	// 平台服务费（分）
	PlatformServiceFee int64 `json:"platform_service_fee"`
	// 佣金（分）
	Commission int64 `json:"commission"`
	// 渠道分成（分）
	GoodLearnChannelFee int64 `json:"good_learn_channel_fee"`
	// 团长服务费（分）
	ColonelServiceFee int64 `json:"colonel_service_fee"`
	// 直播间站外推广（分）
	ChannelPromotionFee int64 `json:"channel_promotion_fee"`
	// 其他分成（分）
	OtherSharingAmount int64 `json:"other_sharing_amount"`
	// 合计支出
	TotalOutcome int64 `json:"total_outcome"`
	// 备注
	Remark string `json:"remark"`
	// 打包费，单位：分
	PackingAmount int64 `json:"packing_amount"`
}
type OrderGetSettleBillDetailV3Data struct {
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
	// 判断查询是否结束。0 未结束, 1 结束。未结束时，需要把next_start_index作为下一次请求的start_index,next_start_time作为下一次请求的start_time
	IsEnd int32 `json:"is_end"`
	// 结果data的大小
	DataSize int64 `json:"data_size"`
}
