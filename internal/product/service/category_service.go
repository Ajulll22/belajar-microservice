package service

import (
	"context"

	"github.com/Ajulll22/belajar-microservice/internal/product/config"
	"github.com/Ajulll22/belajar-microservice/internal/product/model"
	"github.com/Ajulll22/belajar-microservice/internal/product/repository"
	"github.com/Ajulll22/belajar-microservice/pkg/cache"
	"github.com/Ajulll22/belajar-microservice/pkg/constant"
	"github.com/Ajulll22/belajar-microservice/pkg/handling"
	"gorm.io/gorm"
)

type CategoryService interface {
	GetCategories(ctx context.Context, m *[]model.Category) error
	GetCategory(ctx context.Context, m *model.Category, id int) error
	CreateCategory(ctx context.Context, m *model.Category) error
	UpdateCategory(ctx context.Context, m *model.Category) error
	DeleteCategory(ctx context.Context, m *model.Category) error
}

func NewCategoryService(cfg config.Config, db *gorm.DB, cache cache.Cache, categoryRepository repository.CategoryRepository) CategoryService {
	return &categoryService{cfg, db, cache, categoryRepository}
}

type categoryService struct {
	cfg                config.Config
	db                 *gorm.DB
	cache              cache.Cache
	categoryRepository repository.CategoryRepository
}

func (s *categoryService) GetCategories(ctx context.Context, m *[]model.Category) error {
	err := s.cache.Get(ctx, cache.GetCacheKey(s.cfg.CACHE_KEY_CATEGORY), m)
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeServerError, "failed to get categories from cache", nil, nil)
	}

	if len(*m) == 0 {

		err := s.categoryRepository.FindAll(s.db, m)
		if err != nil {
			return handling.NewErrorWrapper(handling.CodeServerError, "failed to get categories from db", nil, nil)
		}

		err = s.cache.Set(ctx, cache.GetCacheKey(s.cfg.CACHE_KEY_PRODUCT), *m, constant.CacheTTLOneDay)
		if err != nil {
			return handling.NewErrorWrapper(handling.CodeServerError, "failed to set categories to cache", nil, nil)
		}

	}
	return nil
}

func (s *categoryService) GetCategory(ctx context.Context, m *model.Category, id int) error {
	categories := []model.Category{}
	err := s.cache.Get(ctx, cache.GetCacheKey(s.cfg.CACHE_KEY_CATEGORY), &categories)
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeServerError, "failed to get categories from cache", nil, nil)
	}

	if len(categories) != 0 {

		for _, val := range categories {
			if val.ID == id {
				*m = val

				return nil
			}

		}

	}

	err = s.categoryRepository.FindByID(s.db, m, id)
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeServerError, "failed to get category from db", nil, nil)
	}

	if m.ID == 0 {
		return handling.NewErrorWrapper(handling.CodeNotFoundError, "data category not found", nil, nil)
	}

	return nil
}

func (s *categoryService) CreateCategory(ctx context.Context, m *model.Category) error {
	return nil
}

func (s *categoryService) UpdateCategory(ctx context.Context, m *model.Category) error {
	return nil
}

func (s *categoryService) DeleteCategory(ctx context.Context, m *model.Category) error {
	return nil
}
