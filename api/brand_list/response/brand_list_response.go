package brand_list_response

import (
	"doudian.com/open/sdk_golang/core"
)

type BrandListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *BrandListData `json:"data"`
}
type AuthBrandListItem struct {
	// 品牌id
	BrandId int64 `json:"brand_id"`
	// 中文名
	NameCn string `json:"name_cn"`
	// 英文名
	NameEn string `json:"name_en"`
}
type BrandListItem struct {
	// 品牌id
	BrandId int64 `json:"brand_id"`
	// 中文名
	NameCn string `json:"name_cn"`
	// 英文名
	NameEn string `json:"name_en"`
}
type BrandListData struct {
	// （已停止使用）品牌id列表
	BrandIds []int64 `json:"brand_ids"`
	// （已停止使用）品牌信息
	BrandInfos map[int64]BrandInfosItem `json:"brand_infos"`
	// （已停止使用）总数
	Total int64 `json:"total"`
	// （已停止使用）还有更多
	HasMore bool `json:"has_more"`
	// 类目是否要求品牌有授权
	AuthRequired bool `json:"auth_required"`
	// 授权的品牌列表
	AuthBrandList []AuthBrandListItem `json:"auth_brand_list"`
	// 未授权的品牌列表
	BrandList []BrandListItem `json:"brand_list"`
}
type BrandInfosItem struct {
	// 品牌id
	BrandId int64 `json:"brand_id"`
	// 品牌中文名
	BrandNameCN string `json:"brand_name_c_n"`
	// 品牌英文名
	BrandNameEN string `json:"brand_name_e_n"`
	// 品牌评级 0-4
	Level int32 `json:"level"`
	// 品牌状态 1:上线, 2:下线
	Status int32 `json:"status"`
	// 品牌别名
	BrandAlias []string `json:"brand_alias"`
	// 创建时间
	CreateTimestamp int64 `json:"create_timestamp"`
	// 更新时间
	UpdateTimestamp int64 `json:"update_timestamp"`
	// 品牌审核状态 1:审核中, 2:审核通过, 3:审核拒绝, 4:送审失败
	AuditStatus int32 `json:"audit_status"`
	// 业务类型 0:国内, 1:跨境电商, 2:广告
	BizType int32 `json:"biz_type"`
	// 品牌logo地址
	Logo string `json:"logo"`
}
