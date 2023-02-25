package doudian_sdk

type DoudianOpConfig struct {
	AppKey          string
	AppSecret       string
	HttpReadTimeout int64
	OpenRequestUrl  string

	Headers map[string]string
}

func NewDoudianOpConfig() *DoudianOpConfig {
	config := &DoudianOpConfig{
		HttpReadTimeout: 12000, //默认12s超时
		OpenRequestUrl:  "https://openapi-fxg.jinritemai.com",
	}
	return config
}

var GlobalConfig *DoudianOpConfig = NewDoudianOpConfig()
