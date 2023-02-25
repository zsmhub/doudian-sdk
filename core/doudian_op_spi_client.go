package doudian_sdk

import (
	"bytes"
	"doudian.com/open/sdk_golang/errors"
	"doudian.com/open/sdk_golang/utils"
	"encoding/json"
	"strings"
)

type DoudianOpSpiClient interface {
	Request(request DoudianOpSpiRequest) (interface{}, error)
}

type DoudianOpSpiClientImpl struct {
}

func NewDefaultDoudianOpSpiClient() DoudianOpSpiClient {
	return &DoudianOpSpiClientImpl{}
}

var DefaultDoudianOpSpiClient = NewDefaultDoudianOpSpiClient()

func (c *DoudianOpSpiClientImpl) Request(request DoudianOpSpiRequest) (interface{}, error) {
	if request.GetConfig() == nil {
		return nil, errors.NewDoudianOpError(errors.ConfigIsNull)
	}
	if request.GetBizHandler() == nil {
		return nil, errors.NewDoudianOpErrorWithMessage(errors.ParamError, "请注册BizHandler(调用RegisterBizHandler方法)")
	}
	if len(request.GetConfig().AppKey) == 0 {
		return nil, errors.NewDoudianOpError(errors.ConfigAppKeyIsNull)
	}
	if len(request.GetConfig().AppSecret) == 0 {
		return nil, errors.NewDoudianOpError(errors.ConfigAppSecretIsNull)
	}
	if request.GetSpiParam() == nil {
		return nil, errors.NewDoudianOpError(errors.SpiParamIsNull)
	}
	if request.GetSpiParam().AppKey != request.GetConfig().AppKey {
		return nil, errors.NewDoudianOpErrorWithMessage(errors.ParamError, "目标app_key和本地app_key不匹配")
	}

	responseObj := request.GetResponseObject()
	response := responseObj.(DoudianOpSpiResponse)
	spiContext := &DoudianOpSpiContext{request: request, response: response}

	appKey := request.GetConfig().AppKey
	appSecret := request.GetConfig().AppSecret
	timestamp := request.GetSpiParam().Timestamp
	signMethod := request.GetSpiParam().SignMethod
	sign := request.GetSpiParam().Sign
	signV2 := request.GetSpiParam().SignV2

	var paramJsonInterfaceMap = make(map[string]interface{})

	decode := json.NewDecoder(strings.NewReader(request.GetSpiParam().ParamJson))
	decode.UseNumber()
	decode.Decode(&paramJsonInterfaceMap)

	sortedParamJsonStr, _ := json.Marshal(paramJsonInterfaceMap)

	calcSign := utils.SpiSign(appKey, appSecret, timestamp, string(sortedParamJsonStr), signMethod)
	if calcSign != sign && calcSign != signV2 {


		bf := bytes.NewBufferString("")
		jsonEncoder := json.NewEncoder(bf)
		jsonEncoder.SetEscapeHTML(false)
		_ = jsonEncoder.Encode(paramJsonInterfaceMap)

		sortedParamJsonWithEscape := strings.TrimSpace(bf.String())

		calcSign = utils.SpiSign(appKey, appSecret, timestamp, sortedParamJsonWithEscape, signMethod)

		if calcSign != sign && calcSign != signV2 {
			spiContext.WrapError(10001, "验证签名失败")
			return response, nil
		}
	}

	request.GetBizHandler()(spiContext)
	return spiContext.response, nil
}

type DoudianOpSpiContext struct {
	request DoudianOpSpiRequest
	response DoudianOpSpiResponse
}

func (c *DoudianOpSpiContext) GetParamObject() interface{} {
	return c.request.GetParamJsonObject()
}

func (c *DoudianOpSpiContext) GetData() interface{} {
	return c.response.GetData()
}

func (c * DoudianOpSpiContext) WrapSuccess() {
	c.response.SetCode(0)
}

func (c *DoudianOpSpiContext) WrapError(code int64, message string) {
	c.response.SetCode(code)
	c.response.SetMessage(message)
}


