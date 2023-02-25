package product_qualificationConfig_response

import (
	"doudian.com/open/sdk_golang/core"
)

type ProductQualificationConfigResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *ProductQualificationConfigData `json:"data"`
}
type ProductQualificationConfigData struct {
	// 资质列表
	ConfigList []ConfigListItem `json:"config_list"`
}
type RuleClauseItem struct {
	// 属性值id数组
	PropertyValues []int64 `json:"property_values"`
	// 等于、不等于
	OperandStr string `json:"operand_str"`
}
type MatchableRuleItem struct {
	// key为类目属性id，当该属性填了如下属性值时，命中规则
	RuleClause map[int64]RuleClauseItem `json:"rule_clause"`
	// 命中该规则后 该资质是否必填
	IsQualificationRequired bool `json:"is_qualification_required"`
}
type ConfigListItem struct {
	// 资质ID
	Key string `json:"key"`
	// 资质名
	Name string `json:"name"`
	// 填写提示
	TextList []string `json:"text_list"`
	// 是否必填
	IsRequired bool `json:"is_required"`
	// 商品类目属性可能触发的规则（会导致资质是否必填发生变化）
	MatchableRule []MatchableRuleItem `json:"matchable_rule"`
}
