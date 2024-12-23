package model

import "time"

type User struct {
	ID          int           `gorm:"column:id" json:"id"`
	Username    string        `gorm:"column:username" json:"username"`
	Email       string        `gorm:"column:email" json:"email"`
	PhoneNumber string        `gorm:"column:phone_number" json:"phone_number"`
	DOB         string        `gorm:"column:dob" json:"dob"`
	Password    string        `gorm:"column:password" json:"-"`
	UserType    int           `gorm:"column:user_type" json:"user_type"`
	CreatedAt   time.Time     `gorm:"column:created_at" json:"created_at" `
	UpdatedAt   time.Time     `gorm:"column:updated_at" json:"updated_at"`
	Addresses   []UserAddress `gorm:"-" json:"addresses,omitempty"`
}
