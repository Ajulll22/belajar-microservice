package router

import (
	"github.com/Ajulll22/belajar-microservice/internal/product/config"
	"github.com/Ajulll22/belajar-microservice/internal/product/handler"
	"github.com/Ajulll22/belajar-microservice/internal/product/repository"
	"github.com/Ajulll22/belajar-microservice/internal/product/service"
	"github.com/Ajulll22/belajar-microservice/pkg/cache"
	redis "github.com/go-redis/cache/v9"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(router *gin.Engine, db *gorm.DB, redis *redis.Cache, cfg config.Config) {

	redisCache := cache.NewRedisCache(redis)
	api := router.Group("/api")

	productRepository := repository.NewProductRepository()
	categoryRepository := repository.NewCategoryRepository()

	productService := service.NewProductService(cfg, db, redisCache, productRepository, categoryRepository)

	productHandler := handler.NewProductHandler(cfg, productService)

	productRouter := api.Group("/product")
	{
		productRouter.GET("/", productHandler.GetProducts)
		productRouter.GET("/:id", productHandler.GetProduct)
		productRouter.POST("/", productHandler.CreateProduct)
	}

}
