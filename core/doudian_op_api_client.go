package doudian_sdk

import (
	"doudian.com/open/sdk_golang/core/http"
	"doudian.com/open/sdk_golang/errors"
	"doudian.com/open/sdk_golang/utils"
	"fmt"
	"strings"
	"time"
)

type DoudianOpApiClient struct{}

func NewDoudianOpApiClient() *DoudianOpApiClient {
	return &DoudianOpApiClient{}
}

var DefaultDoudianOpApiClient *DoudianOpApiClient = NewDoudianOpApiClient()

func (client *DoudianOpApiClient) Request(request DoudianOpApiRequest, accessToken *AccessToken) (string, error) {
	if request.GetConfig() == nil {
		return "", errors.NewDoudianOpError(errors.ConfigIsNull)
	}

	appKey := request.GetConfig().AppKey
	appSecret := request.GetConfig().AppSecret
	paramJson := request.GetParamObject()
	if len(appKey) == 0 {
		return "", errors.NewDoudianOpErrorWithMessage(errors.ParamError, "app key为空")
	}
	if len(appSecret) == 0 {
		return "", errors.NewDoudianOpErrorWithMessage(errors.ParamError, "app secret为空")
	}
	urlPath := request.GetUrlPath()
	method := urlPathToMethod(urlPath)
	paramJsonString := utils.Marshal(paramJson)
	timestamp := time.Now().Unix()
	sign := utils.Sign(appKey, appSecret, method, timestamp, paramJsonString)

	httpParamsMap := map[string]string{
		"app_key":     appKey,
		"method":      method,
		"timestamp":   fmt.Sprintf("%d", timestamp),
		"sign":        sign,
		"v":           "2",
		"sign_method": "hmac-sha256",
	}
	if accessToken != nil {
		httpParamsMap["access_token"] = accessToken.AccessToken
	}
	//生成logId
	//logId := utils.GenLogID()
	httpHeaderMap := map[string]string{
		"from":     "sdk",
		"sdk-type": "golang",
		//"x-tt-logid": logId,
		"sdk-version":            DoudianSdkVersion,
		"x-open-no-old-err-code": "1",
	}
	if request.GetConfig() != nil {
		for k, v := range request.GetConfig().Headers {
			httpHeaderMap[k] = v
		}
	}

	timeout := request.GetConfig().HttpReadTimeout
	if timeout == 0 {
		timeout = 12000 //默认12s超时
	}
	httpRequest := &http.DoudianOpHttpRequest{
		Url:     fmt.Sprintf("%s/%s", request.GetConfig().OpenRequestUrl, urlPath),
		Params:  httpParamsMap,
		Headers: httpHeaderMap,
		Body:    paramJsonString,
	}

	//log.Printf("[doudian sdk] logid: %s, method: %s, httpRequest: %s\n", logId, method, utils.MarshalNoErr(httpRequest))

	httpResponse, err := http.GetHttpClient(timeout).Post(httpRequest)
	if err != nil {
		return "", err
	}

	return httpResponse.Body, nil
}

func urlPathToMethod(urlPath string) string {
	if len(urlPath) == 0 {
		return urlPath
	}
	if urlPath[0] == '/' {
		urlPath = urlPath[1:]
	}
	return strings.ReplaceAll(urlPath, "/", ".")
}
