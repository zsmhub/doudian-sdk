package order_batchSensitive_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderBatchSensitiveResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderBatchSensitiveData `json:"data"`
}
type OrderBatchSensitiveData struct {
	// 脱敏信息列表
	DecryptInfos []DecryptInfosItem `json:"decrypt_infos"`
	// 业务错误
	CustomErr *CustomErr `json:"custom_err"`
}
type DecryptInfosItem struct {
	// 业务标识，value为抖音订单号
	AuthId string `json:"auth_id"`
	// 密文
	CipherText string `json:"cipher_text"`
	// 13*******64
	DecryptText string `json:"decrypt_text"`
	// 0
	ErrNo int64 `json:"err_no"`
	// ""
	ErrMsg string `json:"err_msg"`
}
type CustomErr struct {
	// 0
	ErrCode int64 `json:"err_code"`
	// ""
	ErrMsg string `json:"err_msg"`
}
