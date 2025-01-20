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
	}
}

func RegisterConsumer(db *mongo.Database, cfg config.Config, rmq broker.RabbitMQ) {
	mediaRepository := repository.NewMediaRepository()

	mediaService := service.NewMediaService(db, cfg, mediaRepository)

	mediaConsumer := consumer.NewMediaConsumer(cfg, mediaService)

	routes := []broker.ConsumerRoute{
		{
			Key:     "delete_media",
			Handler: mediaConsumer.DeleteMedia,
			Async:   false,
			AutoAck: false,
			Queue:   "delete_media_" + cfg.MEDIA_QUEUE,
		},
	}

	err := rmq.DeclareExchange(cfg.MEDIA_EXCHANGE, "direct")
	if err != nil {
		log.Println(err)
	}

	for _, route := range routes {

		queue, err := rmq.DeclareQueue(route.Queue) //
		if err != nil {
			log.Println(err)
		}

		err = rmq.BindQueue(queue.Name, cfg.MEDIA_EXCHANGE, route.Key)
		if err != nil {
			log.Println(err)
		}

	}
	err = rmq.Consume(routes)
	if err != nil {
		log.Println(err)
	}
}
