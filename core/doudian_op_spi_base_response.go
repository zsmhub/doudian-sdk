package doudian_sdk

type DoudianOpSpiResponse interface {
	GetData() interface{}
	SetCode(code int64)
	SetMessage(message string)
}

type BaseDoudianOpSpiResponse struct {
	Code int64 `json:"code"`
	Message string `json:"message"`
}

func (b *BaseDoudianOpSpiResponse) SetCode(code int64) {
	b.Code = code
}

func (b *BaseDoudianOpSpiResponse) SetMessage(message string) {
	b.Message = message
}



