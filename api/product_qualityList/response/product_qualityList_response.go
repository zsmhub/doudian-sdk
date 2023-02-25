package product_qualityList_response

import (
	"doudian.com/open/sdk_golang/core"
)

type ProductQualityListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *ProductQualityListData `json:"data"`
}
type ProblemTypeDistributionItem struct {
	// 待优化问题类型
	TypeName string `json:"type_name"`
	// 问题数量
	Num int64 `json:"num"`
}
type QualityListItem struct {
	// 商品ID
	ProductId int64 `json:"product_id"`
	// 商品名字
	ProductName string `json:"product_name"`
	// 待优化问题数量
	ProblemNumToImprove int64 `json:"problem_num_to_improve"`
	// 待优化问题分布列表
	ProblemTypeDistribution []ProblemTypeDistributionItem `json:"problem_type_distribution"`
	// 商品是否达标，1达标，2不达标
	MeetStandard int64 `json:"meet_standard"`
	// 商品基础分
	BaseScore int64 `json:"base_score"`
}
type ProductQualityListData struct {
	// 商品质量列表
	QualityList []QualityListItem `json:"quality_list"`
	// 店铺待优化商品总量
	Total int64 `json:"total"`
}
