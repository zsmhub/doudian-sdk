package logistics_updateOrder_request

import (
	"doudian.com/open/sdk_golang/api/logistics_updateOrder/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type LogisticsUpdateOrderRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *LogisticsUpdateOrderParam 
}
func (c *LogisticsUpdateOrderRequest) GetUrlPath() string{
	return "/logistics/updateOrder"
}


func New() *LogisticsUpdateOrderRequest{
	request := &LogisticsUpdateOrderRequest{
		Param: &LogisticsUpdateOrderParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *LogisticsUpdateOrderRequest) Execute(accessToken *doudian_sdk.AccessToken) (*logistics_updateOrder_response.LogisticsUpdateOrderResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &logistics_updateOrder_response.LogisticsUpdateOrderResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *LogisticsUpdateOrderRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *LogisticsUpdateOrderRequest) GetParams() *LogisticsUpdateOrderParam{
	return c.Param
}


type Address struct {
	// 国家编码（默认CHN，目前只有国内业务）
	CountryCode *string `json:"country_code,omitempty"`
	// 省名称
	ProvinceName *string `json:"province_name,omitempty"`
	// 市名称
	CityName *string `json:"city_name,omitempty"`
	// 区/县名称
	DistrictName *string `json:"district_name,omitempty"`
	// 街道名称。街道名称（street_name）和街道code（street_code），若传入时，需要一起传入。
	StreetName string `json:"street_name,omitempty"`
	// 剩余详细地址，支持密文
	DetailAddress string `json:"detail_address,omitempty"`
	// 省code编码
	ProvinceCode string `json:"province_code,omitempty"`
	// 市code编码
	CityCode string `json:"city_code,omitempty"`
	// 区code编码
	DistrictCode string `json:"district_code,omitempty"`
	// 街道code编码
	StreetCode string `json:"street_code,omitempty"`
}
type ReceiverInfo struct {
	// 收件人地址信息
	Address *Address `json:"address,omitempty"`
	// 收件人联系信息
	Contact *Contact `json:"contact,omitempty"`
}
type ItemsItem struct {
	// 商品名称
	ItemName *string `json:"item_name,omitempty"`
	// 商品规格
	ItemSpecs string `json:"item_specs,omitempty"`
	// 商品数量
	ItemCount *int32 `json:"item_count,omitempty"`
	// 单件商品体积（cm3）
	ItemVolume int32 `json:"item_volume,omitempty"`
	// 单件商品重量（g)
	ItemWeight int32 `json:"item_weight,omitempty"`
	// 单件总净重量（g）
	ItemNetWeight int32 `json:"item_net_weight,omitempty"`
}
type Warehouse struct {
	// 目前该字段无效，统一传false
	IsSumUp *bool `json:"is_sum_up,omitempty"`
	// 仓库订单号(丹鸟等仓发链路使用)
	WhOrderNo string `json:"wh_order_no,omitempty"`
}
type LogisticsUpdateOrderParam struct {
	// 寄件人信息
	SenderInfo *SenderInfo `json:"sender_info,omitempty"`
	// 收件人信息
	ReceiverInfo *ReceiverInfo `json:"receiver_info,omitempty"`
	// 物流服务商编码
	LogisticsCode *string `json:"logistics_code,omitempty"`
	// 运单号
	TrackNo *string `json:"track_no,omitempty"`
	// 商品明细列表
	Items []ItemsItem `json:"items,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty"`
	// 备用扩展字段
	Extra string `json:"extra,omitempty"`
	// 实际使用取号服务店铺user_id
	UserId int64 `json:"user_id,omitempty"`
	// 总体积 货物的总体积或长，宽，高 ；整数 单位cm
	Volume string `json:"volume,omitempty"`
	// /总重量 ；整数 用于与快递商有计抛信任协议的商家）单位克
	Weight int64 `json:"weight,omitempty"`
	// 仓、门店、总对总发货
	Warehouse *Warehouse `json:"warehouse,omitempty"`
}
type Contact struct {
	// 寄件人姓名
	Name string `json:"name,omitempty"`
	// 寄件人固话（和mobile二选一）
	Phone string `json:"phone,omitempty"`
	// 寄件人移动电话（和phone二选一）
	Mobile string `json:"mobile,omitempty"`
}
type SenderInfo struct {
	// 寄件人联系信息
	Contact *Contact `json:"contact,omitempty"`
}
