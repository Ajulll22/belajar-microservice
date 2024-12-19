package service

import (
	"context"

	"github.com/Ajulll22/belajar-microservice/internal/user/config"
	"github.com/Ajulll22/belajar-microservice/internal/user/dto/response"
	"github.com/Ajulll22/belajar-microservice/internal/user/repository"
	"github.com/Ajulll22/belajar-microservice/pkg/cache"
	"gorm.io/gorm"
)

type UserService interface {
	Login(ctx context.Context, m *response.AuthResponse, username string, password string) error
	RefreshToken(ctx context.Context, m *response.AuthResponse, refreshToken string) error
	Logout(ctx context.Context, refreshToken string) error
}

func NewUserService(cfg config.Config, db *gorm.DB, cache cache.Cache, userRepository repository.UserRepository) UserService {
	return &userService{cfg, db, cache, userRepository}
}

type userService struct {
	cfg            config.Config
	db             *gorm.DB
	cache          cache.Cache
	userRepository repository.UserRepository
}

func (s *userService) Login(ctx context.Context, m *response.AuthResponse, username string, password string) error {
	return nil
}

func (s *userService) RefreshToken(ctx context.Context, m *response.AuthResponse, refreshToken string) error {
	return nil
}

func (s *userService) Logout(ctx context.Context, refreshToken string) error {
	return nil
}
