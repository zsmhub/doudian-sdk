package order_batchDecrypt_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderBatchDecryptResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderBatchDecryptData `json:"data"`
}
type DecryptInfosItem struct {
	// 业务标识；解密抖音官方密文数据使用订单号即可。如果开发者使用【/order/batchEncrypt】加密接口自定义auth_id值，需要传自己定义的标识。
	AuthId string `json:"auth_id"`
	// 密文值，具体内容参考[数据安全加密对接文档](https://op.jinritemai.com/docs/guide-docs/56/589)
	CipherText string `json:"cipher_text"`
	// 解密后明文信息；手机号场景下规则：当is_virtual_tel=fasle时，decrypt_text返回值明文手机号，当is_virtual_tel=true时，并且解密入参是手机号时，decrypt_text返回值为虚拟号主机(phone_no_a)-虚拟号分机号(phone_no_b)的拼接值，返回示例值：”18400913965-7576”
	DecryptText string `json:"decrypt_text"`
	// 错误码
	ErrNo int64 `json:"err_no"`
	// 错误描述
	ErrMsg string `json:"err_msg"`
	// 加密类型 1、 地址加密类型 2、 姓名加密类型 3、 手机号加密类型 4、身份证类型 5、手机号加密类型(不会返回虚拟号)
	DataType int64 `json:"data_type"`
	// 手机号描述；false-真实手机号，true-虚拟手机号
	IsVirtualTel bool `json:"is_virtual_tel"`
	// 虚拟号过期时间（商家自有呼叫中心记录，在到期后还需要则再次请求获取）
	ExpireTime int64 `json:"expire_time"`
	// 虚拟号主机号；当is_virtual_tel=true时，有值返回，当is_virtual_tel=fasle时，返回为空；
	PhoneNoA string `json:"phone_no_a"`
	// 虚拟号分机号；当is_virtual_tel=true时，有值返回，当is_virtual_tel=fasle时，返回为空；
	PhoneNoB string `json:"phone_no_b"`
}
type CustomErr struct {
	// 错误码
	ErrCode int64 `json:"err_code"`
	// 错误内容描述
	ErrMsg string `json:"err_msg"`
}
type OrderBatchDecryptData struct {
	// 解密列表
	DecryptInfos []DecryptInfosItem `json:"decrypt_infos"`
	// 业务错误
	CustomErr *CustomErr `json:"custom_err"`
}
