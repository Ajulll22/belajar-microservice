package model

type ProductPicture struct {
	ID  int    `json:"id" gorm:"column:id"`
	Url string `json:"url" gorm:"column:url"`
}
