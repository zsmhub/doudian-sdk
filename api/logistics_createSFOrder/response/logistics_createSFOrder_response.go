package logistics_createSFOrder_response

import (
	"doudian.com/open/sdk_golang/core"
)

type LogisticsCreateSFOrderResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *LogisticsCreateSFOrderData `json:"data"`
}
type WaybillNoInfoListItem struct {
	// 运单号类型: 1：母单 2 :子单 3 : 签回单
	WaybillType int16 `json:"waybill_type"`
	// 运单号
	TrackNo string `json:"track_no"`
}
type RouteLabelData struct {
	// 运单号
	TrackNo string `json:"track_no"`
	// 原寄地中转场
	SourceTransferCode string `json:"source_transfer_code"`
	// 原寄地城市代码
	SourceCityCode string `json:"source_city_code"`
	// 原寄地网点代码
	SourceDeptCode string `json:"source_dept_code"`
	// 原寄地单元区域
	SourceTeamCode string `json:"sourceTeamCode"`
	// 目的地城市代码, eg:755
	DestCityCode string `json:"destCityCode"`
	// 目的地网点代码, eg:755AQ
	DestDeptCode string `json:"destDeptCode"`
	// 目的地网点代码映射码
	DestDeptCodeMapping string `json:"dest_dept_code_mapping"`
	// 目的地单元区域, eg:001
	DestTeamCode string `json:"dest_team_code"`
	// 目的地单元区域映射码
	DestTeamCodeMapping string `json:"dest_team_code_mapping"`
	// 目的地中转场
	DestTransferCode string `json:"dest_transfer_code"`
	// 打单时的路由标签信息如果 是大网的路由标签,这里的 值是目的地网点代码,如果 是同城配的路由标签,这里 的值是根据同城配的设置映 射出来的值,不同的配置结 果会不一样,不能根据-符 号切分(如:上海同城配,可能 是:集散点-目的地网点-接 驳点,也有可能是目的地网 点代码-集散点-接驳点)
	DestRouteLabel string `json:"dest_route_label"`
	// 产品名称 对应RLS:pro_name
	ProName string `json:"pro_name"`
	// 快件内容: 如:C816、SP601
	CargoTypeCode string `json:"cargo_type_code"`
	// 时效代码, 如:T4
	LimitTypeCode string `json:"limit_type_code"`
	// 产品类型,如:B1
	ExpressTypeCode string `json:"express_type_code"`
	// 入港映射码 eg:S10
	CodingMapping string `json:"coding_mapping"`
	// 出港映射码
	CodingMappingOut string `json:"coding_mapping_out"`
	// XB标志 0:不需要打印XB 1:需要打印XB
	XbFlag string `json:"xb_flag"`
	// 打印标志 返回值总共有9位,每位只 有0和1两种,0表示按丰密 面单默认的规则,1是显示, 顺序如下,如111110000 表示打印寄方姓名、寄方 电话、寄方公司名、寄方 地址和重量,收方姓名、收 方电话、收方公司和收方 地址按丰密面单默认规则 1:寄方姓名 2:寄方电话 3:寄方公司名 4:寄方地址 5:重量 6:收方姓名 7:收方电话 8:收方公司名 9:收方地址
	PrintFlag string `json:"print_flag"`
	// 二维码 根据规则生成字符串信息, 格式为MMM={"k1":"(目的 地中转场代码)","k2":"(目的 地原始网点代码)","k3":"(目 的地单元区域)","k4":"(附件 通过三维码(express_type_code、 limit_type_code、 cargo_type_code)映射时效类型)","k5":"(运单 号)","k6":"(AB标识)","k7":"( 校验码)"} （ISV自身做逻辑展示）
	TwoDimensionCode string `json:"two_dimension_code"`
	// 时效类型: 值为二维码中的K4
	ProCode string `json:"pro_code"`
	// 打印图标,根据托寄物判断需 要打印的图标(重货,蟹类,生鲜,易碎，Z标) 返回值有8位，每一位只有0和1两种， 0表示按运单默认的规则， 1表示显示。后面两位默认0备用。 顺序如下：重货,蟹类,生鲜,易碎,医药类,Z标,0,0 如：00000000表示不需要打印重货，蟹类，生鲜，易碎 ,医药,Z标,备用,备用
	PrintIcon string `json:"print_icon"`
	// AB标
	AbFlag string `json:"ab_flag"`
	// 查询出现异常时返回信息。 返回代码: 0 系统异常 1 未找到面单
	ErrMsg string `json:"err_msg"`
}
type RouteLabelInfo struct {
	// 返回调用结果，1000：调用成功；其他调用失败
	Code string `json:"code"`
	// 路由标签数据详细数据
	RouteLabelData *RouteLabelData `json:"route_label_data"`
	// 失败异常描述
	Message string `json:"message"`
}
type LogisticsCreateSFOrderData struct {
	// 订单号
	OrderId string `json:"order_id"`
	// 包裹id
	PackId string `json:"pack_id"`
	// 原寄地区域代码，可用于顺丰 电子运单标签打印
	OriginCode string `json:"origin_code"`
	// 目的地区域代码，可用于顺丰 电子运单标签打印
	DestCode string `json:"dest_code"`
	// 筛单结果： 1：人工确认 2：可收派 3：不可以收派
	FilterResult int16 `json:"filter_result"`
	// 如果filter_result=3时为必填， 不可以收派的原因代码： 1：收方超范围 2：派方超范围 3：其它原因 高峰管控提示信息 【数字】：【高峰管控提示信息】
	Remark string `json:"remark"`
	// 顺丰运单号
	WaybillNoInfoList []WaybillNoInfoListItem `json:"waybill_no_info_list"`
	// 路由标签
	RouteLabelInfo *RouteLabelInfo `json:"route_label_info"`
	// 2;具体请看文档映射表
	OrderChannel string `json:"order_channel"`
}
