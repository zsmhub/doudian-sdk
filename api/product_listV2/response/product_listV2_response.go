package product_listV2_response

import (
	"doudian.com/open/sdk_golang/core"
)

type ProductListV2Response struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *ProductListV2Data `json:"data"`
}
type CategoryDetail struct {
	// 一级类目
	FirstCid int64 `json:"first_cid"`
	// 二级类目
	SecondCid int64 `json:"second_cid"`
	// 三级类目
	ThirdCid int64 `json:"third_cid"`
	// 四级类目
	FourthCid int64 `json:"fourth_cid"`
	// 一级类目名称
	FirstCname string `json:"first_cname"`
	// 二级类目名称
	SecondCname string `json:"second_cname"`
	// 三级类目名称
	ThirdCname string `json:"third_cname"`
	// 四级类目名称
	FourthCname string `json:"fourth_cname"`
}
type ChannelMainProduct struct {
	// 主品的商品id
	ProductId int64 `json:"product_id"`
	// 店铺id
	ShopId int64 `json:"shop_id"`
}
type DataItem struct {
	// 商品ID，抖店系统生成，店铺下唯一
	ProductId int64 `json:"product_id"`
	// 商品在店铺中状态: 0-在线；1-下线；2-删除；详见商品状态机：https://op.jinritemai.com/docs/question-docs/92/2070
	Status int64 `json:"status"`
	// 商品审核状态: 1-未提交；2-待审核；3-审核通过；4-审核未通过；5-封禁；7-审核通过待上架；详见商品状态机：https://op.jinritemai.com/docs/question-docs/92/2070
	CheckStatus int64 `json:"check_status"`
	// 【已废弃】划线价，单位分
	MarketPrice int64 `json:"market_price"`
	// 【已废弃】售价，单位分
	DiscountPrice int64 `json:"discount_price"`
	// 商品图片url，返回商品主图的第一张图 (主要用来辅助页面展示)
	Img string `json:"img"`
	// 商品标题，规则：至少输入8个字（16个字符）以上~输入30个字（60个字符）以内。；标题不规范会引起商品下架，影响您的正常销售，详见商品发布规范：https://school.jinritemai.com/doudian/web/article/101800?from=shop_article
	Name string `json:"name"`
	// 【已废弃】支持的支付方式：0货到付款 1在线支付 2两者都支持
	PayType int64 `json:"pay_type"`
	// 商品类型；0-普通；1-新客商品；3-虚拟；6-玉石闪购；7-云闪购 ；127-其他类型；
	ProductType int64 `json:"product_type"`
	// 商品规格，新增商品是全局唯一，注意：有部分历史存量商品可能存在规格复用
	SpecId int64 `json:"spec_id"`
	// 【已废弃】佣金比例
	CosRatio int64 `json:"cos_ratio"`
	// 商品创建时间，unix时间戳，单位：秒；
	CreateTime int64 `json:"create_time"`
	// 商品更新时间，unix时间戳，单位：秒；
	UpdateTime int64 `json:"update_time"`
	// 【即将废弃】外部商家编码，商家自定义字段。推荐使用，outer_product_id字段
	OutProductId int64 `json:"out_product_id"`
	// 商品详情，最大支持50张图片，单张详情图宽高比不超2000*2000px，大小5M内，仅支持jpg/jpeg/png格式；注意：商品详情不规范，会引起商品下架，影响您的正常销售：https://school.jinritemai.com/doudian/web/article/101800?from=shop_article
	Description string `json:"description"`
	// 【已废弃】手机号
	Mobile string `json:"mobile"`
	// 额外信息，如资质
	Extra string `json:"extra"`
	// 【已废弃】商家推荐语
	RecommendRemark string `json:"recommend_remark"`
	// 类目详情；商品类目id可使用【/shop/getShopCategory】查询
	CategoryDetail *CategoryDetail `json:"category_detail"`
	// 外部商家编码，商家自定义字段，支持最多 255个字符
	OuterProductId string `json:"outer_product_id"`
	// 是否是组合商品，true-是，false-不是；
	IsPackageProduct bool `json:"is_package_product"`
	// 商品关联的组合主商品ID；当is_package_product=true，返回的是组套商品的product_id； 当is_package_product=false，返回当前商品的product_id；补充返回规则：参与的组套商品，下线会展示，目前组合商品支持下线的商品作为子品进行组套，删除会展示，但重新上架组套商品时，会被从列表中删除。
	PackageProductList []int64 `json:"package_product_list"`
	// 当is_package_product=true，返回的是组套商品的子品product_id； 当is_package_product=false，不返回；补充返回规则：参与的组套商品product_id，下线会展示，目前组合商品支持下线的商品作为子品进行组套，删除会展示，但重新上架组套商品时，会被从列表中删除。
	SubProductList []int64 `json:"sub_product_list"`
	// 小时达子品绑定的主品信息
	ChannelMainProduct *ChannelMainProduct `json:"channel_main_product"`
}
type ProductListV2Data struct {
	// 商品数据
	Data []DataItem `json:"data"`
	// 本次查询返回的商品总数
	Total int64 `json:"total"`
	// 当前页码
	Page int64 `json:"page"`
	// 页数（每页数量）
	Size int64 `json:"size"`
}
