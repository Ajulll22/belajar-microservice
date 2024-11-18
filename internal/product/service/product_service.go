package service

import (
	"context"

	"github.com/Ajulll22/belajar-microservice/config"
	"github.com/Ajulll22/belajar-microservice/internal/product/model"
	"github.com/Ajulll22/belajar-microservice/internal/product/repository"
	"github.com/Ajulll22/belajar-microservice/pkg/cache"
)

type ProductService interface {
	GetProducts(ctx context.Context, m []*model.Product) error
	GetProduct(ctx context.Context, m *model.Product, id int) error
	CreateProduct(ctx context.Context, m *model.Product) error
	UpdateProduct(ctx context.Context, m *model.Product) error
	DeleteProduct(ctx context.Context, m *model.Product) error
}

func NewProductService(cfg config.Config, cache cache.Cache, productRepository repository.ProductRepository) ProductService {
	return &productService{cfg, cache, productRepository}
}

type productService struct {
	cfg               config.Config
	cache             cache.Cache
	productRepository repository.ProductRepository
}

func (s *productService) GetProducts(ctx context.Context, m []*model.Product) error {
	return nil
}

func (s *productService) GetProduct(ctx context.Context, m *model.Product, id int) error {
	return nil
}

func (s *productService) CreateProduct(ctx context.Context, m *model.Product) error {
	return nil
}

func (s *productService) UpdateProduct(ctx context.Context, m *model.Product) error {
	return nil
}

func (s *productService) DeleteProduct(ctx context.Context, m *model.Product) error {
	return nil
}
