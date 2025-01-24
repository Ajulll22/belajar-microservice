package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"

	"github.com/Ajulll22/belajar-microservice/internal/product/config"
	"github.com/Ajulll22/belajar-microservice/internal/product/dto/response"
	"github.com/Ajulll22/belajar-microservice/internal/product/model"
	"github.com/Ajulll22/belajar-microservice/internal/product/repository"
	"github.com/Ajulll22/belajar-microservice/pkg/broker"
	"github.com/Ajulll22/belajar-microservice/pkg/cache"
	"github.com/Ajulll22/belajar-microservice/pkg/constant"
	"github.com/Ajulll22/belajar-microservice/pkg/handling"
	"github.com/Ajulll22/belajar-microservice/pkg/service"
	"github.com/Ajulll22/belajar-microservice/pkg/validator"
	"gorm.io/gorm"
)

type ProductService interface {
	GetProducts(ctx context.Context, m *[]model.Product) error
	GetProduct(ctx context.Context, m *model.Product, id int) error
	CreateProduct(ctx context.Context, m *model.Product, pictures []*multipart.FileHeader) error
	UpdateProduct(ctx context.Context, m *model.Product, pictures []*multipart.FileHeader) error
	DeleteProduct(ctx context.Context, m *model.Product) error
}

func NewProductService(cfg config.Config, db *gorm.DB, cache cache.Cache, rmq broker.RabbitMQ, productRepository repository.ProductRepository, productPictureRepository repository.ProductPictureRepository, categoryRepository repository.CategoryRepository) ProductService {
	return &productService{cfg, db, cache, rmq, productRepository, productPictureRepository, categoryRepository}
}

type productService struct {
	cfg                      config.Config
	db                       *gorm.DB
	cache                    cache.Cache
	rmq                      broker.RabbitMQ
	productRepository        repository.ProductRepository
	productPictureRepository repository.ProductPictureRepository
	categoryRepository       repository.CategoryRepository
}

func (s *productService) GetProducts(ctx context.Context, m *[]model.Product) error {

	err := s.cache.Get(ctx, cache.GetCacheKey(s.cfg.CACHE_KEY_PRODUCT), m)
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeServerError, "failed to get products from cache", nil, nil)
	}

	if len(*m) == 0 {

		err := s.productRepository.FindAll(s.db, m)
		if err != nil {
			return handling.NewErrorWrapper(handling.CodeServerError, "failed to get products from db", nil, nil)
		}

		err = s.cache.Set(ctx, cache.GetCacheKey(s.cfg.CACHE_KEY_PRODUCT), *m, constant.CacheTTLOneDay)
		if err != nil {
			return handling.NewErrorWrapper(handling.CodeServerError, "failed to set products to cache", nil, nil)
		}

	}
	return nil
}

func (s *productService) GetProduct(ctx context.Context, m *model.Product, id int) error {

	products := []model.Product{}
	err := s.cache.Get(ctx, cache.GetCacheKey(s.cfg.CACHE_KEY_PRODUCT), &products)
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeServerError, "failed to get products from cache", nil, nil)
	}

	if len(products) != 0 {

		for _, val := range products {
			if val.ID == id {
				*m = val

				return nil
			}

		}

	}

	err = s.productRepository.FindByID(s.db, m, id)
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeServerError, "failed to get product from db", nil, nil)
	}

	if m.ID == 0 {
		return handling.NewErrorWrapper(handling.CodeNotFoundError, "data product not found", nil, nil)
	}

	return nil
}

func (s *productService) CreateProduct(ctx context.Context, m *model.Product, pictures []*multipart.FileHeader) error {
	dataTransaction := s.db.Begin()
	var err error
	defer func() {
		if err != nil {
			dataTransaction.Rollback()
		} else {
			dataTransaction.Commit()
		}
	}()

	categories := []model.Category{}
	err = s.cache.Get(ctx, cache.GetCacheKey(s.cfg.CACHE_KEY_CATEGORY), &categories)
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeServerError, "failed to get categories from cache", nil, nil)
	}
	if len(categories) == 0 {
		err = s.categoryRepository.FindAll(dataTransaction, &categories)
		if err != nil {
			return handling.NewErrorWrapper(handling.CodeServerError, "failed to get category from db", nil, nil)
		}
	}

	errList := []validator.ErrorValidator{}
	for index, Category := range m.Categories {
		found := false

		for _, category := range categories {
			if Category.ID == category.ID {
				found = true
				m.Categories[index].Name = category.Name
				break
			}
		}

		if !found {
			errList = append(errList, validator.ErrorValidator{
				Key:     fmt.Sprintf("categories[%d]", index),
				Message: "category id not found",
			})
		}
	}
	if len(errList) != 0 {
		return handling.NewErrorWrapper(handling.CodeUnprocessableEntity, "invalid parameter", errList, nil)
	}

	responseMedia := handling.BaseResponse[[]struct {
		ID string `json:"id"`
	}]{}
	url := fmt.Sprintf("http://%s:%s/api/media", s.cfg.MEDIA_SERVICE_NAME, s.cfg.MEDIA_SERVICE_PORT)
	res, err := service.ForwardFilesToService(url, pictures)
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeServerError, "failed to upload picture to media", nil, nil)
	} else if res.StatusCode < 300 {
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			return handling.NewErrorWrapper(handling.CodeServerError, "failed to read response http", nil, nil)
		}

		b := bytes.NewBuffer(resBody)
		d := json.NewDecoder(b)
		err = d.Decode(&responseMedia)
		if err != nil {
			return handling.NewErrorWrapper(handling.CodeServerError, "failed to parse response", nil, nil)
		}
	}
	defer res.Body.Close()

	for _, media := range responseMedia.Data {
		m.Pictures = append(m.Pictures, model.ProductPicture{
			Url: media.ID,
		})
	}

	err = s.productRepository.Insert(s.db, m)
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeServerError, "failed to insert product to db", nil, nil)
	}

	err = s.cache.Set(ctx, cache.GetCacheKey(s.cfg.CACHE_KEY_PRODUCT), nil, constant.CacheTTLInvalidate)
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeServerError, "failed to delete product from cache", nil, nil)
	}

	return nil
}

