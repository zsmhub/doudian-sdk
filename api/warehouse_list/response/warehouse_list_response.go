package warehouse_list_response

import (
	"doudian.com/open/sdk_golang/core"
)

type WarehouseListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *WarehouseListData `json:"data"`
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
	// 更新时间
	UpdateTime int64 `json:"update_time"`
	// 创建时间
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
}
type WarehousesItem struct {
	// 仓库id
	WarehouseId int64 `json:"warehouse_id"`
	// 仓库名字
	Name string `json:"name"`
	// 仓库介绍
	Intro string `json:"intro"`
	// 更新时间
	UpdateTime int64 `json:"update_time"`
	// 创建时间
	CreateTime int64 `json:"create_time"`
	// 仓库外部id
	OutWarehouseId string `json:"out_warehouse_id"`
	// 仓库覆盖范围列表
	Addr []AddrItem `json:"addr"`
	// 店铺id
	ShopId int64 `json:"shop_id"`
	// 外部电子围栏id列表
	OutFenceIds []string `json:"out_fence_ids"`
	// 仓的详细地址
	AddressDetail string `json:"address_detail"`
	// 仓的地址编码
	WarehouseLocation *WarehouseLocation `json:"warehouse_location"`
}
type WarehouseListData struct {
	// 仓库信息
	Warehouses []WarehousesItem `json:"warehouses"`
	// 总数
	Total int64 `json:"total"`
}
