package logistics_listShopNetsite_response

import (
	"doudian.com/open/sdk_golang/core"
)

type LogisticsListShopNetsiteResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *LogisticsListShopNetsiteData `json:"data"`
}
type SenderAddressItem struct {
	// 省名称
	ProvinceName string `json:"province_name"`
	// 市名称
	CityName string `json:"city_name"`
	// 区/县名称
	DistrictName string `json:"district_name"`
	// 街道名称
	StreetName string `json:"street_name"`
	// 剩余详细地址
	DetailAddress string `json:"detail_address"`
}
type NetsitesItem struct {
	// 网点Code
	NetsiteCode string `json:"netsite_code"`
	// 网点名称
	NetsiteName string `json:"netsite_name"`
	// 电子面单余额数量
	Amount string `json:"amount"`
	// 寄件人地址
	SenderAddress []SenderAddressItem `json:"sender_address"`
	// 已取单号数量，若业务本身无值，则传-1，前端可展示为“-”
	AllocatedQuantity int64 `json:"allocated_quantity"`
	// 已取消单号数量，若业务本身无值，则传-1，前端可展示为“-”
	CancelledQuantity int64 `json:"cancelled_quantity"`
	// 已回收单号数量，若业务本身无值，则传-1，前端可展示为“-”
	RecycledQuantity int64 `json:"recycled_quantity"`
	// 快递公司编码
	Company string `json:"company"`
	// 物流服务商业务类型 1：直营  2：加盟 3：落地配 4：直营带网点
	CompanyType int16 `json:"company_type"`
}
type LogisticsListShopNetsiteData struct {
	// 商家已开通的网点列表信息
	Netsites []NetsitesItem `json:"netsites"`
	// 快递公司编码
	LogisticsCode string `json:"logistics_code"`
	// 物流服务商业务类型 1：直营  2：加盟 3：落地配 4：直营带网点
	CompanyType int16 `json:"company_type"`
}
