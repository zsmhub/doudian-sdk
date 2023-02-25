package material_searchMaterial_response

import (
	"doudian.com/open/sdk_golang/core"
)

type MaterialSearchMaterialResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *MaterialSearchMaterialData `json:"data"`
}
type PhotoInfo struct {
	// 图片高度，单位px
	Height int32 `json:"height"`
	// 图片宽度，单位px
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
	// 视频封面地址
	VideoCoverUrl string `json:"video_cover_url"`
	// 视频高度，单位px
	Height int32 `json:"height"`
	// 视频宽度，单位px
	Width int32 `json:"width"`
	// 视频大小，单位B
	Size int64 `json:"size"`
}
type MaterialInfoListItem struct {
	// 素材id
	MaterialId string `json:"material_id"`
	// 父文件夹id
	FolderId string `json:"folder_id"`
	// 创建素材时传入的url
	OriginUrl string `json:"origin_url"`
	// 素材中心返回的url，该字段只对图片有效；如果想获取视频播放信息，请参考video_info/vid字段介绍
	ByteUrl string `json:"byte_url"`
	// 素材名称
	MaterialName string `json:"material_name"`
	// photo-图片 video-视频
	MaterialType string `json:"material_type"`
	// 素材状态，0-待下载 1-有效 4-回收站中 6-已删除
	OperateStatus int32 `json:"operate_status"`
	// 审核状态 1-待审核 2-审核中 3-通过 4-拒绝
	AuditStatus int32 `json:"audit_status"`
	// 审核失败的原因
	AuditRejectDesc string `json:"audit_reject_desc"`
	// 文件大小，单位KB
	Size int64 `json:"size"`
	// 创建素材的时间，格式：yyyy-MM-dd HH:mm:ss
	CreateTime string `json:"create_time"`
	// 最近一次更新素材的时间，格式：yyyy-MM-dd HH:mm:ss
	UpdateTime string `json:"update_time"`
	// 移动到回收站的时间，格式：yyyy-MM-dd HH:mm:ss
	DeleteTime string `json:"delete_time"`
	// 图片信息
	PhotoInfo *PhotoInfo `json:"photoInfo"`
	// 视频信息
	VideoInfo *VideoInfo `json:"videoInfo"`
}
type MaterialSearchMaterialData struct {
	// 素材列表信息
	MaterialInfoList []MaterialInfoListItem `json:"material_info_list"`
	// 总数
	Total int64 `json:"total"`
}
