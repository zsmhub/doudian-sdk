package logistics_createSFOrder_request

import (
	"doudian.com/open/sdk_golang/api/logistics_createSFOrder/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type LogisticsCreateSFOrderRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *LogisticsCreateSFOrderParam 
}
func (c *LogisticsCreateSFOrderRequest) GetUrlPath() string{
	return "/logistics/createSFOrder"
}


func New() *LogisticsCreateSFOrderRequest{
	request := &LogisticsCreateSFOrderRequest{
		Param: &LogisticsCreateSFOrderParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *LogisticsCreateSFOrderRequest) Execute(accessToken *doudian_sdk.AccessToken) (*logistics_createSFOrder_response.LogisticsCreateSFOrderResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &logistics_createSFOrder_response.LogisticsCreateSFOrderResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *LogisticsCreateSFOrderRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *LogisticsCreateSFOrderRequest) GetParams() *LogisticsCreateSFOrderParam{
	return c.Param
}


type CargoDetailsItem struct {
	// 货物名称，如果需要生成电子运单，则为必填
	Name string `json:"name,omitempty"`
	// 货物单价的币别：参照附录币别代码附件
	Currency string `json:"currency,omitempty"`
	// 货物数量, 跨境件报关需要填写
	Count int64 `json:"count,omitempty"`
	// 货物单位，如：个、台、本，跨境件报关需要填写。
	Unit string `json:"unit,omitempty"`
	// 货物单价，精确到小数点后10位，跨境件报关需要填写
	Amount string `json:"amount,omitempty"`
	// 订单货物单位重量，包含子母件，单位千克，精确到小数点后6位跨境件报关需要填写
	Weight string `json:"weight,omitempty"`
	// 原产地国别，跨境件报关需要填写
	SourceArea string `json:"source_area,omitempty"`
}
type ServiceListItem struct {
	// 增值服务名，如COD等
	Name *string `json:"name,omitempty"`
	// 增值服务扩展属性，参考增值
	Value string `json:"value,omitempty"`
	// 增值服务扩展属性
	Value1 string `json:"value1,omitempty"`
	// 增值服务扩展属性2
	Value2 string `json:"value2,omitempty"`
	// 增值服务扩展属性3
	Value3 string `json:"value3,omitempty"`
	// 增值服务扩展属性4
	Value4 string `json:"value4,omitempty"`
}
type ContactInfoListItem struct {
	// 地址类型： 1，寄件方信息 2，到件方信息
	ContactType *int16 `json:"contact_type,omitempty"`
	// 公司名称
	BizCompany string `json:"biz_company,omitempty"`
	// 联系人
	Contact string `json:"contact,omitempty"`
	// 联系电话（二选一）
	Tel string `json:"tel,omitempty"`
	// 手机
	Mobile string `json:"mobile,omitempty"`
	// 国家或地区 2位代码（默认是CN）
	Country string `json:"country,omitempty"`
	// 所在省级行政区名称，必须是 标准的省级行政区名称如：北 京、广东省、广西壮族自治区 等；此字段影响原寄地代码识 别
	Province *string `json:"province,omitempty"`
	// 所在地级行政区名称，必须是 标准的城市称谓 如：北京市、 深圳市、大理白族自治州等； 此字段影响原寄地代码识别
	City string `json:"city,omitempty"`
	// 所在县/区级行政区名称，必须 是标准的县/区称谓，如：福田 区，南涧彝族自治县、准格尔旗等
	County string `json:"county,omitempty"`
	// 剩余详细地址
	Address string `json:"address,omitempty"`
}
type LogisticsCreateSFOrderParam struct {
	// 订单号；非抖音订单长度仅支持32个字符，格式：数字、大小写字母及“-”、“/”2种分隔符的组合字符串，例如：Doudian-123456
	OrderId *string `json:"order_id,omitempty"`
	// 用于拆包场景：包裹id（只能传入数字、字母和下划线；大小写敏感，即123A,123a 不可当做相同ID，否则存在一定可能取号失败）一单一包裹是不需要传，有2个以上时，从第二个开始都需要传不同id 和parcelQty（子母件）最多二选一填，两者可以都不填
	PackId string `json:"pack_id,omitempty"`
	// 托寄物信息
	CargoDetails *[]CargoDetailsItem `json:"cargo_details,omitempty"`
	// 增值服务信息
	ServiceList []ServiceListItem `json:"service_list,omitempty"`
	// 收寄双方信息（数组长度必须为2）
	ContactInfoList *[]ContactInfoListItem `json:"contact_info_list,omitempty"`
	// 付款方式，支持以下值： 1:寄方付 2:收方付
	PayMethod int16 `json:"pay_method,omitempty"`
	// 快件产品类别，仅 可使用与顺丰销售约定的快件产品
	ExpressTypeId int16 `json:"express_type_id,omitempty"`
	// 子母件场景使用，包裹数，一个包裹对应一个运单号；若包裹数大于1，则返回一个母运单号和N-1个子运单号 和packid（拆包场景）二选一填
	ParcelQty int16 `json:"parcel_qty,omitempty"`
	// 订单货物总重量， 若为子母件必填， 单位千克， 精确到小数点后3 位，如果提供此值， 必须>0 (子母件需>6)
	TotalWeight string `json:"total_weight,omitempty"`
	// 是否返回签回单 （签单返还）的运单号， 支持以下值： 1：要求 0：不要求
	IsSignBack int16 `json:"is_sign_back,omitempty"`
	// 备注
	Remark string `json:"remark,omitempty"`
	// 客户订单货物总长，单位厘米， 精确到小数点后3位， 包含子母件
	TotalLength string `json:"total_length,omitempty"`
	// 客户订单货物总宽，单位厘米， 精确到小数点后3位， 包含子母件
	TotalWidth string `json:"total_width,omitempty"`
	// 客户订单货物总高，单位厘米， 精确到小数点后3位， 包含子母件
	TotalHeight string `json:"total_height,omitempty"`
	// 订单货物总体积，单位立方厘米, 精确到小数点后3位，会用于计抛
	Volume string `json:"volume,omitempty"`
	// 共享账号场景下需传，代表实际使用取号服务的shop_id（需与order_id匹配）；若无法获取到该shop_id，value传值 -1
	UserId int64 `json:"user_id,omitempty"`
	// 订单渠道来源编码，具体请参考[下单渠道来源编码表](https://bytedance.feishu.cn/sheets/shtcngIVwcJlgXLzWhEtKrmv7Af)，当order_id订单号为非抖音订单时必传
	OrderChannel string `json:"order_channel,omitempty"`
}
