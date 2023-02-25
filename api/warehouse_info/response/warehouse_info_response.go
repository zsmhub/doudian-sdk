package warehouse_info_response

import (
	"doudian.com/open/sdk_golang/core"
)

type WarehouseInfoResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *Data `json:"data"`
}
type AddrItem struct {
	// 一级地址id
	AddrId1 int64 `json:"addr_id1"`
	// 二级地址id
	AddrId2 int64 `json:"addr_id2"`
	// 三级地址id
	AddrId3 int64 `json:"addr_id3"`
	// 四级地址id
	AddrId4 int64 `json:"addr_id4"`
	// 地址更新时间
	UpdateTime int64 `json:"update_time"`
	// 地址创建时间
	CreateTime int64 `json:"create_time"`
}
type WarehouseLocation struct {
	// 省地址编码
	AddressId1 int64 `json:"address_id1"`
	// 市地址编码
	AddressId2 int64 `json:"address_id2"`
	// 区地址编码
	AddressId3 int64 `json:"address_id3"`
	// 街道地址编码
	AddressId4 int64 `json:"address_id4"`
	// 省地址中文
	AddressName1 string `json:"address_name1"`
	// 市地址中文
	AddressName2 string `json:"address_name2"`
	// 街道地址中文
	AddressName3 string `json:"address_name3"`
	// 街道地址中文
	AddressName4 string `json:"address_name4"`
}
type Data struct {
	// 仓库id
	WarehouseId int64 `json:"warehouse_id"`
	// 仓库外部id
	OutWarehouseId string `json:"out_warehouse_id"`
	// 仓库名称
	Name string `json:"name"`
	// 仓库介绍
	Intro string `json:"intro"`
	// 覆盖区域列表
	Addr []AddrItem `json:"addr"`
	// 更新时间
	UpdateTime int64 `json:"update_time"`
	// 创建时间
	CreateTime int64 `json:"create_time"`
	// 店铺id
	ShopId int64 `json:"shop_id"`
	// 仓库地址
	WarehouseLocation *WarehouseLocation `json:"warehouse_location"`
	// 详细地址
	AddressDetail string `json:"address_detail"`
	// 外部围栏ID列表
	OutFenceIds []string `json:"out_fence_ids"`
}
type WarehouseInfoData struct {
	// 仓库信息
	Data *Data `json:"data"`
}
