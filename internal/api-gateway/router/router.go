package router

import (
	"github.com/Ajulll22/belajar-microservice/internal/api-gateway/config"
	"github.com/Ajulll22/belajar-microservice/internal/api-gateway/handler"
	"github.com/Ajulll22/belajar-microservice/internal/api-gateway/middleware"
	"github.com/Ajulll22/belajar-microservice/internal/api-gateway/service"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine, cfg config.Config) {
	api := router.Group("/api")
	asset := router.Group("/asset")

	gatewayService := service.NewGatewayService()

	gatewayHandler := handler.NewGatewayHandler(cfg, gatewayService)

	asset.GET("/*any", gatewayHandler.MediaProxy)

	api.Any("/auth/*any", gatewayHandler.UserProxy)
	api.Any("/user/*any", middleware.ValidateAccessToken(cfg.ACCESS_SECRET), gatewayHandler.UserProxy)

	api.GET("/product/*any", gatewayHandler.ProductProxy)
	api.POST("/product/*any", middleware.ValidateAccessToken(cfg.ACCESS_SECRET), gatewayHandler.ProductProxy)
	api.PUT("/product/*any", middleware.ValidateAccessToken(cfg.ACCESS_SECRET), gatewayHandler.ProductProxy)
	api.PATCH("/product/*any", middleware.ValidateAccessToken(cfg.ACCESS_SECRET), gatewayHandler.ProductProxy)
	api.DELETE("/product/*any", middleware.ValidateAccessToken(cfg.ACCESS_SECRET), gatewayHandler.ProductProxy)
}
