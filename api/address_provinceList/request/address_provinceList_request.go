package address_provinceList_request

import (
	"doudian.com/open/sdk_golang/api/address_provinceList/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type AddressProvinceListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *AddressProvinceListParam 
}
func (c *AddressProvinceListRequest) GetUrlPath() string{
	return "/address/provinceList"
}


func New() *AddressProvinceListRequest{
	request := &AddressProvinceListRequest{
		Param: &AddressProvinceListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *AddressProvinceListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*address_provinceList_response.AddressProvinceListResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &address_provinceList_response.AddressProvinceListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *AddressProvinceListRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *AddressProvinceListRequest) GetParams() *AddressProvinceListParam{
	return c.Param
}


type AddressProvinceListParam struct {
}
