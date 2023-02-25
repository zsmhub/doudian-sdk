package order_getShopAccountItemFile_response

import (
	"doudian.com/open/sdk_golang/core"
)

type OrderGetShopAccountItemFileResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *OrderGetShopAccountItemFileData `json:"data"`
}
type DataItem struct {
	// 店铺id
	ShopId int64 `json:"shop_id"`
	// 账单日期
	BillDate string `json:"bill_date"`
	// 文件url(有效期为1小时)
	Url string `json:"url"`
}
type OrderGetShopAccountItemFileData struct {
	// data
	Data []DataItem `json:"data"`
	// 返回值
	Code string `json:"code"`
	// 返回参数
	CodeMsg string `json:"code_msg"`
}
