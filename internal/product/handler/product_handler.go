package handler

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/Ajulll22/belajar-microservice/internal/product/config"
	"github.com/Ajulll22/belajar-microservice/internal/product/dto/request"
	"github.com/Ajulll22/belajar-microservice/internal/product/model"
	"github.com/Ajulll22/belajar-microservice/internal/product/service"
	"github.com/Ajulll22/belajar-microservice/pkg/handling"
	v "github.com/Ajulll22/belajar-microservice/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (h *productHandler) GetProducts(c *gin.Context) {
	ctx := c.Request.Context()
	validate := true
	products := []model.Product{}

	res := handling.ResponseSuccess(c, &products, "Get products success", 200)

	if validate {

		err := h.productService.GetProducts(ctx, &products)
		if err != nil {
			validate = false
			res = handling.ResponseError(c, err)
		}

	}

	c.JSON(res.Code, res)
}

func (h *productHandler) GetProduct(c *gin.Context) {
	paramID := c.Param("id")
	ctx := c.Request.Context()
	validate := true

	ID := 0
	product := model.Product{}

	res := handling.ResponseSuccess(c, &product, "Get products success", 200)

	if validate {

		ID := &ID
		val, err := strconv.Atoi(paramID)
		if err != nil {
			validate = false
			res = handling.ResponseError(c, handling.NewErrorWrapper(handling.CodeClientError, "failed to convert id to int", nil, err))
		} else {
			*ID = val
		}

	}

	if validate {

		err := h.productService.GetProduct(ctx, &product, ID)
		if err != nil {
			validate = false
			res = handling.ResponseError(c, err)
		}

	}

	c.JSON(res.Code, res)
}

func (h *productHandler) CreateProduct(c *gin.Context) {
	ctx := c.Request.Context()
	validate := true

	productData := model.Product{}

	res := handling.ResponseSuccess(c, &productData, "Get products success", 200)

	bodyRequest := request.ProductInsert{}
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

	if validate {

		price, _ := bodyRequest.Price.Int64()
		stock, _ := bodyRequest.Stock.Int64()

		productData := &productData
		productData.Name = bodyRequest.Name
		productData.Price = int(price)
		productData.Stock = int(stock)
		productData.Description = bodyRequest.Description

		for _, categoryID := range bodyRequest.Categories {
			productData.Categories = append(productData.Categories, model.ProductCategory{
				ID: categoryID,
			})
		}
		err := h.productService.CreateProduct(ctx, productData)
		if err != nil {
			validate = false
			res = handling.ResponseError(c, err)
		}

	}

	c.JSON(res.Code, res)
}

func (h *productHandler) UpdateProduct(c *gin.Context) {

}

func (h *productHandler) DeleteProduct(c *gin.Context) {

}

type productHandler struct {
	cfg            config.Config
	productService service.ProductService
}

func NewProductHandler(cfg config.Config, productService service.ProductService) productHandler {
	return productHandler{cfg, productService}
}
