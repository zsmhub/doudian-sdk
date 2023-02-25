package material_uploadImageSync_response

import (
	"doudian.com/open/sdk_golang/core"
)

type MaterialUploadImageSyncResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *MaterialUploadImageSyncData `json:"data"`
}
type MaterialUploadImageSyncData struct {
	// 素材id；MaterialId调用【/material/queryMaterialDetail】接口，当audit_status=3时获取byte_url；
	MaterialId string `json:"material_id"`
	// 素材所在文件夹id，0-素材中心的根目录；其他值-表示对应的文件夹id；
	FolderId string `json:"folder_id"`
	// 是否是新建，true-新建
	IsNew bool `json:"is_new"`
	// 素材审核状态; 1-等待审核; 2-审核中; 3-通过; 4-拒绝;注意：只有AuditStatus=3时ByteUrl才会返回；
	AuditStatus int32 `json:"audit_status"`
	// 【已下线】新的URL可使用下列方式获取：方式一：可监听【店铺使用doudian_material_auditResultForShop】或【供应商使用doudian_material_auditResultForBSCP】audit_status=3时可以使用byte_url；方式二：根据响应参数MaterialId调用【/material/queryMaterialDetail】接口，当audit_status=3时获取byte_url；
	ByteUrl string `json:"byte_url"`
}
