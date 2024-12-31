package handler

import (
	"encoding/json"
	"errors"

	"github.com/Ajulll22/belajar-microservice/internal/user/config"
	"github.com/Ajulll22/belajar-microservice/internal/user/dto/request"
	"github.com/Ajulll22/belajar-microservice/internal/user/dto/response"
	"github.com/Ajulll22/belajar-microservice/internal/user/service"
	"github.com/Ajulll22/belajar-microservice/pkg/handling"
	v "github.com/Ajulll22/belajar-microservice/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (h *authHandler) Login(c *gin.Context) {
	ctx := c.Request.Context()
	validate := true

	authResponse := response.AuthResponse{}

	res := handling.ResponseSuccess(c, &authResponse, "Login success", 200)

	bodyRequest := request.Login{}
	err := c.ShouldBindJSON(&bodyRequest)
	if err != nil {
		validate = false

		var jsErr *json.UnmarshalTypeError
		var ve validator.ValidationErrors
		if errors.As(err, &jsErr) {
			res = handling.ResponseError(c, handling.NewErrorWrapper(handling.CodeClientError, "parse failed", nil, err))
		} else if errors.As(err, &ve) {
			errList := v.FormatValidation(ve)
			res = handling.ResponseError(c, handling.NewErrorWrapper(handling.CodeUnprocessableEntity, "invalid parameter", errList, err))
		} else {
			res = handling.ResponseError(c, err)
		}
	}

	if validate {

		err := h.authService.Login(ctx, &authResponse, bodyRequest.Username, bodyRequest.Password)
		if err != nil {
			validate = false
			res = handling.ResponseError(c, err)
		}

	}

	c.JSON(res.Code, res)
}
func (h *authHandler) RefreshToken(c *gin.Context) {
	ctx := c.Request.Context()
	validate := true

	authResponse := response.AuthResponse{}

	res := handling.ResponseSuccess(c, &authResponse, "Refresh token success", 200)

	bodyRequest := request.RefreshToken{}
	err := c.ShouldBindJSON(&bodyRequest)
	if err != nil {
		validate = false

		var jsErr *json.UnmarshalTypeError
		var ve validator.ValidationErrors
		if errors.As(err, &jsErr) {
			res = handling.ResponseError(c, handling.NewErrorWrapper(handling.CodeClientError, "parse failed", nil, err))
		} else if errors.As(err, &ve) {
			errList := v.FormatValidation(ve)
			res = handling.ResponseError(c, handling.NewErrorWrapper(handling.CodeUnprocessableEntity, "invalid parameter", errList, err))
		} else {
			res = handling.ResponseError(c, err)
		}
	}

	if validate {

		err := h.authService.RefreshToken(ctx, &authResponse, bodyRequest.RefreshToken)
		if err != nil {
			validate = false
			res = handling.ResponseError(c, err)
		}

	}

	c.JSON(res.Code, res)
}
func (h *authHandler) Logout(c *gin.Context) {
	ctx := c.Request.Context()
	validate := true

	res := handling.ResponseSuccess(c, nil, "Logout success", 200)

	bodyRequest := request.Logout{}
	err := c.ShouldBindJSON(&bodyRequest)
	if err != nil {
		validate = false

		var jsErr *json.UnmarshalTypeError
		var ve validator.ValidationErrors
		if errors.As(err, &jsErr) {
			res = handling.ResponseError(c, handling.NewErrorWrapper(handling.CodeClientError, "parse failed", nil, err))
		} else if errors.As(err, &ve) {
			errList := v.FormatValidation(ve)
			res = handling.ResponseError(c, handling.NewErrorWrapper(handling.CodeUnprocessableEntity, "invalid parameter", errList, err))
		} else {
			res = handling.ResponseError(c, err)
		}
	}

	if validate {

		err := h.authService.Logout(ctx, bodyRequest.RefreshToken)
		if err != nil {
			validate = false
			res = handling.ResponseError(c, err)
		}

	}

	c.JSON(res.Code, res)
}

type authHandler struct {
	cfg         config.Config
	authService service.AuthService
}

func NewAuthHandler(cfg config.Config, authService service.AuthService) authHandler {
	return authHandler{cfg, authService}
}
