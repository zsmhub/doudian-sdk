package token_create_response

import (
	"doudian.com/open/sdk_golang/core"
)

type TokenCreateResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *TokenCreateData `json:"data"`
}
type TokenCreateData struct {
	// token值。 Tips： 1. 在 access_token 过期前1h之前，ISV使用 refresh_token 刷新时，会返回原来的 access_token 和 refresh_token，但是二者有效期不会变； 2. 在 access_token 过期前1h之内，ISV使用 refresh_token 刷新时，会返回新的 access_token 和 refresh_token，但是原来的 access_token 和 refresh_token 继续有效一个小时； 3. 在 access_token 过期后，ISV使用 refresh_token 刷新时，将获得新的 acces_token 和 refresh_token，同时原来的 acces_token 和 refresh_token 失效；
	AccessToken string `json:"access_token"`
	// 过期时间(秒级时间戳)
	ExpiresIn int64 `json:"expires_in"`
	// 刷新token值。用于刷新access_token的刷新令牌（有效期：14 天）
	RefreshToken string `json:"refresh_token"`
	// 范围
	Scope string `json:"scope"`
	// 店铺ID
	ShopId int64 `json:"shop_id"`
	// 店铺名称
	ShopName string `json:"shop_name"`
	// 授权ID
	AuthorityId string `json:"authority_id"`
	// 授权主体类型
	AuthSubjectType string `json:"auth_subject_type"`
}
