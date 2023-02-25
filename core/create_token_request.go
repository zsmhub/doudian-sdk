package doudian_sdk

import (
	"encoding/json"
)

type CreateTokenRequest struct {
	BaseDoudianOpApiRequest
	param *CreateTokenParam
}

func (r *CreateTokenRequest) GetParamObject() interface{} {
	return r.param
}

func (r *CreateTokenRequest) GetParams() *CreateTokenParam {
	return r.param
}

func (r *CreateTokenRequest) Execute(accessToken *AccessToken) (*CreateTokenResponse, error) {
	responseJson, err := r.GetClient().Request(r, accessToken)
	if err != nil {
		return nil, err
	}
	response := &CreateTokenResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil
}

func (r *CreateTokenRequest) GetUrlPath() string {
	return "token/create"
}

func NewCreateTokenRequest() *CreateTokenRequest {
	request := &CreateTokenRequest{
		param: &CreateTokenParam{},
	}
	request.SetConfig(GlobalConfig)
	request.SetClient(DefaultDoudianOpApiClient)
	return request
}

type CreateTokenResponse struct {
	BaseDoudianOpApiResponse
	Data CreateTokenData `json:"data"`
}

type CreateTokenParam struct {
	Code string `json:"code"`

	GrantType string `json:"grant_type"`

	ShopId *int64 `json:"shop_id"`

	AuthId *string `json:"auth_id"`

	AuthSubjectType *string `json:"auth_subject_type"`
}

type CreateTokenData struct {

	//用于调用API的access_token
	//过期时间为expires_in值
	//可通过refresh_token刷新获取新的access_token，过期时间仍为expires_in值
	AccessToken string `json:"access_token"`

	//access_token接口调用凭证超时时间，单位（秒），默认有效期：7天
	ExpiresIn int64 `json:"expires_in"`

	//授权作用域，使用逗号,分隔。预留字段
	Scope string `json:"scope"`

	//店铺ID
	ShopId string `json:"shop_id"`

	//店铺名称
	ShopName string `json:"shop_name"`

	//用于刷新access_token的刷新令牌（有效期：14 天）
	RefreshToken string `json:"refresh_token"`

	//授权主体id
	AuthorityId string `json:"authority_id"`

	//授权主体类型
	AuthSubjectType string `json:"auth_subject_type"`
}
