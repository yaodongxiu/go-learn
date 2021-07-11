package my_error

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type ErrorCode struct {
	code int
	msg  string
}

func (code *ErrorCode) Code() int {
	return code.code
}

func (code *ErrorCode) Msg() string {
	return code.msg
}

func NewErrorCode(code int, msg string) *ErrorCode {
	return &ErrorCode{
		code: code,
		msg:  msg,
	}
}

func (code *ErrorCode) New(err error) *MyError {
	return &MyError{
		code: code.code,
		msg:  code.msg,
		err:  err,
	}
}

type MyError struct {
	code int
	msg  string
	err  error
}

func (e *MyError) Code() int {
	return e.code
}

func (e *MyError) Msg() string {
	return e.msg
}

func (e *MyError) Cause() string {
	return errors.Cause(e.err).Error()
}

func (e *MyError) Error() string {
	return e.err.Error()
}

func (e *MyError) Detail() string {
	return fmt.Sprintf("%+v", e.err)
}

func (e *MyError) Wrap(message string) *MyError {
	e.err = errors.Wrap(e.err, message)
	return e
}

func (e *MyError) Wrapf(message string, args ...interface{}) *MyError {
	e.err = errors.Wrapf(e.err, message, args...)
	return e
}

func (e *MyError) WithMessage(message string) *MyError {
	e.msg = message
	return e
}

func (e *MyError) WithMessagef(format string, args ...interface{}) *MyError {
	e.msg = fmt.Sprintf(format, args...)
	return e
}

func (e *MyError) StatusCode() int {
	switch e.Code() {
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
	case UnauthorizedApi.Code():
		return http.StatusForbidden
	}

	return http.StatusInternalServerError
}
