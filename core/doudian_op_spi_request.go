package doudian_sdk

type DoudianOpSpiParam struct {

	AppKey string

	Timestamp string

	Sign string

	SignV2 string

	SignMethod string

	ParamJson string
}

type DoudianOpSpiRequest interface {
	GetSpiParam() *DoudianOpSpiParam
	GetConfig() *DoudianOpConfig
	GetResponseObject() interface{}
	GetParamJsonObject() interface{}
	GetBizHandler() BizHandler
	Execute() (interface{}, error)
	ResponseJson()(string, error)
	RegisterHandler(bizHandler BizHandler)
}

type BizHandler func(context *DoudianOpSpiContext)

type BaseDoudianOpSpiRequest struct {
	config *DoudianOpConfig
	bizHandler BizHandler
	spiParam *DoudianOpSpiParam
	client DoudianOpSpiClient
}

func (b *BaseDoudianOpSpiRequest) ResponseJson() (string, error) {
	//do nothing
	return "", nil
}

func (b *BaseDoudianOpSpiRequest) Execute() (interface{}, error) {
	//do nothing
	return nil, nil
}

func (b *BaseDoudianOpSpiRequest) GetSpiParam() *DoudianOpSpiParam {
	return b.spiParam
}

func (b *BaseDoudianOpSpiRequest) GetBizHandler() BizHandler {
	return b.bizHandler
}

func (b *BaseDoudianOpSpiRequest) SetSpiParam(spiParam *DoudianOpSpiParam) {
	b.spiParam = spiParam
}

func (b *BaseDoudianOpSpiRequest) RegisterHandler(bizHandler BizHandler) {
	b.bizHandler = bizHandler
}

func (b *BaseDoudianOpSpiRequest) GetConfig() *DoudianOpConfig {
	return b.config
}

func (b *BaseDoudianOpSpiRequest) SetConfig(config *DoudianOpConfig) {
	b.config = config
}

func (b *BaseDoudianOpSpiRequest) SetClient(client DoudianOpSpiClient) {
	b.client = client
}

func (b *BaseDoudianOpSpiRequest) GetClient() DoudianOpSpiClient {
	return b.client
}

func (b *BaseDoudianOpSpiRequest) GetParamJsonObject() interface{} {
	//do nothing
	return nil
}

func (b *BaseDoudianOpSpiRequest) GetResponseObject() interface{} {
	//do nothing
	return nil
}



