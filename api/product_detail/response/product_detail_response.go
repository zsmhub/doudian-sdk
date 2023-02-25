package product_detail_response

import (
	"doudian.com/open/sdk_golang/core"
)

type ProductDetailResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *ProductDetailData `json:"data"`
}
type UnitPriceInfo struct {
	// 1：投资金单位价格2：首饰金单位价格
	PriceRuleType int32 `json:"price_rule_type"`
}
type SpecPicsItem struct {
	// id
	SpecDetailId int64 `json:"spec_detail_id"`
	// 规格图片uri
	Pic string `json:"pic"`
}
type DeliveryInfosItem struct {
	// 信息类型
	InfoType string `json:"info_type"`
	// 信息值
	InfoValue string `json:"info_value"`
	// 信息计量单位
	InfoUnit string `json:"info_unit"`
}
type SpecsItem struct {
	// 【无需使用】销售属性id
	Id int64 `json:"id"`
	// 【无需使用】规格id
	SpecId int64 `json:"spec_id"`
	// 销售属性
	Name string `json:"name"`
	// 【无需使用】销售属性id
	Pid int64 `json:"pid"`
	// 是否是销售属性值
	IsLeaf int16 `json:"is_leaf"`
	// 销售属性值
	Values []ValuesItem `json:"values"`
}
type QualityAttachmentsItem struct {
	// 1-图片，2-pdf
	MediaType int64 `json:"media_type"`
	// 资质url
	Url string `json:"url"`
}
type QualityListItem struct {
	// 资质key
	QualityKey string `json:"quality_key"`
	// 资质名称
	QualityName string `json:"quality_name"`
	// 资质链接
	QualityAttachments []QualityAttachmentsItem `json:"quality_attachments"`
}
type PoiResource struct {
	// 1 随时退+过期自动退，2 不支持退
	CouponReturnMethods []int64 `json:"coupon_return_methods"`
}
type DelayRule struct {
	// 是否开启特殊日期延迟发货
	Enable bool `json:"enable"`
	// 1 时间点；2 时间范围
	ConfigType int32 `json:"config_type"`
	// 特殊日期延迟发货时间最晚发货时间，时间戳，单位秒
	ConfigValue int64 `json:"config_value"`
	// 特殊日期延迟发货时间下单开始时间，时间戳，单位秒
	StartTime int64 `json:"start_time"`
	// 特殊日期延迟发货时间下单结束时间，时间戳，单位秒
	EndTime int64 `json:"end_time"`
}
type CategoryDetail struct {
	// 一级类目id
	FirstCid int64 `json:"first_cid"`
	// 二级类目id
	SecondCid int64 `json:"second_cid"`
	// 三级类目id
	ThirdCid int64 `json:"third_cid"`
	// 四级类目id
	FourthCid int64 `json:"fourth_cid"`
	// 一级类目名
	FirstCname string `json:"first_cname"`
	// 二级类目名
	SecondCname string `json:"second_cname"`
	// 三级类目名
	ThirdCname string `json:"third_cname"`
	// 四级类目名
	FourthCname string `json:"fourth_cname"`
}
type TaxExemptionSkuInfo struct {
	// 是否套装（0 否 1 是）
	IsSuit int32 `json:"is_suit"`
	// 套装数量
	SuitNum int64 `json:"suit_num"`
	// 净含量
	Volume int64 `json:"volume"`
}
type SpecPricesItem struct {
	// Sku的黄金加工费，单位：分
	GoldProcessCharge int64 `json:"gold_process_charge"`
	// 规格id；抖店系统生成，商品id下唯一。
	SkuId int64 `json:"sku_id"`
	// 外部商家skui_id编码，商家自定义字段；推荐使用outer_sku_id字段
	OutSkuId int64 `json:"out_sku_id"`
	// 外部商家skui_id编码，商家自定义字段
	OuterSkuId string `json:"outer_sku_id"`
	// 规格id列表,多规格以”,“分隔；
	SpecDetailIds []int64 `json:"spec_detail_ids"`
	// 可售库存；当前现货可售库存；
	StockNum int64 `json:"stock_num"`
	// 商品价格；单位：分
	Price int64 `json:"price"`
	// 编码
	Code string `json:"code"`
	// 阶梯库存，规则详见名称解释：https://op.jinritemai.com/docs/guide-docs/202/170
	StepStockNum int64 `json:"step_stock_num"`
	// 活动库存，，规则详见名称解释：https://op.jinritemai.com/docs/guide-docs/202/170
	PromStockNum int64 `json:"prom_stock_num"`
	// 活动阶梯库存，，规则详见名称解释：https://op.jinritemai.com/docs/guide-docs/202/170
	PromStepStockNum int64 `json:"prom_step_stock_num"`
	// 【已废弃，无需使用】规格ID
	SpecDetailId1 int64 `json:"spec_detail_id1"`
	// 【已废弃，无需使用】规格ID
	SpecDetailId2 int64 `json:"spec_detail_id2"`
	// 【已废弃，无需使用】规格ID
	SpecDetailId3 int64 `json:"spec_detail_id3"`
	// sku类型；0-普通库存  1-区域库存  10-阶梯库存
	SkuType int64 `json:"sku_type"`
	// 供应商编码
	SupplierId string `json:"supplier_id"`
	// 活动现货库存
	PromotionStockNum int64 `json:"promotion_stock_num"`
	// 活动阶梯库存
	PromotionStepStockNum int64 `json:"promotion_step_stock_num"`
	// 海关申报要素（仅海淘商品返回信息）
	CustomsReportInfo *CustomsReportInfo `json:"customs_report_info"`
	// 现货订单锁定库存
	LockStockNum int64 `json:"lock_stock_num"`
	// 阶梯订单锁定库存
	LockStepStockNum int64 `json:"lock_step_stock_num"`
	// 仓ID->库存映射
	StockNumMap map[string]int64 `json:"stock_num_map"`
	// 海南免税sku信息（仅海淘商品返回信息）
	TaxExemptionSkuInfo *TaxExemptionSkuInfo `json:"tax_exemption_sku_info"`
	// 发货时间（全款预售模式时的发货时间/阶梯模式下阶梯发货时间），9999是当日发、1次日发、2 是48小时发、5、15、45天等。发货时间规则可使用【product/getProductUpdateRule】获取
	PresellDelay int64 `json:"presell_delay"`
	// SKU物流信息
	DeliveryInfos []DeliveryInfosItem `json:"delivery_infos"`
}
type QualityInspectionInfo struct {
	// 1: 单库存模式，只允许售卖一个sku；2: 多库存模式，不限售卖次数
	Mode int32 `json:"mode"`
	// 质检证书编码
	CertificateCode string `json:"certificate_code"`
	// 机构编码
	Agency string `json:"agency"`
	// 是否勾选支持前置质检
	Supported bool `json:"supported"`
}
type ProductDetailData struct {
	// 商品ID，抖店系统生成，店铺下唯一；长度19位。
	ProductId int64 `json:"product_id"`
	// 商品ID（字符串类型），抖店系统生成，店铺下唯一；长度19位。
	ProductIdStr string `json:"product_id_str"`
	// 【即将废弃】外部商家编码，商家自定义字段。推荐使用，outer_product_id字段
	OutProductId int64 `json:"out_product_id"`
	// 外部商家编码，商家自定义字段，支持最多 255个字符
	OuterProductId string `json:"outer_product_id"`
	// 【已废弃】open应用id
	OpenUserId int64 `json:"open_user_id"`
	// 商品标题，规则：至少输入8个字（16个字符）以上~输入30个字（60个字符）以内。；标题不规范会引起商品下架，影响您的正常销售，详见商品发布规范：https://school.jinritemai.com/doudian/web/article/101800?from=shop_article
	Name string `json:"name"`
	// 商品详情，最大支持50张图片，单张详情图宽高比不超2000*2000px，大小5M内，仅支持jpg/jpeg/png格式，返回HTML格式；注意：商品详情不规范，会引起商品下架，影响您的正常销售：https://school.jinritemai.com/doudian/web/article/101800?from=shop_article
	Description string `json:"description"`
	// 承诺发货时间，单位是天; 不传则默认为2天，当presell_type为0或2均只支持传入2；当presell_type为1时支持可选值为: 2、3、5、7、10、15；
	DeliveryMethod int32 `json:"delivery_method"`
	// 海南免税，海关限购分类编码，仅海淘商品有值返回。
	CdfCategory string `json:"cdf_category"`
	// 商品在店铺中状态: 0-在线；1-下线；2-删除；详见商品状态机：https://op.jinritemai.com/docs/question-docs/92/2070
	Status int64 `json:"status"`
	// 商品规格，新增商品是全局唯一，注意：有部分历史存量商品可能存在规格复用
	SpecId int64 `json:"spec_id"`
	// 商品审核状态: 1-未提交；2-待审核；3-审核通过；4-审核未通过；5-封禁；7-审核通过待上架；详见商品状态机：https://op.jinritemai.com/docs/question-docs/92/2070
	CheckStatus int64 `json:"check_status"`
	// 【已废弃】手机号
	Mobile string `json:"mobile"`
	// 【已废弃】品牌id，推荐使用standard_brand_id字段，通过【/brand/list】获取，无品牌id商品返回默认值：596120136；
	BrandId int64 `json:"brand_id"`
	// 是否是组合商品的子商品；true-是，false-不是；
	IsSubProduct bool `json:"is_sub_product"`
	// 草稿状态；0-无草稿,1-未提审,2-待审核,3-审核通过,4-审核未通过。详见：https://op.jinritemai.com/docs/question-docs/92/2070
	DraftStatus int64 `json:"draft_status"`
	// 类目详情；商品类目id可使用【/shop/getShopCategory】查询
	CategoryDetail *CategoryDetail `json:"category_detail"`
	// 【已废弃】支持的支付方式：0货到付款 1在线支付 2两者都支持
	PayType int64 `json:"pay_type"`
	// 【已废弃】商家推荐语
	RecommendRemark string `json:"recommend_remark"`
	// 额外信息，如资质
	Extra string `json:"extra"`
	// 【已废弃】无业务意义
	IsCreate int64 `json:"is_create"`
	// 商品创建时间，时间格式：yyyy-mm-dd hh:mm:ss
	CreateTime string `json:"create_time"`
	// 商品更新时间，时间格式：yyyy-mm-dd hh:mm:ss
	UpdateTime string `json:"update_time"`
	// 商品主图；最多支持5张图片；仅支持png，jpg，jpeg格式，宽高比例为1:1（至少600*600px），大小5M内
	Pic []string `json:"pic"`
	// 【即将废弃】推荐使用product_format_new；属性名称|属性值 之间用|分隔, 多组之间用^分开
	ProductFormat string `json:"product_format"`
	// 规格图片
	SpecPics []SpecPicsItem `json:"spec_pics"`
	// 商品sku详情
	SpecPrices []SpecPricesItem `json:"spec_prices"`
	// 销售属性
	Specs []SpecsItem `json:"specs"`
	// 头图，主图第一张
	Img string `json:"img"`
	// 预售类型，1-全款预售，0-非预售，2-阶梯库存
	PresellType int64 `json:"presell_type"`
	// 单用户下单限购件数
	MaximumPerOrder int64 `json:"maximum_per_order"`
	// 单用户累计限购件数
	LimitPerBuyer int64 `json:"limit_per_buyer"`
	// 用户每次下单至少购买的件数
	MinimumPerOrder int64 `json:"minimum_per_order"`
	// 资质信息
	QualityList []QualityListItem `json:"quality_list"`
	// 跨境物流信息（仅海淘商品返回）
	LogisticsInfo *LogisticsInfo `json:"logistics_info"`
	// 售后服务
	AfterSaleService string `json:"after_sale_service"`
	// 商品价格是否含税
	PriceHasTax int64 `json:"price_has_tax"`
	// 可预约发货天数
	AppointDeliveryDay int64 `json:"appoint_delivery_day"`
	// 类目属性
	ProductFormatNew string `json:"product_format_new"`
	// 品牌库brand id，原brand_id代表商标关系id
	StandardBrandId int64 `json:"standard_brand_id"`
	// 划线价 单位分
	MarketPrice int64 `json:"market_price"`
	// 售卖价 单位分
	DiscountPrice int64 `json:"discount_price"`
	// 汽车vin码
	CarVinCode string `json:"car_vin_code"`
	// 生活娱乐充值模式
	NeedRechargeMode bool `json:"need_recharge_mode"`
	// 多账号充值账号模板
	AccountTemplateId string `json:"account_template_id"`
	// 发货模式：presell_type = 0 现货；presell_type = 2 阶梯；presell_type = 1 && presell_config_level = 0 全款预售；presell_type = 1 && presell_config_level = 1 sku预售；presell_type = 1 && presell_config_level = 2 现货+预售；presell_type = 1 && presell_config_level = 3 新预售
	PresellConfigLevel int64 `json:"presell_config_level"`
	// 现货模式的发货天数；阶梯模式现货部分的发货天数，9999=当日发、1=次日发
	DeliveryDelayDay int64 `json:"delivery_delay_day"`
	// 阶梯模式阶梯部分的发货天数；商品全款预售模式，预售天数
	PresellDelay int64 `json:"presell_delay"`
	// 卡券信息
	PoiResource *PoiResource `json:"poi_resource"`
	// 特殊日期延迟发货规则
	DelayRule *DelayRule `json:"delay_rule"`
	// 3:4长图url(仅素材中心url有效)
	LongPicUrl string `json:"long_pic_url"`
	// 售卖方式;0:全渠道手售卖,1:仅指定直播间售卖
	SellChannel []int64 `json:"sell_channel"`
	// 运费模版ID
	FreightId int64 `json:"freight_id"`
	// 主图视频ID
	MaterialVideoId string `json:"material_video_id"`
	// 提取方式新字段，推荐使用。"0": 普通商品-使用物流发货, "1": 虚拟商品-无需物流与电子交易凭证, "2": 虚拟商品-使用电子交易凭证,  "3": 虚拟商品-充值直连
	PickupMethod string `json:"pickup_method"`
	// 尺码模板ID
	SizeInfoTemplateId int64 `json:"size_info_template_id"`
	// 白底图url(仅素材中心url有效)
	WhiteBackGroundPicUrl string `json:"white_back_ground_pic_url"`
	// 销售渠道类型，包括纯电商（onlineOnly）、专柜同款（sameAsOffline），云零售商家（https://fxg.jinritemai.com/ffa/merchant-growth/cloud-retail）可以设置
	SaleChannelType string `json:"sale_channel_type"`
	// 门店ID
	StoreId int64 `json:"store_id"`
	// 主商品ID
	MainProductId int64 `json:"main_product_id"`
	// 限售模板ID
	SaleLimitId int64 `json:"sale_limit_id"`
	// 系统推荐的标题前缀
	NamePrefix string `json:"name_prefix"`
	// 商品类型，0-普通，3-虚拟，6玉石闪购，7云闪购
	ProductType int64 `json:"product_type"`
	// 预售发货方式配置 0-预售结束后xx天发货; 1-支付完成后xx天发货
	PresellDeliveryType int64 `json:"presell_delivery_type"`
	// 审核通过后上架售卖时间配置：0-立即上架售卖 1-放入仓库
	StartSaleType int64 `json:"start_sale_type"`
	// 库存扣减方式，1-拍下减库存 2-付款减库存
	ReduceType int64 `json:"reduce_type"`
	// 预售结束时间，格式2020-02-21 18:54:27，最多支持设置距离当前30天
	PresellEndTime string `json:"presell_end_time"`
	// 重量数值
	WeightValue float64 `json:"weight_value"`
	// 重量单位，0-kg, 1-g
	WeightUnit int64 `json:"weight_unit"`
	// 参考价凭证
	ReferencePriceCertificate *ReferencePriceCertificate `json:"reference_price_certificate"`
	// 参考价，单位分
	ReferencePrice int64 `json:"reference_price"`
	// 前置质检相关（特定二手商家、特定二手类目使用）
	QualityInspectionInfo *QualityInspectionInfo `json:"quality_inspection_info"`
	// 单位价格信息
	UnitPriceInfo *UnitPriceInfo `json:"unit_price_info"`
}
type CustomsReportInfo struct {
	// 海关代码
	HsCode string `json:"hs_code"`
	// 法定第一计量数量
	FirstMeasureQty float64 `json:"first_measure_qty"`
	// 法定第二计量数量
	SecondMeasureQty float64 `json:"second_measure_qty"`
	// 法定第一计量单位
	FirstMeasureUnit string `json:"first_measure_unit"`
	// 法定第二计量单位
	SecondMeasureUnit string `json:"second_measure_unit"`
	// 售卖单位
	Unit string `json:"unit"`
	// 品名
	ReportName string `json:"report_name"`
	// 品牌
	ReportBrandName string `json:"report_brand_name"`
	// 用途
	Usage string `json:"usage"`
	// 规格型号
	GModel string `json:"g_model"`
	// 条形码
	BarCode string `json:"bar_code"`
}
type ValuesItem struct {
	// 销售属性值ID
	Id int64 `json:"id"`
	// 规格ID
	SpecId int64 `json:"spec_id"`
	// 销售属性值
	Name string `json:"name"`
	// 销售属性ID
	Pid int64 `json:"pid"`
	// 是否是销售属性值
	IsLeaf int16 `json:"is_leaf"`
	// 在线状态，-2彻底删除、0在线、1下线、2删除
	Status int16 `json:"status"`
}
type LogisticsInfo struct {
	// 通关模式，1BBC 2BC 3邮关
	CustomsClearType int64 `json:"customs_clear_type"`
	// 原产国id
	OriginCountryId int64 `json:"origin_country_id"`
	// 货源国id
	SourceCountryId int64 `json:"source_country_id"`
	// 品牌所在地id
	BrandCountryId int64 `json:"brand_country_id"`
	// 税金承担方，0商家承担，1用户承担
	TaxPayer int64 `json:"tax_payer"`
	// 商品净重
	NetWeightQty float64 `json:"net_weight_qty"`
}
type ReferencePriceCertificate struct {
	// 凭证类型;1:"厂商建议零售价",2:"吊牌价",3:"定价",4:"官网零售售价"
	CertificateType string `json:"certificate_type"`
	// 凭证图片url
	CertificateUrls string `json:"certificate_urls"`
}
