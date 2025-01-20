package repository

import (
	"context"

	"github.com/olivere/elastic/v7"
)

type LoggerRepository interface {
	Index(index string, m interface{}) error
}

func NewLoggerRepository(client *elastic.Client) LoggerRepository {
	return &loggerRepository{client}
}

type loggerRepository struct {
	client *elastic.Client
}

func (r *loggerRepository) Index(index string, m interface{}) error {
	_, err := r.client.Index().Index(index).BodyJson(m).Do(context.Background())

	return err
}
