package consumer

import (
	"log"

	"github.com/Ajulll22/belajar-microservice/internal/media/config"
	"github.com/Ajulll22/belajar-microservice/internal/media/repository"
	"github.com/Ajulll22/belajar-microservice/pkg/broker"
	"go.mongodb.org/mongo-driver/mongo"
)

func Register(db *mongo.Database, cfg config.Config, rmq broker.RabbitMQ) {
	mediaRepository := repository.NewMediaRepository()

	mediaConsumer := NewMediaConsumer(db, cfg, mediaRepository)

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
