package material_queryMaterialDetail_response

import (
	"doudian.com/open/sdk_golang/core"
)

type MaterialQueryMaterialDetailResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *MaterialQueryMaterialDetailData `json:"data"`
}
type MaterialInfo struct {
	// 素材id；可使用【/material/batchUploadImageSync】【/material/uploadImageSync】【/material/searchMaterial  】接口获取
	MaterialId string `json:"material_id"`
	// 文件夹id
	FolderId string `json:"folder_id"`
	// 开发者上传时传入的公网可访问地址
	OriginUrl string `json:"origin_url"`
	// 图片url；当audit_status=3时获取byte_url；有值返回；
	ByteUrl string `json:"byte_url"`
	// 素材名称
	MaterilName string `json:"materil_name"`
	// 素材类型，photo-图片；video-视频；
	MaterialType string `json:"material_type"`
	// 素材状态，0-待下载；1-有效； 4-在回收站中；
	OperateStatus int32 `json:"operate_status"`
	// 审核状态，1-待审核 2-审核中 3-通过 4-拒绝；注意：只有AuditStatus=3时ByteUrl才会返回；
	AuditStatus int32 `json:"audit_status"`
	// 审核失败的原因
	AuditRejectDesc string `json:"audit_reject_desc"`
	// 大小，单位为byte
	Size int64 `json:"size"`
	// 图片信息
	PhotoInfo *PhotoInfo `json:"photo_info"`
	// 视频信息
	VideoInfo *VideoInfo `json:"video_info"`
	// 素材创建时间
	CreateTime string `json:"create_time"`
	// 素材最近一次修改时间
	UpdateTime string `json:"update_time"`
	// 素材移动到回收站的时间，只有在回收站中，该字段才有意义
	DeleteTime string `json:"delete_time"`
}
type MaterialQueryMaterialDetailData struct {
	// 素材详情
	MaterialInfo *MaterialInfo `json:"material_info"`
}
type PhotoInfo struct {
	// 图片高度
	Height int32 `json:"height"`
	// 图片宽度
	Width int32 `json:"width"`
	// 图片格式
	Format string `json:"format"`
}
type VideoInfo struct {
	// 视频格式
	Format string `json:"format"`
	// 视频时长，单位秒
	Duration float64 `json:"duration"`
	// vid，用于获取视频播放地址，接口文档见：https://op.jinritemai.com/docs/api-docs/69/2164
	Vid string `json:"vid"`
	// 视频高度
	Height int32 `json:"height"`
	// 视频宽度
	Width int32 `json:"width"`
	// 视频封面地址
	VideoCoverUrl string `json:"video_cover_url"`
}
