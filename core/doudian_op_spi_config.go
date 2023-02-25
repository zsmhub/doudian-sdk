package doudian_sdk

import (
	"io"
	"net/url"
)

type SpiConfigurator struct {
	request DoudianOpSpiRequest
	handler BizHandler
	spiParam *DoudianOpSpiParam
}

func (c *SpiConfigurator) ResponseJson() (string, error) {
	return c.request.ResponseJson()
}

func ConfigSpiWithParam(request DoudianOpSpiRequest, handler BizHandler, spiParam *DoudianOpSpiParam) *SpiConfigurator {
	request.RegisterHandler(handler)
	request.GetSpiParam().AppKey = spiParam.AppKey
	request.GetSpiParam().ParamJson = spiParam.ParamJson
	request.GetSpiParam().Sign = spiParam.Sign
	request.GetSpiParam().SignV2 = spiParam.SignV2
	request.GetSpiParam().SignMethod = spiParam.SignMethod
	request.GetSpiParam().Timestamp = spiParam.Timestamp
	return &SpiConfigurator{
		request:  request,
		handler:  handler,
		spiParam: spiParam,
	}
}

func ConfigSpiWithUrlQuery(request DoudianOpSpiRequest, handler BizHandler, queryString string, body io.Reader) *SpiConfigurator {
	queries, _ := url.ParseQuery(queryString)
	appKey := queries.Get("app_key")
	paramJson := queries.Get("param_json")
	if len(paramJson) == 0 {
		bs, _ := io.ReadAll(body)
		paramJson = string(bs)
	}
	sign := queries.Get("sign")
	signV2 := queries.Get("sign_v2")
	signMethod := queries.Get("sign_method")
	timestamp := queries.Get("timestamp")
	return ConfigSpiWithParam(request, handler, &DoudianOpSpiParam{
		AppKey:     appKey,
		Timestamp:  timestamp,
		Sign:       sign,
		SignV2:     signV2,
		SignMethod: signMethod,
		ParamJson:  paramJson,
	})
}
