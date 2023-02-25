package product_editCbProduct_request

import (
	"doudian.com/open/sdk_golang/api/product_editCbProduct/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ProductEditCbProductRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ProductEditCbProductParam 
}
func (c *ProductEditCbProductRequest) GetUrlPath() string{
	return "/product/editCbProduct"
}


func New() *ProductEditCbProductRequest{
	request := &ProductEditCbProductRequest{
		Param: &ProductEditCbProductParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ProductEditCbProductRequest) Execute(accessToken *doudian_sdk.AccessToken) (*product_editCbProduct_response.ProductEditCbProductResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &product_editCbProduct_response.ProductEditCbProductResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ProductEditCbProductRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *ProductEditCbProductRequest) GetParams() *ProductEditCbProductParam{
	return c.Param
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
type PoiResource struct {
	// 券码使用条件
	Condition string `json:"condition,omitempty"`
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
	// 可核销总次数
	TotalCanUseCount *int32 `json:"total_can_use_count,omitempty"`
	// 0-不支持二次兑换，1-支持二次兑换
	CouponSecondExchange *int64 `json:"coupon_second_exchange,omitempty"`
	// 兑换链接
	Link string `json:"link,omitempty"`
	// 券码总量，0/-1表示不限制，平台券时须\u003e0
	Count int64 `json:"count,omitempty"`
}
type ProductEditCbProductParam struct {
	// 商品ID
	ProductId *int64 `json:"product_id,omitempty"`
	// 0-普通，3-虚拟，6玉石闪购，7云闪购
	ProductType int64 `json:"product_type,omitempty"`
	// 叶子类目ID通过/shop/getShopCategory接口获取
	CategoryLeafId int64 `json:"category_leaf_id,omitempty"`
	// 属性名称|属性值 之间用|分隔, 多组之间用^分开
	ProductFormat string `json:"product_format,omitempty"`
	// 商品名称，最多60个字符(30个汉字)，不能含emoj表情
	Name string `json:"name,omitempty"`
	// 商家推荐语，不能含emoj表情
	RecommendRemark string `json:"recommend_remark,omitempty"`
	// 商品轮播图，多张图片用 \"|\" 分开，第一张图为主图，最多5张，至少600x600，大小不超过1M 注：\"pic\"、\"description\"、\"spec_pic\"这三个字段里的传入的图片数量总和，不得超过50张
	Pic string `json:"pic,omitempty"`
	// 商品轮播图，多张图片用 \"|\" 分开，第一张图为主图，最多5张，至少600x600，大小不超过1M 注：\"pic\"、\"description\"、\"spec_pic\"这三个字段里的传入的图片数量总和，不得超过50张
	Description string `json:"description,omitempty"`
	// 支付方式，0货到付款 1在线支付，2，货到付款+在线支付
	PayType int64 `json:"pay_type,omitempty"`
	// 海南免税生效。 0:离岛免税、1:邮寄、2:离岛自提+邮寄
	DeliveryMethod int32 `json:"delivery_method,omitempty"`
	// 海南免税：海关限购分类编码
	CdfCategory string `json:"cdf_category,omitempty"`
	// 1 减库存类型：1-拍下减库存 2-付款减库存
	ReduceType int64 `json:"reduce_type,omitempty"`
	// 同店商品推荐：为空表示使用系统推荐；多个product_id使用“|”分开
	AssocIds string `json:"assoc_ids,omitempty"`
	// 运费模板id，传0表示包邮，通过/freightTemplate/list接口获取
	FreightId int64 `json:"freight_id,omitempty"`
	// 重量
	Weight float64 `json:"weight,omitempty"`
	// 重量单位，0-kg, 1-g
	WeightUnit int64 `json:"weight_unit,omitempty"`
	// delivery_delay_day： 承诺发货时间，单位是天,不传则默认为2天。现货发货(presell_type=0)和阶梯发货模式(presell_type=2)时必填，支持传入9999 、1、 2 （分别表示当日发、次日发、48小时发），具体支持传入的参数范围/product/getProductUpdateRule
	DeliveryDelayDay int64 `json:"delivery_delay_day,omitempty"`
	// 发货模式，0-现货发货，1-预售发货，2-阶梯发货，默认0
	PresellType int64 `json:"presell_type,omitempty"`
	// 全款预售模式时的发货时间/阶梯模式下阶梯发货时间，具体支持传入的参数范围/product/getProductUpdateRule
	PresellDelay int64 `json:"presell_delay,omitempty"`
	// 预售结束时间，格式2020-02-21 18:54:27，最多支持设置距离当前30天
	PresellEndTime string `json:"presell_end_time,omitempty"`
	// 是否支持7天无理由，0不支持，1支持，2支持（拆封后不支持）
	Supply7dayReturn int64 `json:"supply_7day_return,omitempty"`
	// 客服电话号码
	Mobile string `json:"mobile,omitempty"`
	// false仅保存，true保存+提审
	Commit *bool `json:"commit,omitempty"`
	// 品牌id (请求店铺授权品牌接口获取) (效果即为商品详情页面中的品牌字段)
	BrandId int64 `json:"brand_id,omitempty"`
	// 商家可见备注
	Remark string `json:"remark,omitempty"`
	// 123
	OutProductId int64 `json:"out_product_id,omitempty"`
	// 资质信息，可通过/product/qualificationConfig获取
	QualityList []QualityListItem `json:"quality_list,omitempty"`
	// 如果不填，则规格名为各级规格名用 \"-\" 自动生成
	SpecName string `json:"spec_name,omitempty"`
	// 店铺通用规格，能被同类商品通用。多规格用^分隔，规格与规格值用|分隔，多个规格值用,分隔。如果创建或编辑现货+预售商品或新预售商品时，须包含发货时效，如：颜色|黑色,白色,黄色^尺码|S,M,L^发货时效|5天内发货,15天内发货
	Specs string `json:"specs,omitempty"`
	// sku详情，数量应该等于规格1规格2规格3，sku数量和规格组合数必须一致 sku不可售时，库存可设置为0
	SpecPrices string `json:"spec_prices,omitempty"`
	// 如颜色-尺寸, 颜色就是主规格, 颜色如黑,白,黄,它们图片url以逗号分隔 注：\"pic\"、\"description\"、\"spec_pic\"这三个字段里的传入的图片数量总和，不得超过50张
	SpecPic string `json:"spec_pic,omitempty"`
	// 每个用户每次下单限购件数
	MaximumPerOrder int64 `json:"maximum_per_order,omitempty"`
	// 每个用户累计限购件数
	LimitPerBuyer int64 `json:"limit_per_buyer,omitempty"`
	// 每个用户每次下单至少购买的件数
	MinimumPerOrder int64 `json:"minimum_per_order,omitempty"`
	// 通过/product/getCatePropertyV2获取 格式：{"property_id":[{"value":value,"name":"property_name","diy_type":0}]}, name的类型是string，value和diy_type的类型是number
	ProductFormatNew string `json:"product_format_new,omitempty"`
	// spu id
	SpuId int64 `json:"spu_id,omitempty"`
	// 可预约发货天数
	AppointDeliveryDay int64 `json:"appoint_delivery_day,omitempty"`
	// third_url
	ThirdUrl string `json:"third_url,omitempty"`
	// extra
	Extra string `json:"extra,omitempty"`
	// src
	Src string `json:"src,omitempty"`
	// 外部product id
	OuterProductId string `json:"outer_product_id,omitempty"`
	// standard_brand_id，通过/brand/list获取，无品牌id则传596120136
	StandardBrandId int64 `json:"standard_brand_id,omitempty"`
	// 卡券类型需要传true
	NeedCheckOut bool `json:"need_check_out,omitempty"`
	// 卡券信息
	PoiResource *PoiResource `json:"poi_resource,omitempty"`
	// 使用qualityList覆盖当前商品资质，qualityList传空代表清空
	ForceUseQualityList bool `json:"force_use_quality_list,omitempty"`
	// VIN11111111111111
	CarVinCode string `json:"car_vin_code,omitempty"`
	// 充值模式
	NeedRechargeMode bool `json:"need_recharge_mode,omitempty"`
	// 0：全款预售，1：sku预售，2：现货+预售 ，3：新预售
	PresellConfigLevel int64 `json:"presell_config_level,omitempty"`
	// 多账号模板
	AccountTemplateId string `json:"account_template_id,omitempty"`
	// 全款预售和sku预售时传递，其他不传： 0 预售结束后发货 1支付完成后发货
	PresellDeliveryType int64 `json:"presell_delivery_type,omitempty"`
	// 白底图url(仅素材中心url有效)
	WhiteBackGroundPicUrl string `json:"white_back_ground_pic_url,omitempty"`
	// 3:4长图url(仅素材中心url有效)
	LongPicUrl string `json:"long_pic_url,omitempty"`
	// 售后服务参数,key:value格式，supply_7day_return(7天无理由)： 0不支持，1支持，2支持（拆封后不支持） supply_day_return_selector(7天无理由标签)：N天-政策代号，N只支持7和15,supply_red_ass_return（红屁屁无忧）：0不支持，1支持 supply_allergy_return（过敏无忧,仅跨境可选）：0不支持，1支持 damaged_order_return（坏损包退）：0不支持，1支持 support_allergy_returnV2（过敏包退，境内可选）：0不支持，1支持
	AfterSaleService map[string]string `json:"after_sale_service,omitempty"`
	// 售卖方式;0:全渠道手售卖,1:仅指定直播间售卖
	SellChannel []int64 `json:"sell_channel,omitempty"`
	// 审核通过后上架售卖时间配置:0-立即上架售卖 1-放入仓库
	StartSaleType int64 `json:"start_sale_type,omitempty"`
	// 跨境物流相关信息
	LogisticsInfo string `json:"logistics_info,omitempty"`
	// 1
	PriceHasTax string `json:"price_has_tax,omitempty"`
	// 尺码表id
	SizeInfoTemplateId int64 `json:"size_info_template_id,omitempty"`
}