package service

import (
	"context"
	"time"

	"github.com/Ajulll22/belajar-microservice/internal/user/config"
	"github.com/Ajulll22/belajar-microservice/internal/user/dto/response"
	"github.com/Ajulll22/belajar-microservice/internal/user/model"
	"github.com/Ajulll22/belajar-microservice/internal/user/repository"
	"github.com/Ajulll22/belajar-microservice/pkg/cache"
	"github.com/Ajulll22/belajar-microservice/pkg/handling"
	"github.com/Ajulll22/belajar-microservice/pkg/validator"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	userData := model.User{}
	errList := []validator.ErrorValidator{}

	err := s.userRepository.FindByUsername(s.db, &userData, username)
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeServerError, "failed to get user from db", nil, err)
	}

	if userData.Username != username {
		errList = append(errList, validator.ErrorValidator{
			Key:     "username",
			Message: "username not found",
		})
		return handling.NewErrorWrapper(handling.CodeClientError, "username not found", errList, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(password))
	if err != nil {
		errList = append(errList, validator.ErrorValidator{
			Key:     "password",
			Message: "invalid password",
		})
		return handling.NewErrorWrapper(handling.CodeClientError, "invalid password", errList, err)
	}

	td := model.TokenDetails{
		AccessUUID:    uuid.New().String(),
		RefreshUUID:   uuid.New().String(),
		AccessExpiry:  time.Now().Add(15 * time.Minute),
		RefreshExpiry: time.Now().Add(7 * 24 * time.Hour),
	}

	// Create access token
	accessClaims := jwt.MapClaims{
		"authorized": true,
		"user_id":    userData.ID,
		"exp":        td.AccessExpiry.Unix(),
		"uuid":       td.AccessUUID,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	td.AccessToken, err = accessToken.SignedString([]byte(s.cfg.ACCESS_SECRET))
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeServerError, "failed to create access token", nil, err)
	}

	// Create refresh token
	refreshClaims := jwt.MapClaims{
		"user_id": userData.ID,
		"exp":     td.RefreshExpiry.Unix(),
		"uuid":    td.RefreshUUID,
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	td.RefreshToken, err = refreshToken.SignedString([]byte(s.cfg.REFRESH_SECRET))
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeServerError, "failed to create refresh token", nil, err)
	}

	s.cache.Set(ctx, cache.GetCacheKey(s.cfg.CACHE_KEY_USER, userData.ID), td, td.RefreshExpiry.Sub(time.Time{}))

	m.AccessToken = td.AccessToken
	m.RefreshToken = td.RefreshToken

	return nil
}

func (s *userService) RefreshToken(ctx context.Context, m *response.AuthResponse, refreshToken string) error {
	return nil
}

func (s *userService) Logout(ctx context.Context, refreshToken string) error {
	return nil
}
