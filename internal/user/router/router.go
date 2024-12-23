package router

import (
	"github.com/Ajulll22/belajar-microservice/internal/user/config"
	"github.com/Ajulll22/belajar-microservice/internal/user/handler"
	"github.com/Ajulll22/belajar-microservice/internal/user/repository"
	"github.com/Ajulll22/belajar-microservice/internal/user/service"
	"github.com/Ajulll22/belajar-microservice/pkg/cache"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(router *gin.Engine, db *gorm.DB, redis cache.Cache, cfg config.Config) {

	api := router.Group("/api")

	userRepository := repository.NewUserRepository()

	authService := service.NewAuthService(cfg, db, redis, userRepository)

	authHandler := handler.NewAuthHandler(cfg, authService)

	authRouter := api.Group("/auth")
	{
		authRouter.POST("/login", authHandler.Login)
		authRouter.POST("/refresh-token", authHandler.RefreshToken)
		authRouter.POST("/logout", authHandler.Logout)
	}

}
