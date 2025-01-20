package repository

import (
	"github.com/Ajulll22/belajar-microservice/internal/product/model"
	"gorm.io/gorm"
)

type ProductPictureRepository interface {
	FindAll(*gorm.DB, *[]model.ProductPicture) error
	FindByProductID(*gorm.DB, *[]model.ProductPicture, int) error
	FindByID(*gorm.DB, *model.ProductPicture, int) error
}

type productPictureRepository struct {
}

func NewProductPictureRepository() ProductPictureRepository {
	return &productPictureRepository{}
}

func (r *productPictureRepository) FindAll(db *gorm.DB, m *[]model.ProductPicture) error {
	query := db.Raw("spMS_product_picture_data").Scan(m)
	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (r *productPictureRepository) FindByProductID(db *gorm.DB, m *[]model.ProductPicture, productID int) error {
	query := db.Raw("spMS_product_picture_data_by_product_id ?", productID).Scan(m)
	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (r *productPictureRepository) FindByID(db *gorm.DB, m *model.ProductPicture, id int) error {
	query := db.Raw("spMS_product_picture_data ?", id).Scan(m)
	if query.Error != nil {
		return query.Error
	}

	return nil
}
