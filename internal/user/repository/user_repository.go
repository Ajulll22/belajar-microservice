package repository

import (
	"github.com/Ajulll22/belajar-microservice/internal/user/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(*gorm.DB, *model.User, string) error
	FindByID(*gorm.DB, *model.User, int) error
}

type userRepository struct {
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindByUsername(db *gorm.DB, m *model.User, username string) error {
	query := db.Raw("spMS_user_data_by_username ?", username).Scan(m)

	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (r *userRepository) FindByID(db *gorm.DB, m *model.User, id int) error {
	query := db.Raw("spMS_user_data ?", id).Scan(m)

	if query.Error != nil {
		return query.Error
	}

	return nil
}
