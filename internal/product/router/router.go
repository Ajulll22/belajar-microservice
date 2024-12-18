package router

import (
	"github.com/Ajulll22/belajar-microservice/internal/product/config"
	"github.com/Ajulll22/belajar-microservice/internal/product/handler"
	"github.com/Ajulll22/belajar-microservice/internal/product/repository"
	"github.com/Ajulll22/belajar-microservice/internal/product/service"
	"github.com/Ajulll22/belajar-microservice/pkg/cache"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(router *gin.Engine, db *gorm.DB, redis cache.Cache, cfg config.Config) {

	api := router.Group("/api")

	productRepository := repository.NewProductRepository()
	categoryRepository := repository.NewCategoryRepository()

	productService := service.NewProductService(cfg, db, redis, productRepository, categoryRepository)

	productHandler := handler.NewProductHandler(cfg, productService)

	productRouter := api.Group("/product")
	{
		productRouter.GET("/", productHandler.GetProducts)
		productRouter.GET("/:id", productHandler.GetProduct)
		productRouter.POST("/", productHandler.CreateProduct)
	}

}
