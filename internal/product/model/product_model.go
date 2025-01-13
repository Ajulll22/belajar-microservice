package model

import "time"

type Product struct {
	ID                int               `json:"id" gorm:"column:id"`
	Name              string            `json:"name" gorm:"column:name"`
	Price             int               `json:"price" gorm:"column:price"`
	Stock             int               `json:"stock" gorm:"column:stock"`
	Description       string            `json:"description" gorm:"column:description"`
	CreatedAt         time.Time         `json:"created_at" gorm:"column:created_at"`
	UpdatedAt         time.Time         `json:"updated_at" gorm:"column:updated_at"`
	Pictures          []ProductPicture  `json:"pictures" gorm:"foreignKey:ProductID"`
	Categories        []ProductCategory `json:"categories" gorm:"foreignKey:ProductID"`
	DeletedPictures   []ProductPicture  `json:"-" gorm:"foreignKey:ProductID"`
	DeletedCategories []ProductCategory `json:"-" gorm:"foreignKey:ProductID"`
}
