package repository

import (
	"bytes"
	"fmt"
	"io"

	"github.com/Ajulll22/belajar-microservice/internal/media/model"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

type MediaRepository interface {
	Upload(*gridfs.Bucket, *model.Media) error
	FindByID(*gridfs.Bucket, string) (bytes.Buffer, error)
	DeleteByID(*gridfs.Bucket, string) error
}

type mediaRepository struct {
}

func NewMediaRepository() MediaRepository {
	return &mediaRepository{}
}

func (r *mediaRepository) Upload(bucket *gridfs.Bucket, media *model.Media) error {
	src, err := media.File.Open()
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer src.Close()

	media.ID = uuid.New().String()

	uploadStream, err := bucket.OpenUploadStreamWithID(media.ID, media.File.Filename)
	if err != nil {
		return fmt.Errorf("failed to open upload stream: %v", err)
	}
	defer uploadStream.Close()

	// Salin data dari file ke stream menggunakan io.Copy
	if _, err := io.Copy(uploadStream, src); err != nil {
		return fmt.Errorf("failed to write file to GridFS: %v", err)
	}

	return nil
}

func (r *mediaRepository) FindByID(bucket *gridfs.Bucket, id string) (buffer bytes.Buffer, err error) {
	_, err = bucket.DownloadToStream(id, &buffer)

	return buffer, err
}

func (r *mediaRepository) DeleteByID(bucket *gridfs.Bucket, id string) (err error) {
	err = bucket.Delete(id)

	return err
}
