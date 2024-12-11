package handling

import "github.com/Ajulll22/belajar-microservice/pkg/validator"

type BaseResponse struct {
	Message string                     `json:"message"`
	Error   []validator.ErrorValidator `json:"error,omitempty"`
	Data    interface{}                `json:"data,omitempty"`
	Code    int                        `json:"code"`
}
