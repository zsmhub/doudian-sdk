package warehouse_createBatch_request

import (
	"doudian.com/open/sdk_golang/api/warehouse_createBatch/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type WarehouseCreateBatchRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *WarehouseCreateBatchParam 
}
func (c *WarehouseCreateBatchRequest) GetUrlPath() string{
	return "/warehouse/createBatch"
}


func New() *WarehouseCreateBatchRequest{
	request := &WarehouseCreateBatchRequest{
		Param: &WarehouseCreateBatchParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *WarehouseCreateBatchRequest) Execute(accessToken *doudian_sdk.AccessToken) (*warehouse_createBatch_response.WarehouseCreateBatchResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &warehouse_createBatch_response.WarehouseCreateBatchResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *WarehouseCreateBatchRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *WarehouseCreateBatchRequest) GetParams() *WarehouseCreateBatchParam{
	return c.Param
}


type WarehouseLocation struct {
	// 省地址编码
	AddressId1 int64 `json:"address_id1,omitempty"`
	// 市地址编码
	AddressId2 int64 `json:"address_id2,omitempty"`
	// 区地址编码
	AddressId3 int64 `json:"address_id3,omitempty"`
	// 街道地址编码
	AddressId4 int64 `json:"address_id4,omitempty"`
}
type OutWarehouseListItem struct {
	// 外部仓库ID，一个店铺下，同一个外部ID只能创建一个仓库
	OutWarehouseId *string `json:"out_warehouse_id,omitempty"`
	// 仓库名称
	Name *string `json:"name,omitempty"`
	// 仓库介绍
	Intro *string `json:"intro,omitempty"`
	// 仓地址详细地址，苹果门店店铺、超市小时达店铺建仓必传
	AddressDetail string `json:"address_detail,omitempty"`
	// 仓地址编码  苹果门店店铺、超市小时达店铺建仓必传
	WarehouseLocation *WarehouseLocation `json:"warehouse_location,omitempty"`
}
type WarehouseCreateBatchParam struct {
	// 外部仓信息列表
	OutWarehouseList *[]OutWarehouseListItem `json:"out_warehouse_list,omitempty"`
}
