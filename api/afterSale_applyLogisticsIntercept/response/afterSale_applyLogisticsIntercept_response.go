package afterSale_applyLogisticsIntercept_response

import (
	"doudian.com/open/sdk_golang/core"
)

type AfterSaleApplyLogisticsInterceptResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *AfterSaleApplyLogisticsInterceptData `json:"data"`
}
type CurProduct struct {
	// 商品单订单号
	OrderId int64 `json:"order_id"`
	// 商品图片
	ProductImage string `json:"product_image"`
	// 商品名称
	ProductName string `json:"product_name"`
	// 商品规格
	ProductSpec string `json:"product_spec"`
	// 商品标签
	Tags []string `json:"tags"`
	// 单价
	Price int64 `json:"price"`
	// 数量
	Amount int64 `json:"amount"`
}
type OtherProductsItem struct {
	// 商品单订单号
	OrderId int64 `json:"order_id"`
	// 商品图片
	ProductImage string `json:"product_image"`
	// 商品名称
	ProductName string `json:"product_name"`
	// 商品规格
	ProductSpec string `json:"product_spec"`
	// 商品标签
	Tags []string `json:"tags"`
	// 单价
	Price int64 `json:"price"`
	// 数量
	Amount int64 `json:"amount"`
}
type InterceptResultsItem struct {
	// 物流公司编码
	CompanyCode string `json:"company_code"`
	// 物流公司名称
	CompanyName string `json:"company_name"`
	// 物流单号
	TrackingNo string `json:"tracking_no"`
	// 包裹价值（分）
	ValueAmount int64 `json:"value_amount"`
	// 是否可拦截（拦截详情时返回）
	CanIntercept bool `json:"can_intercept"`
	// 是否拦截成功（发起拦截时返回）
	IsSuccess bool `json:"is_success"`
	// 不可拦截原因编码
	UnavailableReasonCode int64 `json:"unavailable_reason_code"`
	// 不可拦截原因文案
	UnavailableReason string `json:"unavailable_reason"`
	// 拦截费用（分），（拦截详情时返回，不可拦截时无意义）
	InterceptCost int64 `json:"intercept_cost"`
	// 当前售后商品信息
	CurProduct *CurProduct `json:"cur_product"`
	// 其它商品列表
	OtherProducts []OtherProductsItem `json:"other_products"`
	// 其他商品件数
	OtherProductAmount int64 `json:"other_product_amount"`
	// 结算方式 0是线上结算 1是线下结算
	SettlementMode int32 `json:"settlement_mode"`
}
type AfterSaleApplyLogisticsInterceptData struct {
	// 物流拦截结果
	InterceptResults []InterceptResultsItem `json:"intercept_results"`
	// 拦截成功次数
	SuccessCount int64 `json:"success_count"`
	// 拦截失败次数
	FailedCount int64 `json:"failed_count"`
	// 不可拦截编码(failed_count=1时有意义)
	UnavailableReasonCode int64 `json:"unavailable_reason_code"`
	// 不可拦截原因(failed_count=1时有意义)
	UnavailableReason string `json:"unavailable_reason"`
	// 售后单退款总金额
	RefundAmount int64 `json:"refund_amount"`
}
