package afterSale_List_response

import (
	"doudian.com/open/sdk_golang/core"
)

type AfterSaleListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *AfterSaleListData `json:"data"`
}
type SkuSpecItem struct {
	// 规格类型名称
	Name string `json:"name"`
	// 规格值
	Value string `json:"value"`
}
type AftersaleTagsItem struct {
	// 标签名称
	TagDetail string `json:"tag_detail"`
	// 标签关键字
	TagDetailEn string `json:"tag_detail_en"`
	// 标签悬浮文案（其中${key}占位符对应placeholder中的key对应内容）
	TagDetailText string `json:"tag_detail_text"`
	// 标签跳转链接
	TagLinkUrl string `json:"tag_link_url"`
	// 标签悬浮文案的占位符定义
	Placeholder map[string]PlaceholderItem `json:"placeholder"`
}
type AftersaleInfo struct {
	// 售后单号
	AftersaleId string `json:"aftersale_id"`
	// 售后订单类型，枚举为-1(历史订单),1(商品单),2(店铺单)
	AftersaleOrderType int64 `json:"aftersale_order_type"`
	// 售后类型，枚举为0(退货退款),1(已发货仅退款),2(未发货仅退款),3(换货),6(价保),7(补寄)
	AftersaleType int64 `json:"aftersale_type"`
	// 售后状态和请求参数standard_aftersale_status字段对应；3-换货待买家收货；6-待商家同意；7-待买家退货；8-待商家发货；11-待商家二次同意；12-售后成功；14-换货成功；27-商家一次拒绝；28-售后失败；29-商家二次拒绝；
	AftersaleStatus int64 `json:"aftersale_status"`
	// 关联的订单ID
	RelatedId string `json:"related_id"`
	// 申请时间
	ApplyTime int64 `json:"apply_time"`
	// 最近更新时间
	UpdateTime int64 `json:"update_time"`
	// 当前节点逾期时间
	StatusDeadline int64 `json:"status_deadline"`
	// 售后退款金额，单位为分
	RefundAmount int64 `json:"refund_amount"`
	// 售后退运费金额，单位为分
	RefundPostAmount int64 `json:"refund_post_amount"`
	// 售后数量
	AftersaleNum int64 `json:"aftersale_num"`
	// 部分退类型
	PartType int64 `json:"part_type"`
	// 售后退款类型，枚举为-1(历史数据默认值),0(订单货款/原路退款),1(货到付款线下退款),2(备用金),3(保证金),4(无需退款),5(平台垫付)
	AftersaleRefundType int64 `json:"aftersale_refund_type"`
	// 退款方式，枚举为1(极速退款助手)、2(售后小助手)、3(售后急速退)、4(闪电退货)
	RefundType int64 `json:"refund_type"`
	// 仲裁状态，枚举为0(无仲裁记录),1(仲裁中),2(客服同意),3(客服拒绝),4(待商家举证),5(协商期),255(仲裁结束)
	ArbitrateStatus int64 `json:"arbitrate_status"`
	// 售后单创建时间
	CreateTime int64 `json:"create_time"`
	// 退税费
	RefundTaxAmount int64 `json:"refund_tax_amount"`
	// 商家剩余发送短信（催用户寄回）次数
	LeftUrgeSmsCount int64 `json:"left_urge_sms_count"`
	// 退货物流单号
	ReturnLogisticsCode string `json:"return_logistics_code"`
	// 风控码
	RiskDecisionCode int64 `json:"risk_decision_code"`
	// 风控理由
	RiskDecisionReason string `json:"risk_decision_reason"`
	// 风控描述
	RiskDecisionDescription string `json:"risk_decision_description"`
	// 退优惠金额
	ReturnPromotionAmount int64 `json:"return_promotion_amount"`
	// 退款状态；1-待退款;2-退款中;3-退款成功;4-退款失败;5-追缴成功;
	RefundStatus int64 `json:"refund_status"`
	// 仲裁责任方
	ArbitrateBlame int64 `json:"arbitrate_blame"`
	// 换货SKU信息
	ExchangeSkuInfo *ExchangeSkuInfo `json:"exchange_sku_info"`
	// 退货物流公司名称
	ReturnLogisticsCompanyName string `json:"return_logistics_company_name"`
	// 换货物流公司名称
	ExchangeLogisticsCompanyName string `json:"exchange_logistics_company_name"`
	// 售后商家备注
	Remark string `json:"remark"`
	// 买家是否收到货物，0表示未收到，1表示收到
	GotPkg int64 `json:"got_pkg"`
	// 商家首次发货的正向物流信息
	OrderLogistics []OrderLogisticsItem `json:"order_logistics"`
	// 是否拒签后退款（1：已同意拒签, 2：未同意拒签）
	IsAgreeRefuseSign int64 `json:"is_agree_refuse_sign"`
	// 用户申请售后时选择的二级原因标签
	ReasonSecondLabels []ReasonSecondLabelsItem `json:"reason_second_labels"`
	// 售后标签（含时效延长、风险预警、豁免体验分等标签）标签在平台侧会有更新，标签仅做展示使用，请勿作为系统判断依赖。
	AftersaleTags []AftersaleTagsItem `json:"aftersale_tags"`
	// 门店ID
	StoreId int64 `json:"store_id"`
	// 门店名称
	StoreName string `json:"store_name"`
	// 售后子类型；8001-以换代修。
	AftersaleSubType int64 `json:"aftersale_sub_type"`
}
type TextPart struct {
	// 正向物流发货状态文案
	LogisticsText string `json:"logistics_text"`
	// 售后状态文案
	AftersaleStatusText string `json:"aftersale_status_text"`
	// 售后类型文案
	AftersaleTypeText string `json:"aftersale_type_text"`
	// 退货物流发货状态文案
	ReturnLogisticsText string `json:"return_logistics_text"`
	// 售后退款类型文案
	AftersaleRefundTypeText string `json:"aftersale_refund_type_text"`
	// 售后理由文案
	ReasonText string `json:"reason_text"`
	// 坏单比例文案
	BadItemText string `json:"bad_item_text"`
	// 仲裁状态文案
	ArbitrateStatusText string `json:"arbitrate_status_text"`
}
type ItemsItem struct {
	// 售后信息
	AftersaleInfo *AftersaleInfo `json:"aftersale_info"`
	// 订单信息
	OrderInfo *OrderInfo `json:"order_info"`
	// 文案部分
	TextPart *TextPart `json:"text_part"`
	// 卖家插旗日志
	SellerLogs []SellerLogsItem `json:"seller_logs"`
}
type OrderLogisticsItem struct {
	// 物流单号
	TrackingNo string `json:"tracking_no"`
	// 物流公司名称
	CompanyName string `json:"company_name"`
	// 物流公司编码
	CompanyCode string `json:"company_code"`
	// 物流状态到达时间
	LogisticsTime int64 `json:"logistics_time"`
	// 正向物流状态
	LogisticsState int64 `json:"logistics_state"`
}
type PlaceholderItem struct {
	// 占位符文案
	Text string `json:"text"`
	// 占位符跳转链接
	Url string `json:"url"`
}
type TagsItem struct {
	// 标签中文名称
	TagDetail string `json:"tag_detail"`
	// 标签编号
	TagDetailEn string `json:"tag_detail_en"`
	// 标签链接
	TagLinkUrl string `json:"tag_link_url"`
}
type ExchangeSkuInfo struct {
	// 换货SkuID
	SkuId string `json:"sku_id"`
	// 换货SKU code
	Code string `json:"code"`
	// 换货数目
	Num int64 `json:"num"`
	// 商家编号
	OutSkuId string `json:"out_sku_id"`
	// 区域库存仓ID
	OutWarehouseId string `json:"out_warehouse_id"`
	// sku外部供应商编码供应商ID
	SupplierId string `json:"supplier_id"`
	// 商品图片url
	Url string `json:"url"`
	// 商品名称
	Name string `json:"name"`
	// 换货商品的价格，单位分
	Price string `json:"price"`
	// sku规格信息
	SpecDesc string `json:"spec_desc"`
}
type ReasonSecondLabelsItem struct {
	// 二级原因标签编号
	Code int64 `json:"code"`
	// 二级原因标签名称
	Name string `json:"name"`
}
type RelatedOrderInfoItem struct {
	// 商品单信息
	SkuOrderId string `json:"sku_order_id"`
	// 订单状态，枚举为2(未发货),3(已发货),5(已收货或已完成),255(已完成)
	OrderStatus int64 `json:"order_status"`
	// 付款金额
	PayAmount int64 `json:"pay_amount"`
	// 付运费金额
	PostAmount int64 `json:"post_amount"`
	// 购买数量
	ItemNum int64 `json:"item_num"`
	// 下单时间
	CreateTime int64 `json:"create_time"`
	// 税费
	TaxAmount int64 `json:"tax_amount"`
	// 是否为海外订单
	IsOverseaOrder int64 `json:"is_oversea_order"`
	// 商品名称
	ProductName string `json:"product_name"`
	// 商品ID
	ProductId int64 `json:"product_id"`
	// 商品图片
	ProductImage string `json:"product_image"`
	// 订单标签；标签在平台侧会有更新，标签仅做展示使用，请勿作为系统判断依赖。
	Tags []TagsItem `json:"tags"`
	// 商品规格
	SkuSpec []SkuSpecItem `json:"sku_spec"`
	// 商家SKU编码
	ShopSkuCode string `json:"shop_sku_code"`
	// 发货物流编码
	LogisticsCode string `json:"logistics_code"`
	// 售后退款金额
	AftersalePayAmount int64 `json:"aftersale_pay_amount"`
	// 售后退运费金额
	AftersalePostAmount int64 `json:"aftersale_post_amount"`
	// 售后退税费金额
	AftersaleTaxAmount int64 `json:"aftersale_tax_amount"`
	// 售后商品数量
	AftersaleItemNum int64 `json:"aftersale_item_num"`
	// 优惠券金额
	PromotionPayAmount int64 `json:"promotion_pay_amount"`
	// 价格
	Price int64 `json:"price"`
	// 【已废弃】正向物流公司名称,替代字段：aftersale_info.order_logistics.company_name字段
	LogisticsCompanyName string `json:"logistics_company_name"`
	// 赠品订单id
	GivenSkuOrderIds []string `json:"given_sku_order_ids"`
}
type OrderInfo struct {
	// 店铺单订单ID
	ShopOrderId string `json:"shop_order_id"`
	// 售后关联的订单信息
	RelatedOrderInfo []RelatedOrderInfoItem `json:"related_order_info"`
	// 订单插旗
	OrderFlag int64 `json:"order_flag"`
}
type SellerLogsItem struct {
	// 插旗日志内容
	Content string `json:"content"`
	// 插旗操作人
	OpName string `json:"op_name"`
	// 插旗时间（字符串格式）
	CreateTime string `json:"create_time"`
	// 插旗信息；0：灰 1：紫 2: 青 3：绿 4： 橙 5： 红
	Star int64 `json:"star"`
}
type AfterSaleListData struct {
	// 售后列表元素
	Items []ItemsItem `json:"items"`
	// 是否还有更多
	HasMore bool `json:"has_more"`
	// 当前搜索条件下，匹配到的总数量
	Total int64 `json:"total"`
	// 页码，从0开始
	Page int64 `json:"page"`
	// 当前返回售后数量
	Size int64 `json:"size"`
}
