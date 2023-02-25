package shop_getShopCategory_response

import (
	"doudian.com/open/sdk_golang/core"
)

type ShopGetShopCategoryResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data []DataItem `json:"data"`
}
type DataItem struct {
	// 类目id，用于商品发布和更新，对应请求参数category_leaf_id字段
	Id int64 `json:"id"`
	// 类目名称
	Name string `json:"name"`
	// 类目级别：1，2，3级类目
	Level int64 `json:"level"`
	// 类目父级父类目id
	ParentId int64 `json:"parent_id"`
	// 是否是叶子节点；is_leaf=true表示是叶子节点，最小层级类目id。is_leaf=false表示不是子节点，请求参数cid=上一次响应参数id，直到获取最小层级类目id
	IsLeaf bool `json:"is_leaf"`
	// 类目使用有效；enable=true有效，如果enable=false表示该类目已经失效，请勿使用
	Enable bool `json:"enable"`
}
type ShopGetShopCategoryData struct {
	// 返回参数列表
	Data []DataItem `json:"data"`
}
