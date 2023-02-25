package logistics_getShopKey_response

import (
	"doudian.com/open/sdk_golang/core"
)

type LogisticsGetShopKeyResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *LogisticsGetShopKeyData `json:"data"`
}
type LogisticsGetShopKeyData struct {
	// 公钥加密后的对称密钥，用于解密电子面单密文
	Key string `json:"key"`
}
