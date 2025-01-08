package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Ajulll22/belajar-microservice/internal/media/config"
	"github.com/Ajulll22/belajar-microservice/internal/media/dto/request"
	"github.com/Ajulll22/belajar-microservice/internal/media/model"
	"github.com/Ajulll22/belajar-microservice/internal/media/service"
	"github.com/Ajulll22/belajar-microservice/pkg/handling"
	v "github.com/Ajulll22/belajar-microservice/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (h *mediaHandler) UploadMedia(c *gin.Context) {
	validate := true

	mediaDataList := []model.Media{}

	res := handling.ResponseSuccess(c, &mediaDataList, "Upload Media Success", 200)

	bodyRequest := request.UploadMedia{}
	if validate {

		err := c.ShouldBind(&bodyRequest)
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

	}

	if validate {

		for _, file := range bodyRequest.Files {
			mediaDataList = append(mediaDataList, model.Media{
				File: file,
			})
		}

		err := h.mediaService.UploadMedia(&mediaDataList)
		if err != nil {
			validate = false
			res = handling.ResponseError(c, err)
		}

	}

	c.JSON(res.Code, res)
}

func (h *mediaHandler) DeleteMedia(c *gin.Context) {
	fileID := c.Param("fileID")

	res := handling.ResponseSuccess(c, nil, "Delete Media Success", 200)

	err := h.mediaService.DeleteMedia(fileID)
	if err != nil {
		res = handling.ResponseError(c, err)
	}

	c.JSON(res.Code, res)
}

func (h *mediaHandler) GetMedia(c *gin.Context) {
	fileID := c.Param("fileID")

	buffer, err := h.mediaService.GetMedia(fileID)
	if err != nil {
		res := handling.ResponseError(c, err)
		c.JSON(res.Code, res)
		return
	}

	contentType := http.DetectContentType(buffer.Bytes())
	c.Writer.Header().Set("Content-Type", contentType)
	c.Writer.Header().Set("Content-Disposition", "inline")

	c.Data(http.StatusOK, contentType, buffer.Bytes())
}

type mediaHandler struct {
	cfg          config.Config
	mediaService service.MediaService
}

func NewMediaHandler(cfg config.Config, mediaService service.MediaService) mediaHandler {
	return mediaHandler{cfg, mediaService}
}
