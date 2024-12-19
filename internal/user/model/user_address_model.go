package model

type UserAddress struct {
	ID      int    `gorm:"column:id" json:"id"`
	Name    string `gorm:"column:name" json:"name"`
	Address string `gorm:"column:address" json:"address"`
	Note    string `gorm:"column:note" json:"note"`
	UserID  int    `gorm:"column:user_id" json:"-"`
}
