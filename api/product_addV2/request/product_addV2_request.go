package product_addV2_request

import (
	"doudian.com/open/sdk_golang/api/product_addV2/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ProductAddV2Request struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ProductAddV2Param 
}
func (c *ProductAddV2Request) GetUrlPath() string{
	return "/product/addV2"
}


func New() *ProductAddV2Request{
	request := &ProductAddV2Request{
		Param: &ProductAddV2Param{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ProductAddV2Request) Execute(accessToken *doudian_sdk.AccessToken) (*product_addV2_response.ProductAddV2Response, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &product_addV2_response.ProductAddV2Response{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ProductAddV2Request) GetParamObject() interface{}{
	return c.Param
}


func (c *ProductAddV2Request) GetParams() *ProductAddV2Param{
	return c.Param
}


type PoiResource struct {
	// 有效时间段，从领取日开始计算，优先级高于ValidStart-ValidEnd
	ValidDays int64 `json:"valid_days,omitempty"`
	// 卡券有效开始时间，秒单位时间戳
	ValidStart int64 `json:"valid_start,omitempty"`
	// 卡券有效截止时间，秒单位时间戳
	ValidEnd int64 `json:"valid_end,omitempty"`
	// 客服电话
	ServiceNum string `json:"service_num,omitempty"`
	// 领取须知
	Notification string `json:"notification,omitempty"`
	// 平台产生券码: 1 ; 合作方api实时传入的券码: 4
	CodeType int64 `json:"code_type,omitempty"`
	// 券码总量，0/-1表示不限制，平台券时须\u003e0
	Count int64 `json:"count,omitempty"`
	// 0-不支持二次兑换，1-支持二次兑换
	CouponSecondExchange int64 `json:"couponSecondExchange,omitempty"`
	// 可核销总次数
	TotalCanUseCount int32 `json:"total_can_use_count,omitempty"`
	// 兑换链接
	Link string `json:"link,omitempty"`
	// 券码使用条件
	Condition string `json:"condition,omitempty"`
	// 数组当前只支持一个元素且只可为 1或2，其中 1 表示随时退+过期自动退，2 表示不支持退
	CouponReturnMethods []int64 `json:"coupon_return_methods,omitempty"`
}
type RecruitInfo struct {
	// 线索ID
	RecruitFollowId string `json:"recruit_follow_id,omitempty"`
	// 招商类型
	RecruitType string `json:"recruit_type,omitempty"`
}
type ProductAddV2Param struct {
	// 推荐使用，外部商家编码，支持字符串。超市小时达场景不推荐使用
	OuterProductId string `json:"outer_product_id,omitempty"`
	// 0-普通，3-虚拟，6玉石闪购，7云闪购
	ProductType *int64 `json:"product_type,omitempty"`
	// 叶子类目ID通过/shop/getShopCategory接口获取
	CategoryLeafId *int64 `json:"category_leaf_id,omitempty"`
	// （已废弃，用新字段product_format_new）属性名称|属性值 之间用|分隔, 多组之间用^分开
	ProductFormat string `json:"product_format,omitempty"`
	// 商品名称，最多60个字符(30个汉字)，不能含emoj表情
	Name *string `json:"name,omitempty"`
	// 商家推荐语，不能含emoj表情
	RecommendRemark string `json:"recommend_remark,omitempty"`
	// 商品轮播图，多张图片用 \"|\" 分开，第一张图为主图，最多5张，至少600x600，大小不超过1M。
	Pic *string `json:"pic,omitempty"`
	// 商品描述，目前只支持图片。多张图片用 \"|\" 分开。不能用其他网站的文本粘贴，这样会出现css样式不兼容。总图片数量不得超过50张。
	Description *string `json:"description,omitempty"`
	// 支付方式，0货到付款 1在线支付，2，货到付款+在线支付
	PayType int64 `json:"pay_type,omitempty"`
	// 海南免税生效。 0:离岛免税、1:邮寄、2:离岛自提+邮寄
	DeliveryMethod int32 `json:"delivery_method,omitempty"`
	// 海南免税：海关限购分类编码
	CdfCategory string `json:"cdf_category,omitempty"`
	// 1 减库存类型：1-拍下减库存 2-付款减库存
	ReduceType *int64 `json:"reduce_type,omitempty"`
	// 同店商品推荐：为空表示使用系统推荐；多个product_id使用“|”分开
	AssocIds string `json:"assoc_ids,omitempty"`
	// 运费模板id，传0表示包邮，通过/freightTemplate/list接口获取
	FreightId *int64 `json:"freight_id,omitempty"`
	// 重量
	Weight float64 `json:"weight,omitempty"`
	// 重量单位，0-kg, 1-g
	WeightUnit int64 `json:"weight_unit,omitempty"`
	// delivery_delay_day： 承诺发货时间，单位是天,不传则默认为2天。现货发货(presell_type=0)和阶梯发货模式(presell_type=2)时必填，支持传入9999 、1、 2 （分别表示当日发、次日发、48小时发），具体支持传入的参数范围/product/getProductUpdateRule
	DeliveryDelayDay int64 `json:"delivery_delay_day,omitempty"`
	// 发货模式，0-现货发货，1-预售发货，2-阶梯发货，默认0
	PresellType int64 `json:"presell_type,omitempty"`
	// 全款预售模式时的发货时间/阶梯模式下阶梯发货时间，具体支持传入的参数范围/product/getProductUpdateRule。
	PresellDelay int64 `json:"presell_delay,omitempty"`
	// 预售结束时间，格式2020-02-21 18:54:27，最多支持设置距离当前30天
	PresellEndTime string `json:"presell_end_time,omitempty"`
	// 是否支持7天无理由，0不支持，1支持，2支持（拆封后不支持）
	Supply7dayReturn int64 `json:"supply_7day_return,omitempty"`
	// 客服电话号码
	Mobile *string `json:"mobile,omitempty"`
	// false仅保存，true保存+提审
	Commit *bool `json:"commit,omitempty"`
	// （brand_id已废弃，用新字段standard_brand_id）品牌id (请求店铺授权品牌接口获取) (效果即为商品详情页面中的品牌字段)
	BrandId int64 `json:"brand_id,omitempty"`
	// 商家可见备注
	Remark string `json:"remark,omitempty"`
	// 外部product_id。超市小时达场景不推荐使用
	OutProductId int64 `json:"out_product_id,omitempty"`
	// 资质，可通过/product/qualificationConfig获取
	QualityList []QualityListItem `json:"quality_list,omitempty"`
	// 如果不填，则规格名为各级规格名用 \"-\" 自动生成
	SpecName string `json:"spec_name,omitempty"`
	// 店铺通用规格，能被同类商品通用。多规格用^分隔，规格与规格值用|分隔，多个规格值用,分隔。如果创建或编辑现货+预售商品或新预售商品时，须包含发货时效，如：颜色|黑色,白色,黄色^尺码|S,M,L^发货时效|5天内发货,15天内发货
	Specs *string `json:"specs,omitempty"`
	// sku详情，数量应该等于规格1规格2规格3，sku数量和规格组合数必须一致 sku不可售时，库存可设置为0。price单位为分。 delivery_infos为SKU物流信息，info_value为字符串类型（示例："12"），info_type填weight，info_unit支持mg,g,kg，超市小时达场景主品用普通库存，子品用区域库存（"sku_type": 1 // 区域库存，"stock_num_map": {"123": 99999 // 门店ID:库存数量}）; “gold_process_charge”为黄金加工费，只有计价金类目可填并且必填
	SpecPrices *string `json:"spec_prices,omitempty"`
	// 如颜色-尺寸, 颜色就是主规格, 颜色如黑,白,黄,它们图片url以逗号分隔 注：\"pic\"、\"description\"、\"spec_pic\"这三个字段里的传入的图片数量总和，不得超过50张。
	SpecPic string `json:"spec_pic,omitempty"`
	// 每个用户每次下单限购件数
	MaximumPerOrder int64 `json:"maximum_per_order,omitempty"`
	// 每个用户累计限购件数
	LimitPerBuyer int64 `json:"limit_per_buyer,omitempty"`
	// 每个用户每次下单至少购买的件数
	MinimumPerOrder int64 `json:"minimum_per_order,omitempty"`
	// 属性，通过/product/getCatePropertyV2获取 格式：{"property_id":[{"value":value,"name":"property_name","diy_type":0}]} name的类型是string，value和diy_type的类型是number
	ProductFormatNew string `json:"product_format_new,omitempty"`
	// spu_id
	SpuId int64 `json:"spu_id,omitempty"`
	// 可预约发货天数
	AppointDeliveryDay int64 `json:"appoint_delivery_day,omitempty"`
	// third_url
	ThirdUrl string `json:"third_url,omitempty"`
	// extra
	Extra string `json:"extra,omitempty"`
	// src
	Src string `json:"src,omitempty"`
	// 品牌id，通过/brand/list获取，无品牌id则传596120136
	StandardBrandId int64 `json:"standard_brand_id,omitempty"`
	// 卡券类型需要传true
	NeedCheckOut bool `json:"need_check_out,omitempty"`
	// 卡券信息
	PoiResource *PoiResource `json:"poi_resource,omitempty"`
	// 汽车vin码
	CarVinCode string `json:"car_vin_code,omitempty"`
	// 0：全款预售，1：sku预售，2：现货+预售 ，3：新预售
	PresellConfigLevel int64 `json:"presell_config_level,omitempty"`
	// 充值模式
	NeedRechargeMode bool `json:"need_recharge_mode,omitempty"`
	// 账号模板id
	AccountTemplateId string `json:"account_template_id,omitempty"`
	// 全款预售和sku预售时传递，其他不传： 0 预售结束后发货  1支付完成后发货
	PresellDeliveryType int64 `json:"presell_delivery_type,omitempty"`
	// 白底图url(仅素材中心url有效)，白底图比例要求1:1
	WhiteBackGroundPicUrl string `json:"white_back_ground_pic_url,omitempty"`
	// 3:4长图url(仅素材中心url有效)
	LongPicUrl string `json:"long_pic_url,omitempty"`
	// 推荐传入，售后服务参数,key:value格式。supply_day_return_selector(7天无理由选项)：N天-政策代号，N只支持7和15，政策代号枚举https://bytedance.feishu.cn/docs/doccnF946oh1c98e7mo9DlYZtig 。 supply_red_ass_return（红屁屁无忧）：0不支持，1支持。supply_allergy_return（过敏无忧,仅跨境可选）：0不支持，1支持。 damaged_order_return（坏损包退）：0不支持，1支持。 support_allergy_returnV2（过敏包退，境内可选）：0不支持，1支持
	AfterSaleService map[string]string `json:"after_sale_service,omitempty"`
	// 售卖方式;0:全渠道手售卖,1:仅指定直播间售卖
	SellChannel []int64 `json:"sell_channel,omitempty"`
	// 审核通过后上架售卖时间配置:0-立即上架售卖 1-放入仓库
	StartSaleType int64 `json:"start_sale_type,omitempty"`
	// 特殊日期延迟发货规则
	DelayRule *DelayRule `json:"delay_rule,omitempty"`
	// 主图视频ID，可以先通过https://op.jinritemai.com/docs/api-docs/69/1617接口上传视频，获取审核通过的视频素材ID进行传入 任务需要验证
	MaterialVideoId string `json:"material_video_id,omitempty"`
	// 提取方式新字段，推荐使用。"0": 普通商品-使用物流发货, "1": 虚拟商品-无需物流与电子交易凭证, "2": 虚拟商品-使用电子交易凭证,  "3": 虚拟商品-充值直连
	PickupMethod string `json:"pickup_method,omitempty"`
	// 尺码模板ID
	SizeInfoTemplateId int64 `json:"size_info_template_id,omitempty"`
	// 外部商品url
	SubstituteGoodsUrl string `json:"substitute_goods_url,omitempty"`
	// 销售渠道类型，可选onlineOnly（纯电商，仅在线上售卖）或sameAsOffline（专柜同款，线上线下都有售卖），云零售商家（https://fxg.jinritemai.com/ffa/merchant-growth/cloud-retail）可以设置
	SaleChannelType string `json:"sale_channel_type,omitempty"`
	// 招商信息
	RecruitInfo *RecruitInfo `json:"recruit_info,omitempty"`
	// 门店ID
	StoreId int64 `json:"store_id,omitempty"`
	// 主商品ID
	MainProductId int64 `json:"main_product_id,omitempty"`
	// 限售模板ID
	SaleLimitId int64 `json:"sale_limit_id,omitempty"`
	// 通过/product/getRecommendName接口推荐的商品标题前缀（注意系统根据类目属性生成前缀字符串时，末尾有一个空格）
	NamePrefix string `json:"name_prefix,omitempty"`
	// 参考价，单位分，需大于商品价格并小于商品价格的10倍
	ReferencePrice int64 `json:"reference_price,omitempty"`
	// 参考价凭证信息
	ReferencePriceCertificate *ReferencePriceCertificate `json:"reference_price_certificate,omitempty"`
	// 商品主图3:4，多张图片用 \"|\" 分开，最多5张，大小不超过1M。
	MainImageThreeToFour string `json:"main_image_three_to_four,omitempty"`
	// 商品价格规则信息
	UnitPriceInfo *UnitPriceInfo `json:"unit_price_info,omitempty"`
	// 前置质检相关（特定二手商家、特定二手类目使用）
	QualityInspectionInfo *QualityInspectionInfo `json:"quality_inspection_info,omitempty"`
}
type QualityAttachmentsItem struct {
	// 1-图片，2-pdf
	MediaType *int64 `json:"media_type,omitempty"`
	// 凭证url
	Url *string `json:"url,omitempty"`
}
type QualityListItem struct {
	// 资质key
	QualityKey *string `json:"quality_key,omitempty"`
	// 资质名称
	QualityName *string `json:"quality_name,omitempty"`
	// 资质
	QualityAttachments *[]QualityAttachmentsItem `json:"quality_attachments,omitempty"`
}
type DelayRule struct {
	// 是否开启特殊日期延迟发货
	Enable bool `json:"enable,omitempty"`
	// 1 时间点；2 时间范围
	ConfigType int32 `json:"config_type,omitempty"`
	// 特殊日期延迟发货时间最晚发货时间，时间戳，单位秒：当config_type=1时，传时间戳，代表最晚x发货；当config_type=2时，传天数，代表延迟x天发货
	ConfigValue int64 `json:"config_value,omitempty"`
	// 特殊日期延迟发货时间下单开始时间，时间戳，单位秒
	StartTime int64 `json:"start_time,omitempty"`
	// 特殊日期延迟发货时间下单结束时间，时间戳，单位秒
	EndTime int64 `json:"end_time,omitempty"`
}
type ReferencePriceCertificate struct {
	// 通过/product/getProductUpdateRule获取可选的参考价格类型
	CertificateType int64 `json:"certificate_type,omitempty"`
	// 图片url需要使用商品素材中心的url并且只能有一张
	CertificateUrls []string `json:"certificate_urls,omitempty"`
}
type UnitPriceInfo struct {
	// 1表示投资金2表示首饰金
	PriceRuleType int32 `json:"price_rule_type,omitempty"`
}
type QualityInspectionInfo struct {
	// 是否支持前置质检
	Supported bool `json:"supported,omitempty"`
	// 机构编码，请通过/inspection/QueryBtasAgencyList接口获取
	Agency string `json:"agency,omitempty"`
	// 质检证书编码
	CertificateCode string `json:"certificate_code,omitempty"`
	// 1: 单库存模式，只允许售卖一个sku；2: 多库存模式，不限售卖次数
	Mode int32 `json:"mode,omitempty"`
}
