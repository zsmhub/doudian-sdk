package errors

import "fmt"

type ErroCode int
const(
	ConfigIsNull ErroCode = 10001
	ParamError ErroCode = 10002
	HttpError ErroCode = 10003
	BuildAccessTokenError ErroCode =10004
	ConfigAppKeyIsNull ErroCode =10005
	ConfigAppSecretIsNull ErroCode =10006
	SpiParamIsNull ErroCode = 10007
)

var errMessageMap = map[ErroCode]string{
	ConfigIsNull: "Config配置为空",
	ParamError: "参数错误",
	HttpError:  "处理Http请求错误",
	BuildAccessTokenError: "创建AccessToken错误",
	ConfigAppKeyIsNull:  "Config中AppKey为空",
	ConfigAppSecretIsNull:  "Config中AppSecret为空",
	SpiParamIsNull: "SpiParam参数为空",
}

type DoudianOpError struct {
	Code    ErroCode
	Message string
}

func (err *DoudianOpError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s", err.Code, err.Message)
}

func NewDoudianOpErrorWithMessage(code ErroCode, message string) error {
	return &DoudianOpError{
		Code:    code,
		Message: message,
	}
}

func NewDoudianOpError(code ErroCode) error {
	return &DoudianOpError{
		Code:    code,
		Message: errMessageMap[code],
	}
}