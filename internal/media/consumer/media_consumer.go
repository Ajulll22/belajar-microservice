package consumer

import (
	"encoding/json"
	"log"

	"github.com/Ajulll22/belajar-microservice/internal/media/config"
	"github.com/Ajulll22/belajar-microservice/internal/media/dto/request"
	"github.com/Ajulll22/belajar-microservice/internal/media/service"
	"github.com/Ajulll22/belajar-microservice/pkg/handling"
	"github.com/streadway/amqp"
)

func NewMediaConsumer(cfg config.Config, mediaService service.MediaService) MediaConsumer {
	return &mediaConsumer{cfg, mediaService}
}

type MediaConsumer interface {
	DeleteMedia(msg amqp.Delivery) error
}

type mediaConsumer struct {
	cfg          config.Config
	mediaService service.MediaService
}

func (c *mediaConsumer) DeleteMedia(msg amqp.Delivery) error {
	messageByte := msg.Body
	log.Println("Receive message :", string(messageByte))
	messageBody := request.DeleteMedia{}

	err := json.Unmarshal(messageByte, &messageBody)
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeUnprocessableEntity, "failed to unmarshal message", nil, nil)
	}

	err = c.mediaService.DeleteMedia(messageBody.ID)
	if err != nil {
		return err
	}

	log.Println("Success delete file", messageBody.ID)
	return nil
}
