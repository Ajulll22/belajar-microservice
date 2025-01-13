package router

import (
	"log"

	"github.com/Ajulll22/belajar-microservice/internal/media/config"
	"github.com/Ajulll22/belajar-microservice/internal/media/consumer"
	"github.com/Ajulll22/belajar-microservice/internal/media/handler"
	"github.com/Ajulll22/belajar-microservice/internal/media/repository"
	"github.com/Ajulll22/belajar-microservice/internal/media/service"
	"github.com/Ajulll22/belajar-microservice/pkg/broker"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Register(router *gin.Engine, db *mongo.Database, cfg config.Config, rmq broker.RabbitMQ) {
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

func RegisterConsumer(db *mongo.Database, cfg config.Config, rmq broker.RabbitMQ) {
	mediaRepository := repository.NewMediaRepository()

	mediaConsumer := consumer.NewMediaConsumer(db, cfg, mediaRepository)

	err := rmq.DeclareExchange(cfg.MEDIA_EXCHANGE, "direct")
	if err != nil {
		log.Println(err)
	}
	queue, err := rmq.DeclareQueue(cfg.MEDIA_QUEUE)
	if err != nil {
		log.Println(err)
	}
	err = rmq.BindQueue(queue.Name, cfg.MEDIA_EXCHANGE, "delete_media")
	if err != nil {
		log.Println(err)
	}
	err = rmq.Consume(queue.Name, mediaConsumer.Run)
	if err != nil {
		log.Println(err)
	}
}
