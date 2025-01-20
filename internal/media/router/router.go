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
	"github.com/streadway/amqp"
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

	mainExchange := cfg.MEDIA_EXCHANGE
	retryExchange := "retry_" + cfg.MEDIA_EXCHANGE
	dlxExchange := "dlx_" + cfg.MEDIA_EXCHANGE

	retryQueue := "retry_" + cfg.MEDIA_QUEUE
	dlxQueue := "dlx" + cfg.MEDIA_QUEUE

	err := rmq.DeclareExchange(mainExchange, "direct")
	if err != nil {
		log.Println(err)
		return
	}
	err = rmq.DeclareExchange(retryExchange, "direct")
	if err != nil {
		log.Println(err)
		return
	}
	err = rmq.DeclareExchange(dlxExchange, "direct")
	if err != nil {
		log.Println(err)
		return
	}

	for _, route := range routes {

		queue, err := rmq.DeclareQueue(route.Queue, amqp.Table{
			"x-dead-letter-exchange": retryExchange,
		})
		if err != nil {
			log.Println(err)
			return
		}

		err = rmq.BindQueue(queue.Name, mainExchange, route.Key)
		if err != nil {
			log.Println(err)
			return
		}

	}

	_, err = rmq.DeclareQueue(retryQueue, amqp.Table{
		"x-dead-letter-exchange": mainExchange,
		"x-message-ttl":          int32(5000), // 5 seconds TTL
	})
	if err != nil {
		log.Println(err)
		return
	}
	_, err = rmq.DeclareQueue(dlxQueue, nil)
	if err != nil {
		log.Println(err)
		return
	}

	err = rmq.BindQueue(retryQueue, retryExchange, "retry_routing_key")
	if err != nil {
		log.Println(err)
		return
	}
	err = rmq.BindQueue(dlxQueue, dlxExchange, "dlx_routing_key")
	if err != nil {
		log.Println(err)
		return
	}

	err = rmq.Consume(routes)
	if err != nil {
		log.Println(err)
		return
	}
}
