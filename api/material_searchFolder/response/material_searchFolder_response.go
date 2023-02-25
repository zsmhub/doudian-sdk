package material_searchFolder_response

import (
	"doudian.com/open/sdk_golang/core"
)

type MaterialSearchFolderResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *MaterialSearchFolderData `json:"data"`
}
type FolderInfoListItem struct {
	// 文件夹id
	FolderId string `json:"folder_id"`
	// 文件夹类型 0-用户创建 1-默认 2-系统文件夹
	FolderType int32 `json:"folder_type"`
	// 文件夹名称
	Name string `json:"name"`
	// 文件夹状态。1-有效 4-在回收站中
	OperateStatus int32 `json:"operate_status"`
	// 父文件夹id
	ParentFolderId string `json:"parent_folder_id"`
	// 文件夹创建时间
	CreateTime string `json:"create_time"`
	// 文件夹最近一次编辑时间
	UpdateTime string `json:"update_time"`
	// 文件夹删除时间。当文件夹在回收站时返回，未删除时为""
	DeleteTime string `json:"delete_time"`
	// 文件夹属性，0-文件夹（不限素材类型），1-图片文件夹（只能上传图片），2-视频文件夹（只能上传视频）
	FolderAttr string `json:"folder_attr"`
}
type MaterialSearchFolderData struct {
	// 文件夹列表
	FolderInfoList []FolderInfoListItem `json:"folder_info_list"`
	// 符合筛选条件的文件夹数量
	Total int64 `json:"total"`
}
