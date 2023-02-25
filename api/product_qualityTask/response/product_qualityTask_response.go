package product_qualityTask_response

import (
	"doudian.com/open/sdk_golang/core"
)

type ProductQualityTaskResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *ProductQualityTaskData `json:"data"`
}
type ProblemTypeDistributionItem struct {
	// 优化项代号，1-商品信息不规范，2-关键信息待优化，3-商品素材缺失
	TypeKey int64 `json:"type_key"`
	// 优化项具体名字
	TypeName string `json:"type_name"`
	// 具体问题数量
	Num int64 `json:"num"`
}
type ProductQualityTaskData struct {
	// 总共需要诊断的商品数
	ProductNumTotal int64 `json:"product_num_total"`
	// 已经诊断过的商品数
	ProductNumFinished int64 `json:"product_num_finished"`
	// 任务状态，0-初始化，1-进行中，2-已完成
	TaskStatus int64 `json:"task_status"`
	// 待优化商品数，仅brief_only为false返回
	ProductNumToImproveTotal int64 `json:"product_num_to_improve_total"`
	// 可优化项总数，仅brief_only为false返回
	ProblemNumTotal int64 `json:"problem_num_total"`
	// 待优化项数量，仅brief_only为false返回
	ProblemNumToImprove int64 `json:"problem_num_to_improve"`
	// 任务完成时间，仅brief_only为false返回
	TaskFinishTime string `json:"task_finish_time"`
	// 问题类型分布，仅brief_only为false返回
	ProblemTypeDistribution []ProblemTypeDistributionItem `json:"problem_type_distribution"`
	// 任务ID
	TaskId int64 `json:"task_id"`
	// 达标率，百分比
	StandardRate float64 `json:"standard_rate"`
	// 是否达标
	IsStandard bool `json:"is_standard"`
	// 达标商品数
	MeetStandardNum int64 `json:"meet_standard_num"`
}
