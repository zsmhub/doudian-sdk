package warehouse_removeAddr_request

import (
	"doudian.com/open/sdk_golang/api/warehouse_removeAddr/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type WarehouseRemoveAddrRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *WarehouseRemoveAddrParam 
}
func (c *WarehouseRemoveAddrRequest) GetUrlPath() string{
	return "/warehouse/removeAddr"
}


func New() *WarehouseRemoveAddrRequest{
	request := &WarehouseRemoveAddrRequest{
		Param: &WarehouseRemoveAddrParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *WarehouseRemoveAddrRequest) Execute(accessToken *doudian_sdk.AccessToken) (*warehouse_removeAddr_response.WarehouseRemoveAddrResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &warehouse_removeAddr_response.WarehouseRemoveAddrResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *WarehouseRemoveAddrRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *WarehouseRemoveAddrRequest) GetParams() *WarehouseRemoveAddrParam{
	return c.Param
}


type Addr struct {
	// 一级地址id
	AddrId1 *int64 `json:"addr_id1,omitempty"`
	// 二级地址id
	AddrId2 int64 `json:"addr_id2,omitempty"`
	// 三级地址id
	AddrId3 int64 `json:"addr_id3,omitempty"`
	// 四级地址id
	AddrId4 int64 `json:"addr_id4,omitempty"`
}
type WarehouseRemoveAddrParam struct {
	// 外部仓库ID
	OutWarehouseId *string `json:"out_warehouse_id,omitempty"`
	// 删除的地址结构体
	Addr *Addr `json:"addr,omitempty"`
}
