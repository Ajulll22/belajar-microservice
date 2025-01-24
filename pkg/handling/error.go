package handling

import (
	"errors"
	"net/http"
	"runtime"

	"github.com/Ajulll22/belajar-microservice/pkg/validator"
	"github.com/gin-gonic/gin"
)

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func (tcf Block) Do() {
	if tcf.Finally != nil {
		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}

var (
	CodeClientError         = 1001
	CodeNotFoundError       = 1002
	CodeConflictError       = 1003
	CodeServerError         = 1004
	CodeClientUnauthorized  = 1005
	CodeClientForbidden     = 1006
	CodeUnprocessableEntity = 1007

	CodeCacheMiss = 1008
)

var (
	MsgServerError               = "server error"
	MsgClientBadFormattedRequest = "bad format request"
	MsgClientNotFoundRequest     = "resource not found"
	MsgClientUnauthorized        = "unauthorized"
	MsgClientForbidden           = "forbidden"
	MsgValidationError           = "validation error"
	MsgConflictRequest           = "resource conflict"
)

type ErrorWrapper struct {
	Message    string                     `json:"message"` // human readable error
	Validation []validator.ErrorValidator `json:"-"`       //
	Code       int                        `json:"-"`       // code
	Err        error                      `json:"-"`       // original error
	Filename   string                     `json:"-"`
	LineNumber int                        `json:"-"`
}

func (w *ErrorWrapper) Error() string {
	// guard against panics
	if w.Err != nil {
		return w.Err.Error()
	}
	return w.Message
}

func NewErrorWrapper(code int, msg string, validation []validator.ErrorValidator, err error) *ErrorWrapper {
	// getting previous call stack file and line info
	_, filename, line, _ := runtime.Caller(1)
	return &ErrorWrapper{
		Code:       code,
		Message:    msg,
		Err:        err,
		Validation: validation,
		Filename:   filename,
		LineNumber: line,
	}
}

func RetryError(err error) bool {
	var ew *ErrorWrapper
	if errors.As(err, &ew) {
		if ew.Code != CodeServerError || ew.Code != CodeConflictError {
			return false
		}
	} else {
		return false
	}

	return true
}

func ResponseError(c *gin.Context, err error) BaseResponse[any] {
	resCode := http.StatusInternalServerError
	msg := MsgServerError
	var ew *ErrorWrapper
	if errors.As(err, &ew) {
		switch ew.Code {
		case CodeClientError:
			resCode = http.StatusBadRequest
			msg = MsgClientBadFormattedRequest
		case CodeNotFoundError:
			resCode = http.StatusNotFound
			msg = MsgClientNotFoundRequest
		case CodeConflictError:
			resCode = http.StatusConflict
			msg = MsgConflictRequest
		case CodeClientUnauthorized:
			resCode = http.StatusUnauthorized
			msg = MsgClientUnauthorized
		case CodeClientForbidden:
			resCode = http.StatusForbidden
			msg = MsgClientForbidden
		case CodeUnprocessableEntity:
			resCode = http.StatusUnprocessableEntity
			msg = MsgValidationError
		}

		return BaseResponse[any]{
			Message: msg,
			Code:    resCode,
			Error:   ew.Validation,
		}
	}

	return BaseResponse[any]{
		Message: msg,
		Code:    resCode,
	}
}
