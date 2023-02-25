package product_getProductUpdateRule_response

import (
	"doudian.com/open/sdk_golang/core"
)

type ProductGetProductUpdateRuleResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *ProductGetProductUpdateRuleData `json:"data"`
}
type AfterSaleRule struct {
	// 无理由退货规则
	SupplyDayReturnRule *SupplyDayReturnRule `json:"supply_day_return_rule"`
}
type ReferencePriceRule struct {
	// true表示可填写参考价，false表示不可填写参考价
	IsSupport bool `json:"is_support"`
	// true表示必填参考价，false表示不必填参考价
	IsRequired bool `json:"is_required"`
	// 参考价凭证类型枚举值与对应的显示名称，比如{     "1" : "厂商建议零售价",     "2" : "吊牌价",     "3" : "定价",     "4" : "官网零售价" }
	CertificateTypes map[int64]string `json:"certificate_types"`
	// 参考价最高高于最低SKU价格的倍数，比如最低SKU价格为1，此值为10，则参考价不可高于10
	LowerSkuPriceTimes int64 `json:"lower_sku_price_times"`
}
type ProductSpecRule struct {
	// totally_use_given_spec：完全使用系统下发规格；partly_use_given_spec部分使用系统下发规格；not_use_given_spec：不管控规格;"":为空时表示不管控规格
	SpecRuleCode string `json:"spec_rule_code"`
	// 商品规格列表
	RequiredSpecDetails []RequiredSpecDetailsItem `json:"required_spec_details"`
	// 最大可支持的规格层级数量
	MaxSpecNumLimit int64 `json:"max_spec_num_limit"`
	// sku组合数量上限
	SpecCombinationLimit int64 `json:"spec_combination_limit"`
	// 单个规格的规格值数量上限
	SpecSingleLimit int64 `json:"spec_single_limit"`
}
type GoldPriceRule struct {
	// 价格规则信息，内有黄金的不同的单位价格以及枚举
	PriceRules []PriceRulesItem `json:"price_rules"`
	// 销售属性（规格）id，价格基于这个规格填写的数值按照公式进行计算
	SellPropertyId string `json:"sell_property_id"`
	// 是否必填，true表示必填，false表示非必填
	IsRequired bool `json:"is_required"`
	// 是否需要展示（可填），true表示需要展示（可填），false表示无需展示（不可填）
	IsShow bool `json:"is_show"`
}
type ProductPresellRule struct {
	// 是否支持全款预售
	Support bool `json:"support"`
	// 预售门槛价
	MinPresellPrice int64 `json:"min_presell_price"`
	// 是否支持支付成功后发货
	SupportDeliveryAfterPay bool `json:"support_delivery_after_pay"`
	// 是否支持预售结束后发货
	SupportDeliveryAfterPresell bool `json:"support_delivery_after_presell"`
	// 支付结束后规则
	DeliveryAfterPayConfig *DeliveryAfterPayConfig `json:"delivery_after_pay_config"`
	// 预售结束后规则
	DeliveryAfterPresellConfig *DeliveryAfterPresellConfig `json:"delivery_after_presell_config"`
}
type TimeSkuSpecDetialItem_4_4 struct {
	// 规格值
	SpecValue string `json:"spec_value"`
	// 是否预售时效
	IsPresellSpec bool `json:"is_presell_spec"`
}
type FulfillmentRule struct {
	// 现货发货模式规则
	NormalRule *NormalRule `json:"normal_rule"`
	// 阶梯发货模式规则
	StepRule *StepRule `json:"step_rule"`
	// 全款预售发货模式规则
	ProductPresellRule *ProductPresellRule `json:"product_presell_rule"`
	// SKU预售发货模式规则
	SkuPresellRule *SkuPresellRule `json:"sku_presell_rule"`
	// 现货+预售发货规则
	TimeSkuPresellWithNormalRule *TimeSkuPresellWithNormalRule `json:"time_sku_presell_with_normal_rule"`
	// 新预售发货模式规则
	TimeSkuPurePresellRule *TimeSkuPurePresellRule `json:"time_sku_pure_presell_rule"`
}
type OptionsItem struct {
	// 天无理由退货文案描述
	Name string `json:"name"`
	// 无理由退货的类型
	Value string `json:"value"`
}
type RequiredSpecDetailsItem struct {
	// 规格项名称
	SellPropertyName string `json:"sell_property_name"`
	// 规格项id
	SellPropertyId int64 `json:"sell_property_id"`
	// 规格值选项
	PropertyValues []PropertyValuesItem `json:"property_values"`
	// 是否支持备注
	SupportRemark bool `json:"support_remark"`
	// 是否可以自定义规格值
	SupportDiy bool `json:"support_diy"`
}
type NormalRule struct {
	// 特殊时间延迟发货规则
	DelayRule *DelayRule `json:"delay_rule"`
	// 是否支持
	Support bool `json:"support"`
	// 发货时效选项，如当日发、次日发、48小时
	DelayOptions []int64 `json:"delay_options"`
	// 是否是特殊的时间发货，可忽略
	IsSpecialDelayOption bool `json:"is_special_delay_option"`
}
type StepRule struct {
	// 是否支持
	Support bool `json:"support"`
	// 阶梯现货部分延迟返货时间范围
	DelayOptions []int64 `json:"delay_options"`
	// 是特殊时间延迟发货，可忽略
	IsSpecialDelayOption bool `json:"is_special_delay_option"`
	// 阶梯发货阶梯部分发货时间最小值
	StepMinDelay int64 `json:"step_min_delay"`
	// 阶梯发货阶梯部分发货时间最大值
	StepMaxDelay int64 `json:"step_max_delay"`
}
type TimeSkuSpecDetialItem struct {
	// 规格值
	SpecValue string `json:"spec_value"`
	// 是否预售时效
	IsPresellSpec bool `json:"is_presell_spec"`
	// product_time_spec_same_day 当日发 product_time_spec_next_day 次日发 product_time_spec_48h  48小时内发货 product_time_spec_5d  5天内发货 product_time_spec_15d  15天内发货 product_time_spec_45d  45天内发货
	SpecCode string `json:"spec_code"`
}
type SpuControlRule struct {
	// 是否支持spu发品
	SupportSpuProduct bool `json:"support_spu_product"`
	// 0不管控 1弱管控 2强管控
	ControlType int64 `json:"control_type"`
	// 是否支持新建spu
	SupportCreateSpu bool `json:"support_create_spu"`
	// 是否支持spu纠错
	SupportRectifySpu bool `json:"support_rectify_spu"`
	// 是否支持spu举报
	SupportReportSpu bool `json:"support_report_spu"`
	// 是否spu免审
	NoNeedAuditSpu bool `json:"no_need_audit_spu"`
}
type DeliveryAfterPresellConfig struct {
	// 延迟发货时间最小值
	MinDeliveryDays int64 `json:"min_delivery_days"`
	// 延迟发货时间最大值
	MaxDeliveryDays int64 `json:"max_delivery_days"`
	// 最长预售结束时间
	MaxPresellEndDays int64 `json:"max_presell_end_days"`
	// 是否需要人审，可忽略
	NeedAudit bool `json:"need_audit"`
}
type TimeSkuPresellWithNormalRule struct {
	// 是否支持
	Support bool `json:"support"`
	// 发货时效规格名称
	TimeSkuSpecName string `json:"time_sku_spec_name"`
	// 发货时效规格选项
	TimeSkuSpecDetial []TimeSkuSpecDetialItem `json:"time_sku_spec_detial"`
	// 预售门槛价，单位分
	MinPresellPrice int64 `json:"min_presell_price"`
}
type TimeSkuPurePresellRule struct {
	// 是否支持
	Support bool `json:"support"`
	// 发货时效规格名称
	TimeSkuSpecName string `json:"time_sku_spec_name"`
	// 发货时效规格选项
	TimeSkuSpecDetial []TimeSkuSpecDetialItem_4_4 `json:"time_sku_spec_detial"`
	// 预售门槛价，单位分
	MinPresellPrice int64 `json:"min_presell_price"`
}
type PriceRulesItem struct {
	// 是否可选，true表示可选，false表示不可选，如果不可选则表示商家没有设置过基础金价，需要前往抖店后台设置
	CanSelect bool `json:"can_select"`
	// 金价描述
	Desc string `json:"desc"`
	// 每克重黄金对应的价格，单位：分
	Value int64 `json:"value"`
	// 价格规则枚举：1是投资金2是首饰金
	Type int64 `json:"type"`
}
type SupplyDayReturnRule struct {
	// true 支持七天无理由，false 不支持七天无理由
	Enable bool `json:"enable"`
	// 可选的无理由退货选项列表
	Options []OptionsItem `json:"options"`
}
type PropertyValuesItem struct {
	// 规格值id
	SellPropertyValueId int64 `json:"sell_property_value_id"`
	// 规格值名称
	SellPropertyValueName string `json:"sell_property_value_name"`
}
type ComponentTemplateRule struct {
	// 是否展示尺码信息
	IsShow bool `json:"is_show"`
	// 尺码信息是否必填
	MustInput bool `json:"must_input"`
}
type MainImageThreeToFourRule struct {
	// 是否展示主图3:4信息
	IsShow bool `json:"is_show"`
	// 主图3:4信息是否必填
	MustInput bool `json:"must_input"`
}
type DelayRule struct {
	// 规则结束时间
	EndTime int64 `json:"end_time"`
	// 规则开始时间
	StartTime int64 `json:"start_time"`
	// 与规则配置类型对应的配置值
	ConfigValue int64 `json:"config_value"`
	// 支持的配置类型1：代表时间点，此时对应的config_value的值为支持的最晚发货的秒级时间戳2：代表相对时间，此时对应的config_value的值表示支付后几天发货；比如config_value=3表示支付后三天发货
	ConfigType int32 `json:"config_type"`
}
type DeliveryAfterPayConfig struct {
	// 延迟发货时间最小值
	MinDeliveryDays int64 `json:"min_delivery_days"`
	// 延迟发货时间最大值
	MaxDeliveryDays int64 `json:"max_delivery_days"`
	// 最长预售结束时间
	MaxPresellEndDays int64 `json:"max_presell_end_days"`
	// 是否需要人审，可忽略
	NeedAudit bool `json:"need_audit"`
}
type SkuPresellRule struct {
	// 是否支持
	Support bool `json:"support"`
	// 预售门槛价
	MinPresellPrice int64 `json:"min_presell_price"`
	// 是否支持支付结束后发货
	SupportDeliveryAfterPay bool `json:"support_delivery_after_pay"`
	// 是否支持预售结束后发货
	SupportDeliveryAfterPresell bool `json:"support_delivery_after_presell"`
	// 支付结束后规则
	DeliveryAfterPayConfig *DeliveryAfterPayConfig `json:"delivery_after_pay_config"`
	// 预售结束后规则
	DeliveryAfterPresellConfig *DeliveryAfterPresellConfig `json:"delivery_after_presell_config"`
}
type RecommendNameRule struct {
	// 当前类目id是否命中前缀推荐规则
	SatisfyPrefix bool `json:"satisfy_prefix"`
	// 命中规则的属性id详情
	PropertyIds []int64 `json:"property_ids"`
}
type ProductGetProductUpdateRuleData struct {
	// 履约规则
	FulfillmentRule *FulfillmentRule `json:"fulfillment_rule"`
	// 商品标题推荐规则
	RecommendNameRule *RecommendNameRule `json:"recommend_name_rule"`
	// 售后服务规则
	AfterSaleRule *AfterSaleRule `json:"after_sale_rule"`
	// 参考价相关规则
	ReferencePriceRule *ReferencePriceRule `json:"reference_price_rule"`
	// spu管控规则
	SpuControlRule *SpuControlRule `json:"spu_control_rule"`
	// 商品规格约束
	ProductSpecRule *ProductSpecRule `json:"product_spec_rule"`
	// 商品尺码模板配置规则
	ComponentTemplateRule *ComponentTemplateRule `json:"component_template_rule"`
	// 商品主图3:4配置规则
	MainImageThreeToFourRule *MainImageThreeToFourRule `json:"main_image_three_to_four_rule"`
	// 金价信息，计价金类目下sku价格可以按照公式进行计算（价格=克重*每克价格+加工费），本字段提供相关信息
	GoldPriceRule *GoldPriceRule `json:"gold_price_rule"`
}
