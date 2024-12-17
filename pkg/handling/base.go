package handling

import "github.com/Ajulll22/belajar-microservice/pkg/validator"

type BaseResponse[T any] struct {
	Message string                     `json:"message"`
	Error   []validator.ErrorValidator `json:"error,omitempty"`
	Data    T                          `json:"data,omitempty"`
	Code    int                        `json:"code"`
}
