package service

import (
	"bytes"

	"github.com/Ajulll22/belajar-microservice/internal/media/config"
	"github.com/Ajulll22/belajar-microservice/internal/media/model"
	"github.com/Ajulll22/belajar-microservice/internal/media/repository"
	"github.com/Ajulll22/belajar-microservice/pkg/handling"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

type MediaService interface {
	UploadMedia(*[]model.Media) error
	DeleteMedia(string) error
	GetMedia(string) (bytes.Buffer, error)
}

type mediaService struct {
	db              *mongo.Database
	cfg             config.Config
	mediaRepository repository.MediaRepository
}

func NewMediaService(db *mongo.Database, cfg config.Config, mediaRepository repository.MediaRepository) MediaService {
	return &mediaService{db, cfg, mediaRepository}
}

func (s *mediaService) UploadMedia(mediaList *[]model.Media) error {
	bucket, err := gridfs.NewBucket(s.db)
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeServerError, "failed to initialize gridfs", nil, err)
	}

	data := *mediaList

	for i := 0; i < len(data); i++ {
		err := s.mediaRepository.Upload(bucket, &data[i])
		if err != nil {
			return handling.NewErrorWrapper(handling.CodeServerError, "failed to upload file to gridfs", nil, err)
		}
	}

	*mediaList = data

	return nil
}

func (s *mediaService) DeleteMedia(fileID string) error {
	bucket, err := gridfs.NewBucket(s.db)
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeServerError, "failed to initialize gridfs", nil, err)
	}

	err = s.mediaRepository.DeleteByID(bucket, fileID)
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeNotFoundError, "file not found", nil, err)
	}

	return nil
}

func (s *mediaService) GetMedia(fileID string) (buffer bytes.Buffer, err error) {
	bucket, err := gridfs.NewBucket(s.db)
	if err != nil {
		return buffer, handling.NewErrorWrapper(handling.CodeServerError, "failed to initialize gridfs", nil, err)
	}

	buffer, err = s.mediaRepository.FindByID(bucket, fileID)
	if err != nil {
		return buffer, handling.NewErrorWrapper(handling.CodeNotFoundError, "file not found", nil, err)
	}

	return buffer, nil
}
