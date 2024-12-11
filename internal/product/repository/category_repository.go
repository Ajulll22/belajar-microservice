package repository

import (
	"github.com/Ajulll22/belajar-microservice/internal/product/model"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll(*gorm.DB, *[]model.Category) error
	FindByID(*gorm.DB, *model.Category, int) error
	Insert(*gorm.DB, *model.Category) error
	Update(*gorm.DB, *model.Category) error
	Destroy(*gorm.DB, *model.Category) error
}

type categoryRepository struct {
}

func NewCategoryRepository() CategoryRepository {
	return &categoryRepository{}
}

func (r *categoryRepository) FindAll(db *gorm.DB, m *[]model.Category) error {
	query := db.Raw("spMS_category_data 0").Scan(m)
	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (r *categoryRepository) FindByID(db *gorm.DB, m *model.Category, id int) error {
	query := db.Raw("spMS_category_data ?", id).Scan(m)
	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (r *categoryRepository) Insert(db *gorm.DB, m *model.Category) error {
	query := db.Raw("spMS_category_data_insert ?", m.Name).Scan(m)
	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (r *categoryRepository) Update(db *gorm.DB, m *model.Category) error {
	query := db.Exec("spMS_category_data_update ?, ?", m.ID, m.Name)
	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (r *categoryRepository) Destroy(db *gorm.DB, m *model.Category) error {
	query := db.Exec("spMS_category_data_delete ?", m.ID)
	if query.Error != nil {
		return query.Error
	}

	return nil
}
