package token_refresh_response

import (
	"doudian.com/open/sdk_golang/core"
)

type TokenRefreshResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *TokenRefreshData `json:"data"`
}
type TokenRefreshData struct {
	// 用于调用API的access_token 过期时间为expires_in值 可通过refresh_token刷新获取新的access_token，过期时间仍为expires_in值
	AccessToken string `json:"access_token"`
	// access_token过期时间；Unix时间戳：秒
	ExpiresIn int64 `json:"expires_in"`
	// 用于刷新access_token的刷新令牌（有效期：14 天）
	RefreshToken string `json:"refresh_token"`
	// 权限范围
	Scope string `json:"scope"`
	// 店铺ID
	ShopId int64 `json:"shop_id"`
	// 店铺名称
	ShopName string `json:"shop_name"`
	// 授权ID,店铺授权为店铺id，达人授权为达人id；
	AuthorityId string `json:"authority_id"`
	// 授权主体类型
	AuthSubjectType string `json:"auth_subject_type"`
}
