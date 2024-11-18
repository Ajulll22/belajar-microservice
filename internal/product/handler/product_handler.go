package handler

import (
	"github.com/Ajulll22/belajar-microservice/config"
	"github.com/Ajulll22/belajar-microservice/internal/product/service"
	"github.com/gin-gonic/gin"
)

func (h *productHandler) GetProducts(c *gin.Context) {

}

func (h *productHandler) GetProduct(c *gin.Context) {

}

func (h *productHandler) CreateProduct(c *gin.Context) {

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
