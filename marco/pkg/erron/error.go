package erron

import (
	"fmt"
	"strings"
)

//MarcoError error for marco
type MarcoError struct {
	Code    int
	Message string //error message
	ERR     error  // wrapped error
}

//Error implement error interface, 将内部error和Message整合在一起输出
func (r *MarcoError) Error() string {
	var sb strings.Builder
	if len(r.Message) > 0 {
		sb.WriteString(r.Message)
	}
	if r.ERR != nil {
		if sb.Len() > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("with internal error: %s", r.ERR.Error()))
	}
	return sb.String()
}

// UnWrapError 解包响应错误
func UnWrapError(err error) *MarcoError {
	if v, ok := err.(*MarcoError); ok {
		return v
	}
	return nil
}

//WrapError wrap error with code and message
func WrapError(err error, code int, msg string) error {
	res := &MarcoError{
		Code:    code,
		Message: msg,
		ERR:     err,
	}
	return res
}

//New create new error of marco error
func New(err error, code int, msg string) error {
	return &MarcoError{
		Code:    code,
		Message: msg,
		ERR:     err,
	}
}
