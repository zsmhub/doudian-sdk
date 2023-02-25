package product_CategoryDimList_response

import (
	"doudian.com/open/sdk_golang/core"
)

type ProductCategoryDimListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data []DataItem `json:"data"`
}
type DataItem struct {
	// 类目ID
	Id int64 `json:"id"`
	// 类目名
	Name string `json:"name"`
	// 类目层级
	Level int32 `json:"level"`
	// 上级ID，一级类目上级ID为0
	ParentId int64 `json:"parent_id"`
	// 是否叶子类目
	IsLeaf int32 `json:"is_leaf"`
	// 行业名
	VerticalCategoryNew string `json:"vertical_category_new"`
	// 行业英文名
	VerticalCategoryCodeNew string `json:"vertical_category_code_new"`
}
type ProductCategoryDimListData struct {
	// 返回结果
	Data []DataItem `json:"data"`
}
