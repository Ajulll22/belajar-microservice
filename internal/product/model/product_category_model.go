package model

type ProductCategory struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ProductID int    `json:"-" gorm:"column:product_id"`
}
