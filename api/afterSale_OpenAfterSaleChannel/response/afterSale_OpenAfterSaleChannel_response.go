package afterSale_OpenAfterSaleChannel_response

import (
	"doudian.com/open/sdk_golang/core"
)

type AfterSaleOpenAfterSaleChannelResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *AfterSaleOpenAfterSaleChannelData `json:"data"`
}
type Conclusion struct {
	// 匹配到超级售后的类型，1是超售后期售后，2是EP订单超级售后，3是虚拟订单超级售后，4是超售后次数售后
	MatchConclusion int32 `json:"match_conclusion"`
	// 匹配结果的解释，成功时为空
	MatchMessage string `json:"match_message"`
	// 当前超级售后可以发起的售后类型，0是退货退款，1是已发货仅退款，2是未发货仅退款，3是换货
	CanApplyTypeList []int64 `json:"can_apply_type_list"`
	// 匹配是否成功，当match_conclusion不为0且can_apply_type_list不是空的时候，此值为true
	MatchSuccess bool `json:"match_success"`
}
type AfterSaleOpenAfterSaleChannelData struct {
	// 打开售后通道结论
	Conclusion *Conclusion `json:"conclusion"`
}
