package model

type ProductCategory struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	ProductID int    `json:"-" gorm:"column:product_id"`
}

type MergeOutputProductCategory struct {
	Action       string `gorm:"column:action"`
	InsertedID   int    `gorm:"column:inserted_id"`
	InsertedName string `gorm:"column:inserted_name"`
	DeletedID    int    `gorm:"column:deleted_id"`
	DeletedName  string `gorm:"column:deleted_name"`
}
