package model

type ProductPicture struct {
	ID        int    `json:"id" gorm:"column:id"`
	Url       string `json:"url" gorm:"column:url"`
	ProductID int    `json:"-" gorm:"column:product_id"`
}

type MergeOutputProductPicture struct {
	Action      string `gorm:"column:action"`
	InsertedID  int    `gorm:"column:inserted_id"`
	InsertedUrl string `gorm:"column:inserted_url"`
	DeletedID   int    `gorm:"column:deleted_id"`
	DeletedUrl  string `gorm:"column:deleted_url"`
}
