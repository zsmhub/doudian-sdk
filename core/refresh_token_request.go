package doudian_sdk

import "encoding/json"

type RefreshTokenRequest struct {
	BaseDoudianOpApiRequest
	param *RefreshTokenParam
}

func (r *RefreshTokenRequest) GetParamObject() interface{}{
	return r.param
}

func(r *RefreshTokenRequest) GetParam() *RefreshTokenParam {
	return r.param
}

func (r *RefreshTokenRequest) Execute(accessToken *AccessToken) (*CreateTokenResponse, error) {
	responseJson, err := r.GetClient().Request(r, accessToken)
	if err != nil {
		return nil, err
	}
	response := &CreateTokenResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil
}

func (r *RefreshTokenRequest) GetUrlPath() string {
	return "token/refresh"
}

func NewRefreshTokenRequest() *RefreshTokenRequest {
	request := &RefreshTokenRequest{
		param:  &RefreshTokenParam{},
	}
	request.SetConfig(GlobalConfig)
	request.SetClient(DefaultDoudianOpApiClient)
	return request
}

type RefreshTokenParam struct {

	GrantType string `json:"grant_type"`

	RefreshToken string `json:"refresh_token"`
}