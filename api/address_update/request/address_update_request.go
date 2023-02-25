package address_update_request

import (
	"doudian.com/open/sdk_golang/api/address_update/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type AddressUpdateRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *AddressUpdateParam 
}
func (c *AddressUpdateRequest) GetUrlPath() string{
	return "/address/update"
}


func New() *AddressUpdateRequest{
	request := &AddressUpdateRequest{
		Param: &AddressUpdateParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *AddressUpdateRequest) Execute(accessToken *doudian_sdk.AccessToken) (*address_update_response.AddressUpdateResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &address_update_response.AddressUpdateResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *AddressUpdateRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *AddressUpdateRequest) GetParams() *AddressUpdateParam{
	return c.Param
}


type AddressUpdateParam struct {
	// 地址信息
	Address *Address `json:"address,omitempty"`
	// 门店场景需要传门店ID，其他场景默认值0即可
	StoreId int64 `json:"store_id,omitempty"`
}
type Address struct {
	// 地址ID
	Id *int64 `json:"id,omitempty"`
	// 联系人姓名
	UserName *string `json:"user_name,omitempty"`
	// 手机号码
	Mobile string `json:"mobile,omitempty"`
	// 省ID
	ProvinceId *int64 `json:"province_id,omitempty"`
	// 城市ID
	CityId *int64 `json:"city_id,omitempty"`
	// 区ID
	TownId *int64 `json:"town_id,omitempty"`
	// 详细地址
	Detail *string `json:"detail,omitempty"`
	// 街道ID
	StreetId *int64 `json:"street_id,omitempty"`
	// 联系方式类型(0-手机,1-普通座机,2-企业座机)
	LinkType *int32 `json:"link_type,omitempty"`
	// 普通座机格式：区号-座机号-分机号(分机号选填)、区号3~4位、座机号7~8位、分机号不超过5位。企业座机：400/800开头不超过10位、95开头在5~8we
	FixedPhone string `json:"fixed_phone,omitempty"`
	// 售后备注，限制200个字符
	Remark string `json:"remark,omitempty"`
}
