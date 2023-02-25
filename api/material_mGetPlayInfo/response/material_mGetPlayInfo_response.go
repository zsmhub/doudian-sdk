package material_mGetPlayInfo_response

import (
	"doudian.com/open/sdk_golang/core"
)

type MaterialMGetPlayInfoResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *MaterialMGetPlayInfoData `json:"data"`
}
type MaterialMGetPlayInfoData struct {
	// 获取成功map，key为vid，value是视频的播放信息
	SuccessMap map[string]SuccessMapItem `json:"success_map"`
	// 获取失败map，key为vid
	FailedMap map[string]FailedMapItem `json:"failed_map"`
}
type SuccessMapItem struct {
	// vid，素材中心搜索素材或者查素材详情接口返回的vid字段
	Vid string `json:"Vid"`
	// 视频高度
	Height int32 `json:"Height"`
	// 视频宽度
	Width int32 `json:"Width"`
	// 视频格式
	Format string `json:"Format"`
	// 视频大小，单位B
	Size int64 `json:"Size"`
	// 视频原地址
	URI string `json:"URI"`
	// 视频长度(s)
	Duration float64 `json:"Duration"`
	// 视频封面地址
	VideoCoverUrl string `json:"VideoCoverUrl"`
	// 视频播放地址
	MainUrl string `json:"MainUrl"`
	// 视频播放地址和视频封面地址的过期时间
	MainUrlExpire int64 `json:"MainUrlExpire"`
}
type FailedMapItem struct {
	// 错误码
	ErrCode int32 `json:"ErrCode"`
	// 错误信息
	ErrMsg string `json:"ErrMsg"`
}
