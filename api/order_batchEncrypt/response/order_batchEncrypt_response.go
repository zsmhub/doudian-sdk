package order_batchEncrypt_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderBatchEncryptResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderBatchEncryptData `json:"data"`
}
type EncryptInfosItem struct {
	// 鉴权id
	AuthId string `json:"auth_id"`
	// 密文
	CipherText string `json:"cipher_text"`
	// 明文
	DecryptText string `json:"decrypt_text"`
	// 错误码
	ErrNo int64 `json:"err_no"`
	// 错误描述
	ErrMsg string `json:"err_msg"`
}
type CustomErr struct {
	// 错误码
	ErrCode int64 `json:"err_code"`
	// 错误描述
	ErrMsg string `json:"err_msg"`
}
type OrderBatchEncryptData struct {
	// 加密之后的数据
	EncryptInfos []EncryptInfosItem `json:"encrypt_infos"`
	// 业务错误
	CustomErr *CustomErr `json:"custom_err"`
}
