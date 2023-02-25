package logistics_getUserDeliverServiceInfos_response

import (
	"doudian.com/open/sdk_golang/core"
)

type LogisticsGetUserDeliverServiceInfosResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *LogisticsGetUserDeliverServiceInfosData `json:"data"`
}
type UserServiceInfosItem struct {
	// 发货服务信息
	ServiceInfo *ServiceInfo `json:"service_info"`
	// 0：未生效 1：已生效 2：已失效
	EffectStatus int32 `json:"effect_status"`
}
type LogisticsGetUserDeliverServiceInfosData struct {
	// 用户发货服务信息
	UserServiceInfos []UserServiceInfosItem `json:"user_service_infos"`
}
type ServiceInfo struct {
	// GetRecommendedAndDeliveryExpressByOrder
	ServiceCode string `json:"service_code"`
	// GetRecommendedAndDeliveryExpressByOrder
	ServiceName string `json:"service_name"`
}