func (s *productService) UpdateProduct(ctx context.Context, m *model.Product, pictures []*multipart.FileHeader) error {
	dataTransaction := s.db.Begin()
	var err error
	defer func() {
		if err != nil {
			dataTransaction.Rollback()
		} else {
			dataTransaction.Commit()
		}
	}()

	oldProduct := model.Product{}
	err = s.GetProduct(ctx, &oldProduct, m.ID)
	if err != nil {
		return err
	}

	categories := []model.Category{}
	err = s.cache.Get(ctx, cache.GetCacheKey(s.cfg.CACHE_KEY_CATEGORY), &categories)
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeServerError, "failed to get categories from cache", nil, nil)
	}
	if len(categories) == 0 {
		err = s.categoryRepository.FindAll(dataTransaction, &categories)
		if err != nil {
			return handling.NewErrorWrapper(handling.CodeServerError, "failed to get category from db", nil, nil)
		}
	}

	errList := []validator.ErrorValidator{}
	for index, Category := range m.Categories {
		found := false

		for _, category := range categories {
			if Category.ID == category.ID {
				found = true
				m.Categories[index].Name = category.Name
				break
			}
		}

		if !found {
			errList = append(errList, validator.ErrorValidator{
				Key:     fmt.Sprintf("categories[%d]", index),
				Message: "category id not found",
			})
		}
	}
	for index, Picture := range m.Pictures {
		found := false

		for _, picture := range oldProduct.Pictures {
			if Picture.Url == picture.Url {
				found = true
				m.Pictures[index].ID = picture.ID
				break
			}
		}

		if !found {
			errList = append(errList, validator.ErrorValidator{
				Key:     fmt.Sprintf("existing_pictures[%d]", index),
				Message: "existing picture url not found",
			})
		}
	}
	if len(errList) != 0 {
		return handling.NewErrorWrapper(handling.CodeUnprocessableEntity, "invalid parameter", errList, nil)
	}

	if len(pictures) > 0 {

		responseMedia := handling.BaseResponse[[]response.MediaResponse]{}
		url := fmt.Sprintf("http://%s:%s/api/media", s.cfg.MEDIA_SERVICE_NAME, s.cfg.MEDIA_SERVICE_PORT)
		res, err := service.ForwardFilesToService(url, pictures)
		if err != nil {
			return handling.NewErrorWrapper(handling.CodeServerError, "failed to upload picture to media", nil, nil)
		} else if res.StatusCode < 300 {
			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				return handling.NewErrorWrapper(handling.CodeServerError, "failed to read response http", nil, nil)
			}

			b := bytes.NewBuffer(resBody)
			d := json.NewDecoder(b)
			err = d.Decode(&responseMedia)
			if err != nil {
				return handling.NewErrorWrapper(handling.CodeServerError, "failed to parse response", nil, nil)
			}
		}
		defer res.Body.Close()

		for _, media := range responseMedia.Data {
			m.Pictures = append(m.Pictures, model.ProductPicture{
				Url: media.ID,
			})
		}

	}

	err = s.productRepository.Update(s.db, m)
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeServerError, "failed to update product to db", nil, nil)
	}

	err = s.cache.Set(ctx, cache.GetCacheKey(s.cfg.CACHE_KEY_PRODUCT), nil, constant.CacheTTLInvalidate)
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeServerError, "failed to delete product from cache", nil, nil)
	}

	for _, picture := range m.MergeOutputPictures {
		if picture.Action == "INSERT" {

			for i := 0; i < len(m.Pictures); i++ {
				if m.Pictures[i].ID == 0 && m.Pictures[i].Url == picture.InsertedUrl {
					m.Pictures[i].ID = picture.InsertedID
				}
			}

		}

		if picture.Action == "DELETE" {

			messageBody := response.MediaResponse{
				ID: picture.DeletedUrl,
			}
			messageByte, err := json.Marshal(messageBody)
			if err != nil {
				log.Println("Error marshal message, ", err.Error())
				return handling.NewErrorWrapper(handling.CodeServerError, "error marshal message", nil, nil)
			}

			s.rmq.Publish(
				s.cfg.MEDIA_EXCHANGE,
				"delete_media",
				messageByte,
				nil,
			)

		}
	}

	return nil
}

func (s *productService) DeleteProduct(ctx context.Context, m *model.Product) error {
	return nil
}
