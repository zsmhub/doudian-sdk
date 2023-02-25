package material_getFolderInfo_response

import (
	"doudian.com/open/sdk_golang/core"
)

type MaterialGetFolderInfoResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *MaterialGetFolderInfoData `json:"data"`
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
	// 视频时长
	Duration float64 `json:"duration"`
	// 视频云vid，根据该字段获取视频播放信息等https://op.jinritemai.com/docs/api-docs/69/2164
	Vid string `json:"vid"`
	// 视频封面
	VideoCoverUrl string `json:"video_cover_url"`
}
type ChildMaterialItem struct {
	// 素材id
	MaterialId string `json:"material_id"`
	// 文件夹id
	FolderId string `json:"folder_id"`
	// 开发者创建素材是传入的公网可访问url
	OriginUrl string `json:"origin_url"`
	// 图片url，当audit_status=3时才会返回byte_url；视频请使用vid字段；
	ByteUrl string `json:"byte_url"`
	// 素材名称
	MaterilName string `json:"materil_name"`
	// 素材类型，photo-图片；video-视频；
	MaterialType string `json:"material_type"`
	// 状态，0-待下载；1-有效；4-在回收站中；
	OperateStatus int32 `json:"operate_status"`
	// 审核状态，1-待审核；2-审核中；3-通过 4-拒绝；注意：只有audit_status =3时byte_url才会返回；
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
type FolderInfo struct {
	// 文件夹id
	FolderId string `json:"folder_id"`
	// 文件夹类型
	FolderType int32 `json:"folder_type"`
	// 文件夹名称
	FolderName string `json:"folder_name"`
	// 文件夹状态，1-有效 4-在回收站中
	OperateStatus int32 `json:"operate_status"`
	// 子文件夹列表
	ChildFolder []ChildFolderItem `json:"child_folder"`
	// 素材列表
	ChildMaterial []ChildMaterialItem `json:"child_material"`
	// 父文件夹id，0-素材中心根目录
	ParentFolderId string `json:"parent_folder_id"`
	// 文件夹创建时间
	CreateTime string `json:"create_time"`
	// 文件夹最近一次修改时间
	UpdateTime string `json:"update_time"`
	// 文件夹移动到回收站的时间，只有在回收站中，该字段才有意义
	DeleteTime string `json:"delete_time"`
	// 文件夹下素材总数目
	TotalChildMaterialNum int64 `json:"total_child_material_num"`
	// 文件夹属性，0-文件夹（不限素材类型），1-图片文件夹（只能上传图片），2-视频文件夹（只能上传视频）
	FolderAttr int32 `json:"folder_attr"`
}
type MaterialGetFolderInfoData struct {
	// 文件夹信息
	FolderInfo *FolderInfo `json:"folder_info"`
}
type ChildFolderItem struct {
	// 文件夹id
	FolderId string `json:"folder_id"`
	// 文件夹类型
	FolderType int32 `json:"folder_type"`
	// 文件夹名称
	FolderName string `json:"folder_name"`
	// 文件夹状态，1-有效 4-在回收站中
	OperateStatus int32 `json:"operate_status"`
	// 父文件夹id
	ParentFolderId string `json:"parent_folder_id"`
	// 文件夹创建时间
	CreateTime string `json:"create_time"`
	// 文件夹最近一次修改时间
	UpdateTime string `json:"update_time"`
	// 文件夹移动到回收站的时间，只有在回收站中，该字段才有意义
	DeleteTime string `json:"delete_time"`
	// 文件夹属性，0-文件夹（不限素材类型），1-图片文件夹（只能上传图片），2-视频文件夹（只能上传视频）
	FolderAttr int32 `json:"folder_attr"`
}
