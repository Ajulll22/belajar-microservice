package consumer

import (
	"encoding/json"
	"log"

	"github.com/Ajulll22/belajar-microservice/internal/media/config"
	"github.com/Ajulll22/belajar-microservice/internal/media/dto/request"
	"github.com/Ajulll22/belajar-microservice/internal/media/repository"
	"github.com/Ajulll22/belajar-microservice/pkg/handling"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

func NewMediaConsumer(db *mongo.Database, cfg config.Config, mediaRepository repository.MediaRepository) MediaConsumer {
	return &mediaConsumer{db, cfg, mediaRepository}
}

type MediaConsumer interface {
	Run([]byte) error
}

type mediaConsumer struct {
	db              *mongo.Database
	cfg             config.Config
	mediaRepository repository.MediaRepository
}

func (c *mediaConsumer) Run(messageByte []byte) error {
	log.Println("Receive message :", string(messageByte))
	messageBody := request.DeleteMedia{}

	err := json.Unmarshal(messageByte, &messageBody)
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeUnprocessableEntity, "failed to unmarshal message", nil, err)
	}

	bucket, err := gridfs.NewBucket(c.db)
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeServerError, "failed to initialize gridfs", nil, err)
	}

	err = c.mediaRepository.DeleteByID(bucket, messageBody.ID)
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeNotFoundError, "file not found", nil, err)
	}

	log.Println("Success delete file", messageBody.ID)
	return nil
}
