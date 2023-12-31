package errcode

import "fmt"

type ErrorCode struct {
	code    int32
	msg     string
	details []string
}

func NewError(code int32, msg string) *ErrorCode {
	return &ErrorCode{code: code, msg: msg}
}

func (e *ErrorCode) GetCode() int32 {
	return e.code
}

func (e *ErrorCode) GetMessage() string {
	return e.msg
}

func (e *ErrorCode) GetDetails() []string {
	return e.details
}

func (e *ErrorCode) WithDetails(details ...string) *ErrorCode {
	newError := *e
	// 切片是引用类型，需要 copy 一份
	// 注意，不推荐这种方式，底层都会引用同一个数组
	// newError.details = details[:]
	// 常规做法是用 append 将一个切片追加到另一个切片中
	// 还有一种更简单的写法
	newError.details = append([]string(nil), details...)
	// 还有另一种写法
	//newError.details = append(newError.details, details...)
	return &newError
}

func (e *ErrorCode) Error() string {
	return fmt.Sprintf("%#v", e)
}

var (
	Success       = NewError(0, "")
	ServerError   = NewError(10000000, "服务内部错误")
	InvalidParams = NewError(10000001, "入参错误")
	NotFound      = NewError(10000002, "找不到")
	Unauthorized  = NewError(10000003, "鉴权失败")
)

func NewServerError(err error) *ErrorCode {
	return ServerError.WithDetails(err.Error())
}

func NewInvalidError(err error) *ErrorCode {
	return InvalidParams.WithDetails(err.Error())
}

func NewNotFoundError(err error) *ErrorCode {
	return NotFound.WithDetails(err.Error())
}

func NewUnauthorized(err error) *ErrorCode {
	return Unauthorized.WithDetails(err.Error())
}
