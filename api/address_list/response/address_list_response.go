package address_list_response

import (
	"doudian.com/open/sdk_golang/core"
)

type AddressListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *AddressListData `json:"data"`
}
type AddressListItem struct {
	// 地址库id
	AddressId int64 `json:"address_id"`
	// 收/发件人
	RecieverName string `json:"reciever_name"`
	// 是否为退货默认
	IsDefault int64 `json:"is_default"`
	// 是否为发货默认
	IsSendDefault int64 `json:"is_send_default"`
	// 创建时间，时间戳，秒
	CreateTime string `json:"create_time"`
	// 更新时间，时间戳，秒
	UpdateTime string `json:"update_time"`
	// 收件人省份
	ReceiverProvinc string `json:"receiver_provinc"`
	// 收件人城市
	ReceiverCity string `json:"receiver_city"`
	// 收件人地区
	ReceiverDistrict string `json:"receiver_district"`
	// 收件人详情地址
	ReceiverDetail string `json:"receiver_detail"`
	// 收件人街道
	ReceiverStreet string `json:"receiver_street"`
	// 地址备注信息
	Remark string `json:"remark"`
}
type AddressListData struct {
	// 地址总数
	Total int64 `json:"total"`
	// 页码
	PageSize int64 `json:"page_size"`
	// 每页数量
	PageNo int64 `json:"page_no"`
	// 地址列表
	AddressList []AddressListItem `json:"address_list"`
}
