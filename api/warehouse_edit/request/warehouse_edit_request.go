package warehouse_edit_request

import (
	"doudian.com/open/sdk_golang/api/warehouse_edit/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type WarehouseEditRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *WarehouseEditParam 
}
func (c *WarehouseEditRequest) GetUrlPath() string{
	return "/warehouse/edit"
}


func New() *WarehouseEditRequest{
	request := &WarehouseEditRequest{
		Param: &WarehouseEditParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *WarehouseEditRequest) Execute(accessToken *doudian_sdk.AccessToken) (*warehouse_edit_response.WarehouseEditResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &warehouse_edit_response.WarehouseEditResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *WarehouseEditRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *WarehouseEditRequest) GetParams() *WarehouseEditParam{
	return c.Param
}


type WarehouseEditParam struct {
	// 外部仓库id
	OutWarehouseId *string `json:"out_warehouse_id,omitempty"`
	// 仓库名称
	Name string `json:"name,omitempty"`
	// 仓库介绍
	Intro string `json:"intro,omitempty"`
	// 省地址编码
	AddressId1 int64 `json:"address_id1,omitempty"`
	// 市地址编码
	AddressId2 int64 `json:"address_id2,omitempty"`
	// 区地址编码
	AddressId3 int64 `json:"address_id3,omitempty"`
	// 街道地址编码
	AddressId4 int64 `json:"address_id4,omitempty"`
	// 详细地址
	AddressDetail string `json:"address_detail,omitempty"`
}
