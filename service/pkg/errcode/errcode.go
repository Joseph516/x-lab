package errcode

import (
	"fmt"
	"log"
	"net/http"
)

// 错误码操作程序包

// 已经定义错误码
var codes = map[int]string{}

type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		log.Panicf("%v has already be defined", code)
	}

	codes[code] = msg
	return &Error{
		code: code,
		msg:  msg,
	}
}

// Error 字符串格式化输出错误码
func (e *Error) Error() string {
	return fmt.Sprintf("Error code: %v, msg: %v", e.code, e.msg)
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Details() []string {
	return e.details
}

// AddDetails 将错误详情details写入Error
func (e *Error) AddDetails(details ...string) {
	for _, detail := range details {
		e.details = append(e.details, detail)
	}
}

// WithDetails 将原Error新增details，并返回新的Error
func (e *Error) WithDetails(details ...string) *Error {
	newErr := *e
	newErr.AddDetails(details...)
	return &newErr
}

// TODO: 没啥必要，后续去除
func (e *Error) ToHttpStatusCode() int {
	code := e.Code()
	switch code {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}
