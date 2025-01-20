package model

import "time"

type Product struct {
	ID                    int                          `json:"id" gorm:"column:id"`
	Name                  string                       `json:"name" gorm:"column:name"`
	Price                 int                          `json:"price" gorm:"column:price"`
	Stock                 int                          `json:"stock" gorm:"column:stock"`
	Description           string                       `json:"description" gorm:"column:description"`
	CreatedAt             time.Time                    `json:"created_at" gorm:"column:created_at"`
	UpdatedAt             time.Time                    `json:"updated_at" gorm:"column:updated_at"`
	Pictures              []ProductPicture             `json:"pictures" gorm:"foreignKey:ProductID"`
	Categories            []ProductCategory            `json:"categories" gorm:"foreignKey:ProductID"`
	MergeOutputPictures   []MergeOutputProductPicture  `json:"-" gorm:"-"`
	MergeOutputCategories []MergeOutputProductCategory `json:"-" gorm:"-"`
}

type ProductRawData []struct {
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

// Map data to nested structures
func (rawData ProductRawData) MapDataToStruct() map[int]*Product {
	// Map data to nested structures
	productMap := make(map[int]*Product)
	for _, item := range rawData {
		// Ensure the product exists in the map
		if _, exists := productMap[item.ProductID]; !exists {
			productMap[item.ProductID] = &Product{
				ID:          item.ProductID,
				Name:        item.ProductName,
				Price:       item.ProductPrice,
				Stock:       item.ProductStock,
				Description: item.ProductDescription,
				CreatedAt:   item.CreatedAt,
				UpdatedAt:   item.UpdatedAt,
				Categories:  []ProductCategory{},
				Pictures:    []ProductPicture{},
			}
		}

		product := productMap[item.ProductID]

		// Add category if it exists
		if item.CategoryID != 0 {
			exists := false
			for _, category := range product.Categories {
				if category.ID == item.CategoryID {
					exists = true
					break
				}
			}
			if !exists {
				product.Categories = append(product.Categories, ProductCategory{
					ID:   item.CategoryID,
					Name: item.CategoryName,
				})
			}
		}

		// Add picture if it exists
		if item.PictureID != 0 {
			exists := false
			for _, picture := range product.Pictures {
				if picture.ID == item.PictureID {
					exists = true
					break
				}
			}
			if !exists {
				product.Pictures = append(product.Pictures, ProductPicture{
					ID:  item.PictureID,
					Url: item.PictureUrl,
				})
			}
		}
	}

	return productMap
}
