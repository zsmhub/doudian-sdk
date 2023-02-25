package spu_queryBookNameByISBN_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SpuQueryBookNameByISBNResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SpuQueryBookNameByISBNData `json:"data"`
}
type DataItem struct {
	// spu id
	SpuId string `json:"spu_id"`
	// 图书的书名
	BookName string `json:"book_name"`
	// 类目ID
	CategoryLeafId int64 `json:"category_leaf_id"`
}
type SpuQueryBookNameByISBNData struct {
	// 查询到的数据
	Data []DataItem `json:"data"`
	// 总数
	Total int64 `json:"total"`
}
