package warehouse_create_request

import (
	"doudian.com/open/sdk_golang/api/warehouse_create/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type WarehouseCreateRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *WarehouseCreateParam 
}
func (c *WarehouseCreateRequest) GetUrlPath() string{
	return "/warehouse/create"
}


func New() *WarehouseCreateRequest{
	request := &WarehouseCreateRequest{
		Param: &WarehouseCreateParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *WarehouseCreateRequest) Execute(accessToken *doudian_sdk.AccessToken) (*warehouse_create_response.WarehouseCreateResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &warehouse_create_response.WarehouseCreateResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *WarehouseCreateRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *WarehouseCreateRequest) GetParams() *WarehouseCreateParam{
	return c.Param
}


type WarehouseCreateParam struct {
	// 外部仓库ID，一个店铺下，同一个外部ID只能创建一个仓库
	OutWarehouseId *string `json:"out_warehouse_id,omitempty"`
	// 仓库名称
	Name *string `json:"name,omitempty"`
	// 仓库介绍
	Intro *string `json:"intro,omitempty"`
}
