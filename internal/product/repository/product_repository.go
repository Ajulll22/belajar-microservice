package repository

import (
	"time"

	"github.com/Ajulll22/belajar-microservice/internal/product/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll([]*model.Product) error
	FindByID(*model.Product, int) error
	Insert(*model.Product) error
	Update(*model.Product) error
	Destroy(*model.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) FindAll(m []*model.Product) error {
	rawData := productRawData{}

	query := r.db.Raw("").Scan(&rawData)

	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (r *productRepository) FindByID(m *model.Product, id int) error {
	query := r.db.Raw("").Scan(m)

	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (r *productRepository) Insert(m *model.Product) error {
	query := r.db.Raw("").Scan(m)

	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (r *productRepository) Update(m *model.Product) error {
	query := r.db.Exec("")

	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (r *productRepository) Destroy(m *model.Product) error {
	query := r.db.Exec("")

	if query.Error != nil {
		return query.Error
	}

	return nil
}

type productRawData struct {
	ProductID          int       `gorm:"column:product_id"`
	ProductName        string    `gorm:"column:product_name"`
	ProductPrice       int       `gorm:"column:product_price"`
	ProductStock       int       `gorm:"column:product_stock"`
	ProductDescription string    `gorm:"column:product_description"`
	PictureID          int       `gorm:"column:picture_id"`
	PictureUrl         string    `gorm:"column:picture_url"`
	CategoryID         int       `gorm:"column:category_id"`
	CategoryName       string    `gorm:"column:category_name"`
	CreatedAt          time.Time `gorm:"column:created_at"`
	UpdatedAt          time.Time `gorm:"column:updated_at"`
}
