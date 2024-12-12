package router

import (
	"github.com/Ajulll22/belajar-microservice/internal/media/config"
	"github.com/Ajulll22/belajar-microservice/internal/media/handler"
	"github.com/Ajulll22/belajar-microservice/internal/media/repository"
	"github.com/Ajulll22/belajar-microservice/internal/media/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Register(router *gin.Engine, db *mongo.Database, cfg config.Config) {
	mediaRepository := repository.NewMediaRepository()

	mediaService := service.NewMediaService(db, cfg, mediaRepository)

	mediaHandler := handler.NewMediaHandler(cfg, mediaService)

	asset := router.Group("/asset")
	{
		asset.GET("/:fileID", mediaHandler.GetMedia)
	}

	api := router.Group("/api")
	mediaRouter := api.Group("/media")
	{
		mediaRouter.POST("/", mediaHandler.UploadMedia)
		mediaRouter.DELETE("/:fileID", mediaHandler.DeleteMedia)
	}
}
