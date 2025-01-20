package repository

import (
	"encoding/json"

	"github.com/Ajulll22/belajar-microservice/internal/product/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(*gorm.DB, *[]model.Product) error
	FindByID(*gorm.DB, *model.Product, int) error
	Insert(*gorm.DB, *model.Product) error
	Update(*gorm.DB, *model.Product) error
	Destroy(*gorm.DB, *model.Product) error
}

type productRepository struct {
}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (r *productRepository) FindAll(db *gorm.DB, m *[]model.Product) error {
	rawData := model.ProductRawData{}

	query := db.Raw("spMS_product_data 0").Scan(&rawData)

	if query.Error != nil {
		return query.Error
	}

	productMap := rawData.MapDataToStruct()

	// Convert map to slice
	for _, product := range productMap {
		*m = append(*m, *product)
	}

	return nil
}

func (r *productRepository) FindByID(db *gorm.DB, m *model.Product, id int) error {
	rawData := model.ProductRawData{}

	query := db.Raw("spMS_product_data ?", id).Scan(&rawData)

	if query.Error != nil {
		return query.Error
	}

	productMap := rawData.MapDataToStruct()

	for _, product := range productMap {
		*m = *product
	}

	return nil
}

func (r *productRepository) Insert(db *gorm.DB, m *model.Product) error {

	data := model.Product{}
	query := db.Raw("spMS_product_data_insert ?, ?, ?, ?", m.Name, m.Price, m.Stock, m.Description).Scan(&data)

	if query.Error != nil {
		return query.Error
	}

	m.ID = data.ID
	m.CreatedAt = data.CreatedAt
	m.UpdatedAt = data.UpdatedAt

	array_url := []string{}
	for _, val := range m.Pictures {
		array_url = append(array_url, val.Url)
	}
	if len(array_url) > 0 {
		string_url, err := json.Marshal(array_url)
		if err != nil {
			return err
		}
		query := db.Raw("spMS_product_picture_data_insert ?, ?", m.ID, string(string_url)).Scan(&(m.Pictures))
		if query.Error != nil {
			return query.Error
		}
	}

	array_category_id := []int{}
	for _, val := range m.Categories {
		array_category_id = append(array_category_id, val.ID)
	}
	if len(array_category_id) > 0 {
		string_category_id, err := json.Marshal(array_category_id)
		if err != nil {
			return err
		}
		query := db.Raw("spMS_product_category_data_insert ?, ?", m.ID, string(string_category_id)).Scan(&(m.Categories))
		if query.Error != nil {
			return query.Error
		}
	}

	return nil
}

func (r *productRepository) Update(db *gorm.DB, m *model.Product) error {
	query := db.Exec("spMS_product_data_update ?, ?, ?, ?, ?", m.ID, m.Name, m.Price, m.Stock, m.Description)

	if query.Error != nil {
		return query.Error
	}

	array_url := []string{}
	for _, val := range m.Pictures {
		array_url = append(array_url, val.Url)
	}
	if len(array_url) > 0 {
		string_url, err := json.Marshal(array_url)
		if err != nil {
			return err
		}
		query = db.Raw("spMS_product_picture_data_update ?, ?", m.ID, string(string_url)).Scan(&(m.MergeOutputPictures))
		if query.Error != nil {
			return query.Error
		}
	}

	array_category_id := []int{}
	for _, val := range m.Categories {
		array_category_id = append(array_category_id, val.ID)
	}
	if len(array_category_id) > 0 {
		string_category_id, err := json.Marshal(array_category_id)
		if err != nil {
			return err
		}
		query = db.Raw("spMS_product_category_data_update ?, ?", m.ID, string(string_category_id)).Scan(&(m.MergeOutputCategories))
		if query.Error != nil {
			return query.Error
		}
	}

	return nil
}

func (r *productRepository) Destroy(db *gorm.DB, m *model.Product) error {
	query := db.Exec("spMS_product_data_delete ?", m.ID)

	if query.Error != nil {
		return query.Error
	}

	return nil
}
