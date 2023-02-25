package brand_getSug_response

import (
	"doudian.com/open/sdk_golang/core"
)

type BrandGetSugResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *BrandGetSugData `json:"data"`
}
type BrandGetSugData struct {
	// 品牌信息列表
	SugList []SugListItem `json:"sug_list"`
}
type SugListItem struct {
	// 品牌ID
	BrandId int64 `json:"brand_id"`
	// 品牌中文名
	BrandNameCN string `json:"brand_name_c_n"`
	// 品牌英文名
	BrandNameEN string `json:"brand_name_e_n"`
	// 品牌等级
	Level int32 `json:"level"`
	// 品牌状态：在线
	Status int32 `json:"status"`
	// 品牌别名
	BrandAlias []string `json:"brand_alias"`
	// 创建时间
	CreateTimestamp int64 `json:"create_timestamp"`
	// 修改时间
	UpdateTimestamp int64 `json:"update_timestamp"`
	// 审核情况 1. 审核中 2. 审核通过 3. 审核拒绝 4. 送审失败
	AuditStatus int32 `json:"audit_status"`
	// 业务线类型: 0. 国内品牌 1. 跨境品牌 3. 广告
	BizType int32 `json:"biz_type"`
	// 品牌logo
	Logo string `json:"logo"`
	// 额外信息
	Extra map[string]string `json:"extra"`
}
