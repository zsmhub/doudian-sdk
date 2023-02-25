package doudian_sdk

type DoudianOpApiRequest interface {
	GetConfig() *DoudianOpConfig
	SetConfig(config *DoudianOpConfig)
	GetParamObject() interface{}
	GetUrlPath() string
}

type BaseDoudianOpApiRequest struct {
	config *DoudianOpConfig
	client *DoudianOpApiClient
}

func(r *BaseDoudianOpApiRequest) GetConfig() *DoudianOpConfig {
	return r.config
}

func (r *BaseDoudianOpApiRequest) SetConfig(config *DoudianOpConfig) {
	r.config = config
}

func (r *BaseDoudianOpApiRequest) GetClient() *DoudianOpApiClient {
	return r.client
}

func(r *BaseDoudianOpApiRequest) SetClient(client *DoudianOpApiClient) {
	r.client = client
}



