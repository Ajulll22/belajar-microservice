package consumer

import (
	"encoding/json"
	"log"

	"github.com/Ajulll22/belajar-microservice/internal/logger/config"
	"github.com/Ajulll22/belajar-microservice/internal/logger/model"
	"github.com/Ajulll22/belajar-microservice/internal/logger/service"
	"github.com/Ajulll22/belajar-microservice/pkg/handling"
	"github.com/go-playground/validator/v10"
	"github.com/streadway/amqp"
)

type LoggerConsumer interface {
	IndexApplicationLog(msg amqp.Delivery) error
	IndexAuditLog(msg amqp.Delivery) error
	IndexPerformLog(msg amqp.Delivery) error
	IndexErrorLog(msg amqp.Delivery) error
}

func NewLoggerService(cfg config.Config, loggerService service.LoggerService) LoggerConsumer {
	return &loggerConsumer{cfg, loggerService}
}

type loggerConsumer struct {
	cfg           config.Config
	loggerService service.LoggerService
}

func (c *loggerConsumer) IndexApplicationLog(msg amqp.Delivery) error {
	messageByte := msg.Body
	log.Println("Receive message :", string(messageByte))

	messageBody := model.ApplicationLog{}
	validate := validator.New()

	if err := json.Unmarshal(messageByte, &messageBody); err != nil {
		return handling.NewErrorWrapper(handling.CodeClientError, "parse failed", nil, nil)
	}
	if err := validate.Struct(messageBody); err != nil {
		return handling.NewErrorWrapper(handling.CodeUnprocessableEntity, "invalid parameter", nil, nil)
	}

	if err := c.loggerService.IndexApplicationLog(&messageBody); err != nil {
		return err
	}

	return nil
}

func (c *loggerConsumer) IndexAuditLog(msg amqp.Delivery) error {
	messageByte := msg.Body
	log.Println("Receive message :", string(messageByte))

	messageBody := model.AuditLog{}
	validate := validator.New()

	if err := json.Unmarshal(messageByte, &messageBody); err != nil {
		return handling.NewErrorWrapper(handling.CodeClientError, "parse failed", nil, nil)
	}
	if err := validate.Struct(messageBody); err != nil {
		return handling.NewErrorWrapper(handling.CodeUnprocessableEntity, "invalid parameter", nil, nil)
	}

	if err := c.loggerService.IndexAuditLog(&messageBody); err != nil {
		return err
	}

	return nil
}

func (c *loggerConsumer) IndexPerformLog(msg amqp.Delivery) error {
	messageByte := msg.Body
	log.Println("Receive message :", string(messageByte))

	messageBody := model.PerformLog{}
	validate := validator.New()

	if err := json.Unmarshal(messageByte, &messageBody); err != nil {
		return handling.NewErrorWrapper(handling.CodeClientError, "parse failed", nil, nil)
	}
	if err := validate.Struct(messageBody); err != nil {
		return handling.NewErrorWrapper(handling.CodeUnprocessableEntity, "invalid parameter", nil, nil)
	}

	if err := c.loggerService.IndexPerformLog(&messageBody); err != nil {
		return err
	}

	return nil
}

func (c *loggerConsumer) IndexErrorLog(msg amqp.Delivery) error {
	messageByte := msg.Body
	log.Println("Receive message :", string(messageByte))

	messageBody := model.ErrorLog{}
	validate := validator.New()

	if err := json.Unmarshal(messageByte, &messageBody); err != nil {
		return handling.NewErrorWrapper(handling.CodeClientError, "parse failed", nil, nil)
	}
	if err := validate.Struct(messageBody); err != nil {
		return handling.NewErrorWrapper(handling.CodeUnprocessableEntity, "invalid parameter", nil, nil)
	}

	if err := c.loggerService.IndexErrorLog(&messageBody); err != nil {
		return err
	}

	return nil
}
