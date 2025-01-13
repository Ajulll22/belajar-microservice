package router

import (
	"log"

	"github.com/Ajulll22/belajar-microservice/internal/product/config"
	"github.com/Ajulll22/belajar-microservice/internal/product/handler"
	"github.com/Ajulll22/belajar-microservice/internal/product/repository"
	"github.com/Ajulll22/belajar-microservice/internal/product/service"
	"github.com/Ajulll22/belajar-microservice/pkg/broker"
	"github.com/Ajulll22/belajar-microservice/pkg/cache"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(router *gin.Engine, db *gorm.DB, redis cache.Cache, cfg config.Config, rmq broker.RabbitMQ) {

	api := router.Group("/api")

	productRepository := repository.NewProductRepository()
	categoryRepository := repository.NewCategoryRepository()

	err := rmq.DeclareExchange(cfg.MEDIA_EXCHANGE, "direct")
	if err != nil {
		log.Println(err)
	}

	productService := service.NewProductService(cfg, db, redis, rmq, productRepository, categoryRepository)

	productHandler := handler.NewProductHandler(cfg, productService)

	productRouter := api.Group("/product")
	{
		productRouter.GET("/", productHandler.GetProducts)
		productRouter.GET("/:id", productHandler.GetProduct)
		productRouter.POST("/", productHandler.CreateProduct)
		productRouter.PUT("/:id", productHandler.UpdateProduct)
	}

}
