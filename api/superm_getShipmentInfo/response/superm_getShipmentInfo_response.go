package superm_getShipmentInfo_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SupermGetShipmentInfoResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SupermGetShipmentInfoData `json:"data"`
}
type ShipmentInfo struct {
	// 店铺父订单号
	ShopOrderId int64 `json:"shop_order_id"`
	// 上门取运力对应售后单号
	AftersaleId int64 `json:"aftersale_id"`
	// 骑手配送运单号
	TrackNo string `json:"track_no"`
	// 抖店店铺ID
	ShopId int64 `json:"shop_id"`
	// 配送状态：    CALLING   = 1 // 骑手呼叫中     RECEIVED  = 2 //骑手已接单     ARRIVED   = 3 //骑手已到取货点     PICKUPPED = 4 //骑手已取货     DELIVERING= 5 //骑手配送中     REJECTED  = 6 //收货人已拒收     RETURNING = 7 //返回中     RETURNED  = 8 //返回完成     DELIVERED = 9 //订单妥投
	ShipmentStatus int32 `json:"shipment_status"`
	// 配送异常：    MERCHANT_DELIVER_SLOW   = 1001  // 商家出货慢,     MERCHANT_DELIVER_ERROR  = 1002  // 商家出错货,     MERCHANT_STOCKOUT       = 1003  // 商家缺货,     GOODS_DAMAGED           = 2001  // 破损,     GOODS_LOST              = 2002  // 丢失,     RIDER_WRONG_GOODS       = 2003  // 骑手取错货,     RIDER_UNDELIVERABLE     = 2004  // 骑手无法送达     WRONG_RECEIPT_CODE      = 3001  // 小票码错误,     MERCHANT_WRONG_ADDRESS  = 3002  // 商家地址错误,     CONSUMER_WRONG_ADDRESS  = 3003  // 消费者地址错误,     WEIGHT_DEVIATION_TOO_LARGE = 3004  //  订单重量偏差过大,     CONSUMER_CALL_NO_ANSWER = 4002  // 消费者电话未接通,     CONSUMER_STAY_OUT       = 4003  // 消费者不在家,     CONSUMER_MODIFY_ADDRESS = 4004  // 消费者线下联系修改地址,     OTHER                   = 5001  // 其他
	ShipmentError int32 `json:"shipment_error"`
	// 配送骑手名称
	RiderName string `json:"rider_name"`
	// 配送骑手联系方式
	RiderPhone string `json:"rider_phone"`
	// 配送骑手经度；地图坐标GCJ02
	RiderLongitude string `json:"rider_longitude"`
	// 配送骑手维度；地图坐标GCJ02
	RiderLatitude string `json:"rider_latitude"`
	// 当前状态的变更时间；时间格式：yyyy-MM-dd HH:mm:ss
	OccurredTime string `json:"occurred_time"`
}
type SupermGetShipmentInfoData struct {
	// 配送状态信息
	ShipmentInfo *ShipmentInfo `json:"shipment_info"`
}
