package doudian_message

type DoudianOpConfig struct {
	AppKey    string
	AppSecret string
}

func NewDoudianOpConfig() *DoudianOpConfig {
	config := &DoudianOpConfig{}
	return config
}

var GlobalConfig = NewDoudianOpConfig()
