package order_searchList_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderSearchListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderSearchListData `json:"data"`
}
type SkuCustomizationInfoItem struct {
	// 定制详情
	Detail *Detail `json:"detail"`
}
type MaskPostAddr struct {
	// 省
	Province *Province `json:"province"`
	// 市
	City *City `json:"city"`
	// 县/区
	Town *Town `json:"town"`
	// 街道
	Street *Street `json:"street"`
	// 详细地址
	Detail string `json:"detail"`
}
type RealYunWarehouseInfo struct {
	// 云仓code
	YunWarehouseCode string `json:"yun_warehouse_code"`
	// 云仓名称
	YunWarehouseName string `json:"yun_warehouse_name"`
}
type ShopOrderTagUiItem struct {
	// 标签Key
	Key string `json:"key"`
	// 标签说明
	Text string `json:"text"`
	// 帮助文档地址
	HelpDoc string `json:"help_doc"`
}
type AddressTagUiItem struct {
	// 双地址标签 key
	Key string `json:"key"`
	// 双地址
	Text string `json:"text"`
	// 双地址 hover 提示
	HoverText string `json:"hover_text"`
}
type Town struct {
	// 名称
	Name string `json:"name"`
	// 地区ID
	Id string `json:"id"`
}
type SpecItem struct {
	// 规格名称
	Name string `json:"name"`
	// 规格值
	Value string `json:"value"`
}
type PreAllocatedYunWarehouseInfo struct {
	// 云仓code
	YunWarehouseCode string `json:"yun_warehouse_code"`
	// 云仓名称
	YunWarehouseName string `json:"yun_warehouse_name"`
}
type SkuOrderListItem struct {
	// 商品订单号
	OrderId string `json:"order_id"`
	// 父订单号（店铺订单号）
	ParentOrderId string `json:"parent_order_id"`
	// 订单层级
	OrderLevel int64 `json:"order_level"`
	// 【业务来源】 1、鲁班 2、小店 3、好好学习 4、ev 5、虚拟 6、建站 7、核销 8、玉石 9、ez 10、ep 11、虚拟卡券 12、服务市场 13、EP视频课 14、EP直播课 21、跨境BBC 22、跨境BC 23、跨境CC,UPC 24、手机充值 25、拍卖保证金 26、懂车帝抵扣券 27、懂车帝返现券 28、离岛免税 29、海南会员购 30、抽奖 32、dou+券 76、大闸蟹 99、保险 102、小店海外
	Biz int64 `json:"biz"`
	// 业务来源描述
	BizDesc string `json:"biz_desc"`
	// 【订单类型】 0、普通订单 2、虚拟商品订单 4、电子券（poi核销） 5、三方核销
	OrderType int64 `json:"order_type"`
	// 订单类型描述
	OrderTypeDesc string `json:"order_type_desc"`
	// 交易类型；0-普通；1-拼团；2-定金预售；3-订金找货；4-拍卖；5-0元单；6-回收；7-寄卖；10-寄样；11-零元抽奖；12-达人买样；13-普通定制；16-大众竞拍；18-小时达；102-定金预售的赠品单；103-收款；和trade_type_desc成对出现
	TradeType int64 `json:"trade_type"`
	// 交易类型描述；和trade_type成对出现
	TradeTypeDesc string `json:"trade_type_desc"`
	// 订单状态
	OrderStatus int64 `json:"order_status"`
	// 订单状态描述
	OrderStatusDesc string `json:"order_status_desc"`
	// 主流程状态
	MainStatus int64 `json:"main_status"`
	// 主流程状态描述
	MainStatusDesc string `json:"main_status_desc"`
	// 支付时间
	PayTime int64 `json:"pay_time"`
	// 订单过期时间
	OrderExpireTime int64 `json:"order_expire_time"`
	// 订单完成时间
	FinishTime int64 `json:"finish_time"`
	// 下单时间
	CreateTime int64 `json:"create_time"`
	// 订单更新时间
	UpdateTime int64 `json:"update_time"`
	// 取消原因
	CancelReason string `json:"cancel_reason"`
	// 【下单端】 0、站外 1、火山 2、抖音 3、头条 4、西瓜 5、微信 6、值点app 7、头条lite 8、懂车帝 9、皮皮虾 11、抖音极速版 12、TikTok 13、musically 14、穿山甲 15、火山极速版 16、服务市场 26、番茄小说 27、UG教育营销电商平台 28、Jumanji 29、电商SDK
	BType int64 `json:"b_type"`
	// 下单端描述
	BTypeDesc string `json:"b_type_desc"`
	// 【下单场景】 0、未知 1、app内-原生 2、app内-小程序 3、H5 13、电商SDK-头条 35、电商SDK-头条lite
	SubBType int64 `json:"sub_b_type"`
	// 下单场景描述
	SubBTypeDesc string `json:"sub_b_type_desc"`
	// 已废弃，不推荐使用。流量来源：1-鲁班广告 2-联盟 3-商城 4-自主经营 5-线索通支付表单 6-抖音门店 7-抖+ 8-穿山甲
	SendPay int64 `json:"send_pay"`
	// 已废弃，不推荐使用。流量来源描述
	SendPayDesc string `json:"send_pay_desc"`
	// 直播主播id（达人）
	AuthorId int64 `json:"author_id"`
	// 直播主播名称
	AuthorName string `json:"author_name"`
	// 【下单来源】 0、其他 1、直播间
	ThemeType string `json:"theme_type"`
	// 下单来源描述
	ThemeTypeDesc string `json:"theme_type_desc"`
	// 小程序id
	AppId int64 `json:"app_id"`
	// 直播间id
	RoomId int64 `json:"room_id"`
	// 内容id
	ContentId string `json:"content_id"`
	// 视频id
	VideoId string `json:"video_id"`
	// 流量来源id
	OriginId string `json:"origin_id"`
	// 创意id（creative_id）
	Cid int64 `json:"cid"`
	// 【C端流量来源】 0、unknown 2、精选联盟 8、小店自卖
	CBiz int64 `json:"c_biz"`
	// C端流量来源业务类型描述
	CBizDesc string `json:"c_biz_desc"`
	// 鲁班广告落地页ID
	PageId int64 `json:"page_id"`
	// 【支付类型】 0、货到付款 1 、微信 2、支付宝 3、小程序 4、银行卡 5、余额 7、无需支付（0元单） 8、DOU分期（信用支付） 9、新卡支付
	PayType int64 `json:"pay_type"`
	// 支付渠道的流水号
	ChannelPaymentNo string `json:"channel_payment_no"`
	// 订单金额（单位：分）
	OrderAmount int64 `json:"order_amount"`
	// 支付金额（单位：分）
	PayAmount int64 `json:"pay_amount"`
	// 运费险金额（单位：分）
	PostInsuranceAmount int64 `json:"post_insurance_amount"`
	// 改价金额变化量（单位：分）
	ModifyAmount int64 `json:"modify_amount"`
	// 改价运费金额变化量（单位：分）
	ModifyPostAmount int64 `json:"modify_post_amount"`
	// 订单优惠总金额（单位：分）= 店铺优惠金额+ 平台优惠金额+ 达人优惠金额
	PromotionAmount int64 `json:"promotion_amount"`
	// 店铺优惠金额（单位：分），属于店铺的优惠活动、优惠券、红包的总优惠金额
	PromotionShopAmount int64 `json:"promotion_shop_amount"`
	// 平台优惠金额（单位：分），属于平台的优惠活动、优惠券、红包的总优惠金额
	PromotionPlatformAmount int64 `json:"promotion_platform_amount"`
	// 商家承担金额（单位：分），订单参与活动和优惠中商家承担部分的总金额
	ShopCostAmount int64 `json:"shop_cost_amount"`
	// 平台承担金额（单位：分），订单参与活动和优惠中平台+作者（达人）承担部分的总金额,包含作者（达人）承担金额：platform_cost_amount = only_platform_cost_amount + author_cost_amount
	PlatformCostAmount int64 `json:"platform_cost_amount"`
	// 达人优惠金额（单位：分），属于达人的优惠活动、优惠券、红包的总优惠金额
	PromotionTalentAmount int64 `json:"promotion_talent_amount"`
	// 支付优惠金额（单位：分），支付渠道上的优惠金额
	PromotionPayAmount int64 `json:"promotion_pay_amount"`
	// 商家后台商品编码
	Code string `json:"code"`
	// 收件人电话
	PostTel string `json:"post_tel"`
	EncryptPostTel string `json:"encrypt_post_tel"`
	// 收件人姓名
	PostReceiver string `json:"post_receiver"`
	EncryptPostReceiver string `json:"encrypt_post_receiver"`
	// 收件人地址
	PostAddr *PostAddr `json:"post_addr"`
	// 预计发货时间
	ExpShipTime int64 `json:"exp_ship_time"`
	// 发货时间
	ShipTime int64 `json:"ship_time"`
	// 物流收货时间
	LogisticsReceiptTime int64 `json:"logistics_receipt_time"`
	// 用户确认收货时间
	ConfirmReceiptTime int64 `json:"confirm_receipt_time"`
	// 【商品类型】 0、实体 1、虚拟
	GoodsType int64 `json:"goods_type"`
	// 商品ID
	ProductId int64 `json:"product_id"`
	// 商品skuId
	SkuId int64 `json:"sku_id"`
	// 规格信息
	Spec []SpecItem `json:"spec"`
	// 一级类目
	FirstCid int64 `json:"first_cid"`
	// 二级类目
	SecondCid int64 `json:"second_cid"`
	// 三级类目
	ThirdCid int64 `json:"third_cid"`
	// 四级类目
	FourthCid int64 `json:"fourth_cid"`
	// 外部Skuid
	OutSkuId string `json:"out_sku_id"`
	// sku外部供应商编码
	SupplierId string `json:"supplier_id"`
	// 商品外部编码
	OutProductId string `json:"out_product_id"`
	// 仓id，废弃，使用inventory_list
	WarehouseIds []string `json:"warehouse_ids"`
	// 外部仓id，废弃，使用inventory_list
	OutWarehouseIds []string `json:"out_warehouse_ids"`
	// 库存类型，普通库存/区域库存，废弃，使用inventory_list
	InventoryType string `json:"inventory_type"`
	// 库存类型描述，废弃，使用inventory_list
	InventoryTypeDesc string `json:"inventory_type_desc"`
	// 【库存扣减方式】 1、下单减库存 2、支付减库存
	ReduceStockType int64 `json:"reduce_stock_type"`
	// 库存扣减方式名称
	ReduceStockTypeDesc string `json:"reduce_stock_type_desc"`
	// 商品现价（单位：分）
	OriginAmount int64 `json:"origin_amount"`
	// 是否包税
	HasTax bool `json:"has_tax"`
	// 订单商品数量
	ItemNum int64 `json:"item_num"`
	// 商品现价*件数
	SumAmount int64 `json:"sum_amount"`
	// 商品来源平台
	SourcePlatform string `json:"source_platform"`
	// 商品图片
	ProductPic string `json:"product_pic"`
	// 是否评价 :1已评价，0未评价，2表示追评
	IsComment int64 `json:"is_comment"`
	// 商品名称
	ProductName string `json:"product_name"`
	// 仓库信息
	InventoryList []InventoryListItem `json:"inventory_list"`
	// 运费（单位：分）
	PostAmount int64 `json:"post_amount"`
	// 预售类型 ，0 现货类型，1 全款预售 2 阶梯发货
	PreSaleType int64 `json:"pre_sale_type"`
	// 售后信息
	AfterSaleInfo *AfterSaleInfo `json:"after_sale_info"`
	// 红包优惠金额（单位：分）
	PromotionRedpackAmount int64 `json:"promotion_redpack_amount"`
	// 平台红包优惠金额（单位：分），属于平台的红包的优惠金额
	PromotionRedpackPlatformAmount int64 `json:"promotion_redpack_platform_amount"`
	// 达人红包优惠金额（单位：分），属于达人的红包的优惠金额
	PromotionRedpackTalentAmount int64 `json:"promotion_redpack_talent_amount"`
	// 1:邮寄，2:自提
	ReceiveType int64 `json:"receive_type"`
	// 是否需要上传序列号，用于判断发货时是否需要上传商品序列号（IMEI码或SN码），true 表示是3C数码商品，需要上传序列号
	NeedSerialNumber bool `json:"need_serial_number"`
	// 广告来源，video：短视频 live：直播
	AdEnvType string `json:"ad_env_type"`
	// 商品单标签
	SkuOrderTagUi []SkuOrderTagUiItem `json:"sku_order_tag_ui"`
	// 商品ID，字符串型
	ProductIdStr string `json:"product_id_str"`
	// 预约发货时间
	AppointmentShipTime int64 `json:"appointment_ship_time"`
	// 直播房间ID，字符串类型
	RoomIdStr string `json:"room_id_str"`
	// 绑定类型 MASTER-主品单 FREE-免费赠品
	GivenProductType string `json:"given_product_type"`
	// 废弃，使用master_sku_order_id_list；绑定的主品sku单单号 如果given_product_type=FREE 则master_sku_order_id_list为主品sku单单号
	MasterSkuOrderId string `json:"master_sku_order_id"`
	// 商品卡券基本信息
	CardVoucher *CardVoucher `json:"card_voucher"`
	// 组套商品子商品列表
	BundleSkuInfo []BundleSkuInfoItem `json:"bundle_sku_info"`
	// 账号信息list
	AccountList *AccountList `json:"account_list"`
	// 定制商品信息
	SkuCustomizationInfo []SkuCustomizationInfoItem `json:"sku_customization_info"`
	// 作者（达人）承担金额（单位：分），订单参与活动和优惠中作者（达人）承担部分的总金额
	AuthorCostAmount int64 `json:"author_cost_amount"`
	// 仅平台承担金额（单位：分），订单参与活动和优惠中平台承担部分的总金额
	OnlyPlatformCostAmount int64 `json:"only_platform_cost_amount"`
	// 履约时效信息(json串)
	PromiseInfo string `json:"promise_info"`
	// 门店信息
	StoreInfo *StoreInfo `json:"store_info"`
	// 收件人姓名（脱敏后）
	MaskPostReceiver string `json:"mask_post_receiver"`
	// 收件人电话（脱敏后）
	MaskPostTel string `json:"mask_post_tel"`
	// 收件人地址（脱敏后）
	MaskPostAddr *MaskPostAddr `json:"mask_post_addr"`
	// 合约信息
	ContractInfo *ContractInfo `json:"contract_info"`
	// 绑定的主品sku单单号 如果given_product_type=FREE 则master_sku_order_id为主品sku单单号
	MasterSkuOrderIdList []string `json:"master_sku_order_id_list"`
	// 赠品活动信息
	GivenProductActivityInfo *GivenProductActivityInfo `json:"given_product_activity_info"`
	// 质检信息
	QualityCheckInfo *QualityCheckInfo `json:"quality_check_info"`
	// sku云仓择仓信息（云仓业务使用，非商品区域仓功能使用）
	SkuYunWarehouseInfo []SkuYunWarehouseInfoItem `json:"sku_yun_warehouse_info"`
	// 税费
	TaxAmount int64 `json:"tax_amount"`
}
type UserIdInfo struct {
	// 证件号
	IdCardNo string `json:"id_card_no"`
	EncryptIdCardNo string `json:"encrypt_id_card_no"`
	// 证件姓名
	IdCardName string `json:"id_card_name"`
	EncryptIdCardName string `json:"encrypt_id_card_name"`
}
type InventoryListItem struct {
	// 仓id
	WarehouseId string `json:"warehouse_id"`
	// 外部仓id
	OutWarehouseId string `json:"out_warehouse_id"`
	// 库存类型，普通库存/区域库存
	InventoryType int64 `json:"inventory_type"`
	// 库存类型描述
	InventoryTypeDesc string `json:"inventory_type_desc"`
	// 库存变更数量
	Count int64 `json:"count"`
	// 仓类型：0-商家仓，1-跨境仓，2-云仓
	WarehouseType int64 `json:"warehouse_type"`
}
type BundleSkuInfoItem struct {
	// 组套商品子商品ID
	ProductId string `json:"product_id"`
	// 组套商品子商品SKU_ID
	SkuId string `json:"sku_id"`
	// 组套商品子商品名称
	ProductName string `json:"product_name"`
	// 组套商品子商品数量
	ItemNum int64 `json:"item_num"`
	// 组商品子商品图片
	PictureUrl string `json:"picture_url"`
	// 组套子商品外部编码
	Code string `json:"code"`
}
type StoreAddress struct {
	// 省
	Province *Province `json:"province"`
	// 市
	City *City `json:"city"`
	// 县/区
	Town *Town `json:"town"`
	// 街道
	Street *Street `json:"street"`
	// 详细信息
	Detail string `json:"detail"`
}
type OrderPhaseListItem struct {
	// 阶段单id
	PhaseOrderId string `json:"phase_order_id"`
	// 总共有几阶段
	TotalPhase int64 `json:"total_phase"`
	// 第几阶段
	CurrentPhase int64 `json:"current_phase"`
	// 是否支付成功
	PaySuccess bool `json:"pay_success"`
	// 商品单ID
	SkuOrderId string `json:"sku_order_id"`
	// 活动id
	CampaignId int64 `json:"campaign_id"`
	// 阶段价格：定金单价、尾款单价，单位：分
	PhasePayablePrice int64 `json:"phase_payable_price"`
	// 支付类型:0-货到付款, 1-微信, 2-支付宝, 3-小程序, 4-银行卡, 5-余额, 7-无需支付, 8-放心花(信用支付), 9-新卡支付
	PhasePayType int64 `json:"phase_pay_type"`
	// 阶段单开始开启支付时间，时间戳
	PhaseOpenTime int64 `json:"phase_open_time"`
	// 阶段单支付成功时间，时间戳
	PhasePayTime int64 `json:"phase_pay_time"`
	// 阶段单关闭时间，时间戳
	PhaseCloseTime int64 `json:"phase_close_time"`
	// 渠道支付订单号
	ChannelPaymentNo string `json:"channel_payment_no"`
	// 阶段单总金额，单位分，为订单金额phase_sum_amount+运费phase_post_amount,为支付金额phase_pay_amount+优惠金额phase_promotion_amount
	PhaseOrderAmount int64 `json:"phase_order_amount"`
	// 阶段单订单金额，单位分
	PhaseSumAmount int64 `json:"phase_sum_amount"`
	// 阶段单运费金额，单位分
	PhasePostAmount int64 `json:"phase_post_amount"`
	// 阶段单用户实际支付金额，单位分
	PhasePayAmount int64 `json:"phase_pay_amount"`
	// 阶段单总优惠金额，单位分
	PhasePromotionAmount int64 `json:"phase_promotion_amount"`
	// 阶段状态描述
	CurrentPhaseStatusDesc string `json:"current_phase_status_desc"`
	// sku单价，单位：分
	SkuPrice int64 `json:"sku_price"`
}
type ShopOrderListItem struct {
	// 小时达订单接单状态；0-未接单；1-已接单；2-超时取消，或商家取消
	AcceptOrderStatus int64 `json:"accept_order_status"`
	// 店铺ID
	ShopId int64 `json:"shop_id"`
	// 商户名称
	ShopName string `json:"shop_name"`
	// 抖音小程序ID
	OpenId string `json:"open_id"`
	// 店铺订单号
	OrderId string `json:"order_id"`
	// 订单层级
	OrderLevel int64 `json:"order_level"`
	// 【业务来源】 1、鲁班 2、小店 3、好好学习 4、ev 5、虚拟 6、建站 7、核销 8、玉石 9、ez 10、ep 11、虚拟卡券 12、服务市场 13、EP视频课 14、EP直播课 21、跨境BBC 22、跨境BC 23、跨境CC,UPC 24、手机充值 25、拍卖保证金 26、懂车帝抵扣券 27、懂车帝返现券 28、离岛免税 29、海南会员购 30、抽奖 32、dou+券 76、大闸蟹 99、保险 102、小店海外
	Biz int64 `json:"biz"`
	// 业务来源描述
	BizDesc string `json:"biz_desc"`
	// 【订单类型】 0、普通订单 2、虚拟商品订单 4、电子券（poi核销） 5、三方核销
	OrderType int64 `json:"order_type"`
	// 订单类型描述
	OrderTypeDesc string `json:"order_type_desc"`
	// 【已废弃】请使用sku_order_list.trade_type
	TradeType int64 `json:"trade_type"`
	// 【已废弃】请使用sku_order_list. trade_type_desc
	TradeTypeDesc string `json:"trade_type_desc"`
	// 订单状态;1 待确认/待支付-订单创建完毕;105-已支付; 2-备货中； 101-部分发货; 3-已发货（全部发货）;4- 已取消;5 已完成（已收货）;
	OrderStatus int64 `json:"order_status"`
	// 订单状态描述
	OrderStatusDesc string `json:"order_status_desc"`
	// 主流程状态
	MainStatus int64 `json:"main_status"`
	// 主流程状态描述
	MainStatusDesc string `json:"main_status_desc"`
	// 支付时间
	PayTime int64 `json:"pay_time"`
	// 订单过期时间
	OrderExpireTime int64 `json:"order_expire_time"`
	// 订单完成时间
	FinishTime int64 `json:"finish_time"`
	// 下单时间
	CreateTime int64 `json:"create_time"`
	// 订单更新时间
	UpdateTime int64 `json:"update_time"`
	// 取消原因
	CancelReason string `json:"cancel_reason"`
	// 买家留言
	BuyerWords string `json:"buyer_words"`
	// 商家备注
	SellerWords string `json:"seller_words"`
	// 【下单端】 0、站外 1、火山 2、抖音 3、头条 4、西瓜 5、微信 6、值点app 7、头条lite 8、懂车帝 9、皮皮虾 11、抖音极速版 12、TikTok 13、musically 14、穿山甲 15、火山极速版 16、服务市场 26、番茄小说 27、UG教育营销电商平台 28、Jumanji 29、电商SDK
	BType int64 `json:"b_type"`
	// 下单端描述
	BTypeDesc string `json:"b_type_desc"`
	// 【下单场景】 0、未知 1、app内-原生 2、app内-小程序 3、H5 13、电商SDK-头条 35、电商SDK-头条lite
	SubBType int64 `json:"sub_b_type"`
	// 下单场景描述
	SubBTypeDesc string `json:"sub_b_type_desc"`
	// 具体某个小程序的ID
	AppId int64 `json:"app_id"`
	// 【支付类型】 0、货到付款 1 、微信 2、支付宝 3、小程序 4、银行卡 5、余额 7、无需支付（0元单） 8、DOU分期（信用支付） 9、新卡支付 12、先用后付
	PayType int64 `json:"pay_type"`
	// 支付渠道的流水号
	ChannelPaymentNo string `json:"channel_payment_no"`
	// 订单金额（单位：分）
	OrderAmount int64 `json:"order_amount"`
	// 支付金额（单位：分）
	PayAmount int64 `json:"pay_amount"`
	// 快递费（单位：分）
	PostAmount int64 `json:"post_amount"`
	// 运费险金额（单位：分）
	PostInsuranceAmount int64 `json:"post_insurance_amount"`
	// 改价金额变化量（单位：分）
	ModifyAmount int64 `json:"modify_amount"`
	// 改价运费金额变化量（单位：分）
	ModifyPostAmount int64 `json:"modify_post_amount"`
	// 订单优惠总金额（单位：分）= 店铺优惠金额+ 平台优惠金额+ 达人优惠金额
	PromotionAmount int64 `json:"promotion_amount"`
	// 店铺优惠金额（单位：分），属于店铺的优惠活动、优惠券、红包的总优惠金额
	PromotionShopAmount int64 `json:"promotion_shop_amount"`
	// 平台优惠金额（单位：分），属于平台的优惠活动、优惠券、红包的总优惠金额
	PromotionPlatformAmount int64 `json:"promotion_platform_amount"`
	// 商家承担金额（单位：分），订单参与活动和优惠中商家承担部分的总金额
	ShopCostAmount int64 `json:"shop_cost_amount"`
	// 平台承担金额（单位：分），订单参与活动和优惠中平台+作者（达人）承担部分的总金额,包含作者（达人）承担金额：platform_cost_amount = only_platform_cost_amount + author_cost_amount
	PlatformCostAmount int64 `json:"platform_cost_amount"`
	// 达人优惠金额（单位：分），属于达人的优惠活动、优惠券、红包的总优惠金额
	PromotionTalentAmount int64 `json:"promotion_talent_amount"`
	// 支付优惠金额（单位：分），支付渠道上的优惠金额
	PromotionPayAmount int64 `json:"promotion_pay_amount"`
	// 收件人电话
	PostTel string `json:"post_tel"`
	EncryptPostTel string `json:"encrypt_post_tel"`
	// 收件人姓名
	PostReceiver string `json:"post_receiver"`
	EncryptPostReceiver string `json:"encrypt_post_receiver"`
	// 收件人地址
	PostAddr *PostAddr `json:"post_addr"`
	// 预计发货时间
	ExpShipTime int64 `json:"exp_ship_time"`
	// 发货时间
	ShipTime int64 `json:"ship_time"`
	// 物流信息
	LogisticsInfo []LogisticsInfoItem `json:"logistics_info"`
	// 商品单信息
	SkuOrderList []SkuOrderListItem `json:"sku_order_list"`
	// 卖家订单标记 小旗子star取值0～5，分别表示 灰紫青绿橙红
	SellerRemarkStars int64 `json:"seller_remark_stars"`
	// 定金预售阶段单
	OrderPhaseList []OrderPhaseListItem `json:"order_phase_list"`
	// 用户加密串
	DoudianOpenId string `json:"doudian_open_id"`
	// 商品序列号（IMEI码或SN码）
	SerialNumberList []string `json:"serial_number_list"`
	// 红包优惠金额（单位：分）
	PromotionRedpackAmount int64 `json:"promotion_redpack_amount"`
	// 平台红包优惠金额（单位：分），属于平台的红包的优惠金额
	PromotionRedpackPlatformAmount int64 `json:"promotion_redpack_platform_amount"`
	// 达人红包优惠金额（单位：分），属于达人的红包的优惠金额
	PromotionRedpackTalentAmount int64 `json:"promotion_redpack_talent_amount"`
	// 用户证件信息
	UserIdInfo *UserIdInfo `json:"user_id_info"`
	// 预约发货时间
	AppointmentShipTime int64 `json:"appointment_ship_time"`
	// 懂车帝购车信息
	DCarShopBizData *DCarShopBizData `json:"d_car_shop_biz_data"`
	// 店铺单标签
	ShopOrderTagUi []ShopOrderTagUiItem `json:"shop_order_tag_ui"`
	// 总优惠金额（单位：分），total_promotion_amount = promotion_amount + post_promotion_amount
	TotalPromotionAmount int64 `json:"total_promotion_amount"`
	// 运费原价（单位：分），post_origin_amount = post_amount + post_promotion_amount
	PostOriginAmount int64 `json:"post_origin_amount"`
	// 运费优惠金额（单位：分）
	PostPromotionAmount int64 `json:"post_promotion_amount"`
	// 用户特征标签
	UserTagUi []UserTagUiItem `json:"user_tag_ui"`
	// 作者（达人）承担金额（单位：分），订单参与活动和优惠中作者（达人）承担部分的总金额
	AuthorCostAmount int64 `json:"author_cost_amount"`
	// 仅平台承担金额（单位：分），订单参与活动和优惠中平台承担部分的总金额
	OnlyPlatformCostAmount int64 `json:"only_platform_cost_amount"`
	// 履约时效信息(json串)
	PromiseInfo string `json:"promise_info"`
	// 收件人姓名（脱敏后）
	MaskPostReceiver string `json:"mask_post_receiver"`
	// 收件人电话（脱敏后）
	MaskPostTel string `json:"mask_post_tel"`
	// 收件人地址（脱敏后）
	MaskPostAddr *MaskPostAddr `json:"mask_post_addr"`
	// 买家收货地址经纬度信息，高德坐标系；
	UserCoordinate *UserCoordinate `json:"user_coordinate"`
	// 预计最早送达时间，Unix时间戳：秒；仅小时达业务返回
	EarliestReceiptTime int64 `json:"earliest_receipt_time"`
	// 预计最晚送达时间，Unix时间戳：秒；仅小时达业务返回
	LatestReceiptTime int64 `json:"latest_receipt_time"`
	// 是否尽快送达，true-是，false-否，仅小时达业务返回
	EarlyArrival bool `json:"early_arrival"`
	// 尽快送达的时间点，Unix时间戳：秒；仅小时达业务返回
	TargetArrivalTime int64 `json:"target_arrival_time"`
	// 打包费，单位：分
	PackingAmount int64 `json:"packing_amount"`
	// 门店流水号；仅小时达业务返回
	SupermarketOrderSerialNo string `json:"supermarket_order_serial_no"`
	// 税费（子单税费之和）
	TaxAmount int64 `json:"tax_amount"`
	// 地址标签列表
	AddressTagUi []AddressTagUiItem `json:"address_tag_ui"`
}
type PostAddr struct {
	// 省/直辖市
	Province *Province `json:"province"`
	// 市
	City *City `json:"city"`
	// 区县
	Town *Town `json:"town"`
	// 街道
	Street *Street `json:"street"`
	// 详细地址
	Detail string `json:"detail"`
	EncryptDetail string `json:"encrypt_detail"`
}
type PicItem struct {
	// id
	Id int64 `json:"id"`
	// url
	Url string `json:"url"`
}
type AfterSaleInfo struct {
	// 售后状态，0-售后初始化， 6-售后申请， 7-售后退货中， 27-拒绝售后申请， 12-售后成功， 28-售后失败， 11-售后已发货， 29-退货后拒绝退款， 13-售后换货商家发货， 14-售后换货用户收货， 51-取消成功， 53-逆向交易完成
	AfterSaleStatus int64 `json:"after_sale_status"`
	// 售后类型:0 售后退货退款:1-售后退款 2-售前退款 3-换货 4-系统取消 5-用户取消
	AfterSaleType int64 `json:"after_sale_type"`
	// 退款状态:1-待退款；3-退款成功； 4-退款失败；当买家发起售后后又主动取消售后，此时after_sale_status=28并且refund_status=1的状态不变，不会流转至4状态；
	RefundStatus int64 `json:"refund_status"`
}
type AccountList struct {
	// 账号信息
	AccountInfo []AccountInfoItem `json:"account_info"`
}
type OrderSearchListData struct {
	// 页数，从0开始
	Page int64 `json:"page"`
	// 总订单数
	Total int64 `json:"total"`
	// 单页大小
	Size int64 `json:"size"`
	// 订单信息
	ShopOrderList []ShopOrderListItem `json:"shop_order_list"`
}
type Province struct {
	// 名称
	Name string `json:"name"`
	// 地区ID
	Id string `json:"id"`
}
type City struct {
	// 名称
	Name string `json:"name"`
	// 地区ID
	Id string `json:"id"`
}
type UserTagUiItem struct {
	// 标签key
	Key string `json:"key"`
	// 标签名称
	Text string `json:"text"`
}
type UserCoordinate struct {
	// 买家收货地址经度信息，高德坐标系；
	UserCoordinateLongitude string `json:"user_coordinate_longitude"`
	// 买家收货地址纬度信息，高德坐标系；
	UserCoordinateLatitude string `json:"user_coordinate_latitude"`
}
type SkuYunWarehouseInfoItem struct {
	// 预分配云仓信息
	PreAllocatedYunWarehouseInfo *PreAllocatedYunWarehouseInfo `json:"pre_allocated_yun_warehouse_info"`
	// 云仓实仓信息
	RealYunWarehouseInfo *RealYunWarehouseInfo `json:"real_yun_warehouse_info"`
	// 当前业务阶段 1-未分配 2-已分配未发货 2-已发货
	CurrentBusinessStage int64 `json:"current_business_stage"`
	// 该仓对应的skuId
	SkuId string `json:"sku_id"`
	// sku数量
	SkuCount int64 `json:"sku_count"`
}
type DCarShopBizData struct {
	// 选择的门店ID
	PoiId string `json:"poi_id"`
	// 选择的门店名称
	PoiName string `json:"poi_name"`
	// 选择的门店地址
	PoiAddr string `json:"poi_addr"`
	// 选择的门店电话
	PoiTel string `json:"poi_tel"`
	// 权益信息
	CouponRight []CouponRightItem `json:"coupon_right"`
	// 选择的门店所在省
	PoiPname string `json:"poi_pname"`
	// 选择的门店所在市
	PoiCityName string `json:"poi_city_name"`
	// 选择的门店所在区县
	PoiAdname string `json:"poi_adname"`
}
type SkuOrderTagUiItem struct {
	// 标签key
	Key string `json:"key"`
	// 标签文案
	Text string `json:"text"`
	// 标签备注文案
	HoverText string `json:"hover_text"`
	// 标签类型，如颜色
	TagType string `json:"tag_type"`
	// 帮助文档
	HelpDoc string `json:"help_doc"`
	// 排序
	Sort int64 `json:"sort"`
	// 标签其他信息
	Extra string `json:"extra"`
}
type Detail struct {
	// 定制图片信息
	Pic []PicItem `json:"pic"`
	// 定制文案信息
	Text []TextItem `json:"text"`
	// 额外信息
	Extra string `json:"extra"`
}
type LogisticsInfoItem struct {
	// 物流单号
	TrackingNo string `json:"tracking_no"`
	// 物流公司
	Company string `json:"company"`
	// 发货时间
	ShipTime int64 `json:"ship_time"`
	// 包裹id
	DeliveryId string `json:"delivery_id"`
	// 物流公司名称
	CompanyName string `json:"company_name"`
	// 商品信息
	ProductInfo []ProductInfoItem `json:"product_info"`
	// 骑手取件码；触发场景①商家：抖超小时达商家通过平台运力发货并选择需要取件码时，该字段有值返回。应用场景①骑手到店后报流水号，门店员工根据流水号【示例值：流水号#1】找到对应包裹。②门店员工与骑手核对取件码，核对无误后交货给骑手；
	HourUpPickupCode string `json:"hour_up_pickup_code"`
}
type AccountInfoItem struct {
	// 账号名称
	AccountName string `json:"account_name"`
	// 账号类型；枚举值：Mobile ：手机号;Email ：邮箱 ; IdCard ：身份证; Passport ：护照; BankCard ：银行卡; Number ：纯数学; NumberLetter ：数字字母混合;
	AccountType string `json:"account_type"`
	// 账号值
	AccountId string `json:"account_id"`
	EncryptAccountId string `json:"encrypt_account_id"`
}
type TextItem struct {
	// id
	Id int64 `json:"id"`
	// key
	Key string `json:"key"`
	// content
	Content string `json:"content"`
}
type StoreInfo struct {
	// 门店id
	StoreId string `json:"store_id"`
	// 门店名称
	StoreName string `json:"store_name"`
	// 门店电话
	StoreTel string `json:"store_tel"`
	// 门店地址
	StoreAddress *StoreAddress `json:"store_address"`
	// 扩展字段
	Extra string `json:"extra"`
}
type ContractInfo struct {
	// 办理合约的手机号
	MobileNo string `json:"mobile_no"`
	EncryptMobileNo string `json:"encrypt_mobile_no"`
}
type QualityCheckInfo struct {
	// 质检结果 WATI_CHECT:待质检, CHECK_PASS:质检通过, CHECK_FAIL:质检不通过
	CheckResultCode string `json:"check_result_code"`
	// 质检异常信息描述文案
	CheckFailMsg string `json:"check_fail_msg"`
	// 重新送检截止时间
	ResendCheckTime int64 `json:"resend_check_time"`
}
type Street struct {
	// 名称
	Name string `json:"name"`
	// 地区ID
	Id string `json:"id"`
}
type SkuSpecsItem struct {
	// 规格名称
	Name string `json:"name"`
	// 规格值
	Value string `json:"value"`
}
type GivenProductActivityInfo struct {
	// NORMAL-普通的买赠，AMOUNT-满件赠，PRICE-满元赠
	GivenProductActivityType string `json:"given_product_activity_type"`
	// 满几件赠或满几元赠，满几元赠场景单位是分
	GivenProductActivityValue string `json:"given_product_activity_value"`
}
type CouponRightItem struct {
	// 权益类型
	RightType int64 `json:"right_type"`
	// 权益名称
	RightName string `json:"right_name"`
	// 权益面值
	Quota int64 `json:"quota"`
}
type ProductInfoItem struct {
	// 商品名称
	ProductName string `json:"product_name"`
	// 商品价格
	Price int64 `json:"price"`
	// 商家编码
	OuterSkuId string `json:"outer_sku_id"`
	// 商品skuId
	SkuId int64 `json:"sku_id"`
	// 规格信息
	SkuSpecs []SkuSpecsItem `json:"sku_specs"`
	// 发货商品数量
	ProductCount int64 `json:"product_count"`
	// 商品ID
	ProductId int64 `json:"product_id"`
	// 商品单ID
	SkuOrderId string `json:"sku_order_id"`
	// 商品ID，字符串型
	ProductIdStr string `json:"product_id_str"`
}
type CardVoucher struct {
	// 自领取之日起有效天数(如果有值优先使用该字段)
	ValidDays int64 `json:"valid_days"`
	// 卡券开始时间
	ValidStart int64 `json:"valid_start"`
	// 卡券截止日期
	ValidEnd int64 `json:"valid_end"`
}
