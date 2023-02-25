package superm_createVirtualMobile_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SupermCreateVirtualMobileResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SupermCreateVirtualMobileData `json:"data"`
}
type MobileInfo struct {
	// 是否是新创建的虚拟号码
	IsNewCreate bool `json:"is_new_create"`
	// 虚拟号码
	MobileVirtual string `json:"mobile_virtual"`
	// 虚拟号码的过期时间
	ExpireTime string `json:"expire_time"`
}
type SupermCreateVirtualMobileData struct {
	// 获取虚拟号返回结果信息
	MobileInfo *MobileInfo `json:"mobile_info"`
}
