package address_getAreasByProvince_request

import (
	"doudian.com/open/sdk_golang/api/address_getAreasByProvince/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type AddressGetAreasByProvinceRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *AddressGetAreasByProvinceParam 
}
func (c *AddressGetAreasByProvinceRequest) GetUrlPath() string{
	return "/address/getAreasByProvince"
}


func New() *AddressGetAreasByProvinceRequest{
	request := &AddressGetAreasByProvinceRequest{
		Param: &AddressGetAreasByProvinceParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *AddressGetAreasByProvinceRequest) Execute(accessToken *doudian_sdk.AccessToken) (*address_getAreasByProvince_response.AddressGetAreasByProvinceResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &address_getAreasByProvince_response.AddressGetAreasByProvinceResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *AddressGetAreasByProvinceRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *AddressGetAreasByProvinceRequest) GetParams() *AddressGetAreasByProvinceParam{
	return c.Param
}


type AddressGetAreasByProvinceParam struct {
	// 省ID
	ProvinceId *int64 `json:"province_id,omitempty"`
}
