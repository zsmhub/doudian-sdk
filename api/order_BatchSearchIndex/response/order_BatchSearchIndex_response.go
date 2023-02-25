package order_BatchSearchIndex_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderBatchSearchIndexResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderBatchSearchIndexData `json:"data"`
}
type PlainToEncryptIndexListItem struct {
	// 明文
	Plain string `json:"plain"`
	// 索引串
	SearchIndex string `json:"search_index"`
}
type CustomErr struct {
	// 被风控状态码
	ErrCode int64 `json:"err_code"`
	// 产生索引串被风控返回次信息
	ErrMsg string `json:"err_msg"`
}
type OrderBatchSearchIndexData struct {
	// 明文转化为索引穿列表
	PlainToEncryptIndexList []PlainToEncryptIndexListItem `json:"plain_to_encrypt_index_list"`
	// 业务错误
	CustomErr *CustomErr `json:"custom_err"`
}
