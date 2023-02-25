package doudian_sdk

import (
	"doudian.com/open/sdk_golang/errors"
	"doudian.com/open/sdk_golang/utils"
)

type AccessToken struct {
	CreateTokenData
	response *CreateTokenResponse
}

func (at *AccessToken) IsSuccess() bool {
	return at.response.Code == 10000
}

type BuildAccessTokenParam struct {
	Code   string
	ShopId *int64
	Config *DoudianOpConfig

	AuthId          *string
	AuthSubjectType *string
}

type RefreshAccessTokenParam struct {
	RefreshToken string
	Config       *DoudianOpConfig
}

func BuildAccessToken(param *BuildAccessTokenParam) (*AccessToken, error) {
	request := NewCreateTokenRequest()
	if len(param.Code) > 0 {
		request.GetParams().GrantType = "authorization_code"
		request.GetParams().Code = param.Code
	} else {
		request.GetParams().GrantType = "authorization_self"
		request.GetParams().Code = ""
		request.GetParams().ShopId = param.ShopId
		request.GetParams().AuthId = param.AuthId
		request.GetParams().AuthSubjectType = param.AuthSubjectType
	}
	if param.Config != nil {
		request.SetConfig(param.Config)
	}

	response, err := request.Execute(nil)
	if err != nil {
		return nil, err
	}

	accessToken := &AccessToken{
		CreateTokenData: response.Data,
		response:        response,
	}

	if !accessToken.IsSuccess() {
		return nil, errors.NewDoudianOpErrorWithMessage(errors.BuildAccessTokenError, utils.MarshalNoErr(response))
	}
	return accessToken, nil
}

func RefreshAccessToken(param *RefreshAccessTokenParam) (*AccessToken, error) {
	request := NewRefreshTokenRequest()
	request.GetParam().GrantType = "refresh_token"
	request.GetParam().RefreshToken = param.RefreshToken
	if param.Config != nil {
		request.SetConfig(param.Config)
	}

	response, err := request.Execute(nil)
	if err != nil {
		return nil, err
	}
	accessToken := &AccessToken{
		CreateTokenData: response.Data,
		response:        response,
	}

	if !accessToken.IsSuccess() {
		return nil, errors.NewDoudianOpErrorWithMessage(errors.BuildAccessTokenError, utils.MarshalNoErr(response))
	}
	return accessToken, nil
}

func ParseAccessToken(accessTokenStr string) *AccessToken {
	accessToken := &AccessToken{}
	accessToken.AccessToken = accessTokenStr
	return accessToken
}
