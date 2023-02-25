package security_batchReportOrderSecurityEvent_request

import (
	"doudian.com/open/sdk_golang/api/security_batchReportOrderSecurityEvent/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SecurityBatchReportOrderSecurityEventRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SecurityBatchReportOrderSecurityEventParam 
}
func (c *SecurityBatchReportOrderSecurityEventRequest) GetUrlPath() string{
	return "/security/batchReportOrderSecurityEvent"
}


func New() *SecurityBatchReportOrderSecurityEventRequest{
	request := &SecurityBatchReportOrderSecurityEventRequest{
		Param: &SecurityBatchReportOrderSecurityEventParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SecurityBatchReportOrderSecurityEventRequest) Execute(accessToken *doudian_sdk.AccessToken) (*security_batchReportOrderSecurityEvent_response.SecurityBatchReportOrderSecurityEventResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &security_batchReportOrderSecurityEvent_response.SecurityBatchReportOrderSecurityEventResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SecurityBatchReportOrderSecurityEventRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SecurityBatchReportOrderSecurityEventRequest) GetParams() *SecurityBatchReportOrderSecurityEventParam{
	return c.Param
}


type PurchaseProductInfoListItem struct {
	// 采购平台 1：拼多多 2：淘宝 3：1688 4：其它
	PurchasePlatformType int32 `json:"purchase_platform_type,omitempty"`
	// 采购商品链接
	PurchaseProductUrl string `json:"purchase_product_url,omitempty"`
	// 采购人id（外部）
	ExternalPurchaserId string `json:"external_purchaser_id,omitempty"`
	// 本平台订单的商品名称
	ProductName string `json:"product_name,omitempty"`
}
type EventsItem struct {
	// HTTP 请求头里的 doudian-event-id 对应的值
	EventId *string `json:"event_id,omitempty"`
	// 商户的账户ID，每个ISV下需要保证唯一。独立生成的账户唯一标识
	AccountId *string `json:"account_id,omitempty"`
	// main_account 服务商账号体系中的主账号、子账号 main_account/sub_account
	AccountType *string `json:"account_type,omitempty"`
	// orderIds 实际归属的店铺id
	OrderRelatedShopId *string `json:"order_related_shop_id,omitempty"`
	// 选填 ，account_id 关联的店铺 ID 列表
	ShopIds []string `json:"shop_ids,omitempty"`
	// 订单 ids, 单次最大数量 50, 超过 50 需分批上传
	OrderIds *[]string `json:"order_ids,omitempty"`
	// 操作类型.支持操作类型:      1:view_order (查看订单)      2:view_order_list (查看订单列表)      3:download_order (下载订单)      4:download_order_list (下载订单列表)      5:print_order (打印订单)      6:print_order_list (打印订单列表)      7:export_order (导出订单)      8:export_order_list (导出订单列表)      9:delete_order (删除订单)
	OperationType *int32 `json:"operation_type,omitempty"`
	// 精确到秒的操作时间戳，格林威治时间，如1522555200
	OperateTime *string `json:"operate_time,omitempty"`
	// isv请求url
	Url *string `json:"url,omitempty"`
	// 客户端IP，须为用户操作时真实客户端外网IP 若使用SLB，客户端IP添加在HTTP请求的X-Forwarded-For末尾；若使用CWAF，客户端IP放在 X-Real-Ip 属性
	Ip *string `json:"ip,omitempty"`
	// 登录设备的mac地址
	Mac string `json:"mac,omitempty"`
	// 对外发送的订单信息明细格式 场景:商家来调解密接口，同时解密接口里面包含了敏感信息
	IdentifyInfoList *[]IdentifyInfoListItem `json:"identify_info_list,omitempty"`
	// iOS / Android / Windows	商户在什么设备上使用 ISV 的软件
	DeviceType string `json:"device_type,omitempty"`
	// 设备id，标识唯一设备
	DeviceId string `json:"device_id,omitempty"`
	// HTTP 请求头里referer对应的值,参考链接
	Referer *string `json:"referer,omitempty"`
	// HTTP 请求头里的 userAgent对应的值
	UserAgent string `json:"user_agent,omitempty"`
	// 采购商品信息
	PurchaseProductInfoList []PurchaseProductInfoListItem `json:"purchase_product_info_list,omitempty"`
}
type SecurityBatchReportOrderSecurityEventParam struct {
	// 订单事件类型 1:订单访问事件, 2:订单流出事件
	EventType *int32 `json:"event_type,omitempty"`
	// 订单事件列表
	Events *[]EventsItem `json:"events,omitempty"`
}
type IdentifyInfoListItem struct {
	// 代表订单某个信息的字段名 (可直接取抖店开放平台返回的订单信息中的字段名)
	Name *string `json:"name,omitempty"`
	// 代表此字段在发送给第三方时是否为加密状态
	Encrypted *bool `json:"encrypted,omitempty"`
}
