package service

import (
	"context"
	"time"

	"github.com/Ajulll22/belajar-microservice/internal/user/config"
	"github.com/Ajulll22/belajar-microservice/internal/user/dto/response"
	"github.com/Ajulll22/belajar-microservice/internal/user/model"
	"github.com/Ajulll22/belajar-microservice/internal/user/repository"
	"github.com/Ajulll22/belajar-microservice/pkg/cache"
	"github.com/Ajulll22/belajar-microservice/pkg/constant"
	"github.com/Ajulll22/belajar-microservice/pkg/handling"
	"github.com/Ajulll22/belajar-microservice/pkg/validator"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(ctx context.Context, m *response.AuthResponse, username string, password string) error
	RefreshToken(ctx context.Context, m *response.AuthResponse, refreshToken string) error
	Logout(ctx context.Context, refreshToken string) error
}

func NewAuthService(cfg config.Config, db *gorm.DB, cache cache.Cache, userRepository repository.UserRepository) AuthService {
	return &authService{cfg, db, cache, userRepository}
}

type authService struct {
	cfg            config.Config
	db             *gorm.DB
	cache          cache.Cache
	userRepository repository.UserRepository
}

func (s *authService) Login(ctx context.Context, m *response.AuthResponse, username string, password string) error {
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
		UserID:        userData.ID,
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

	s.cache.Set(ctx, cache.GetCacheKey(s.cfg.CACHE_KEY_USER, td.RefreshUUID), td, td.RefreshExpiry.Sub(time.Time{}))

	m.AccessToken = td.AccessToken
	m.RefreshToken = td.RefreshToken

	return nil
}

func (s *authService) RefreshToken(ctx context.Context, m *response.AuthResponse, refreshToken string) error {
	userData := model.User{}
	td := model.TokenDetails{}
	errList := []validator.ErrorValidator{}

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(s.cfg.REFRESH_SECRET), nil
	})
	if err != nil {
		errList = append(errList, validator.ErrorValidator{
			Key:     "refresh_token",
			Message: "invalid or expired refresh token",
		})
		return handling.NewErrorWrapper(handling.CodeClientError, "invalid or expired refresh token", errList, nil)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["uuid"] == nil {
		errList = append(errList, validator.ErrorValidator{
			Key:     "refresh_token",
			Message: "invalid token claims",
		})
		return handling.NewErrorWrapper(handling.CodeClientError, "invalid token claims", errList, nil)
	}

	refreshUUID := claims["uuid"].(string)
	err = s.cache.Get(ctx, cache.GetCacheKey(s.cfg.CACHE_KEY_USER, refreshUUID), &td)
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeServerError, "failed to get token from cache", nil, nil)
	}
	if td.UserID == 0 {
		errList = append(errList, validator.ErrorValidator{
			Key:     "refresh_token",
			Message: "invalid refersh token",
		})
		return handling.NewErrorWrapper(handling.CodeClientError, "invalid refersh token", errList, nil)
	}

	err = s.userRepository.FindByID(s.db, &userData, td.UserID)
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeServerError, "failed to get user from db", nil, nil)
	}
	if userData.ID == 0 {
		errList = append(errList, validator.ErrorValidator{
			Key:     "refresh_token",
			Message: "refresh token user not found",
		})
		return handling.NewErrorWrapper(handling.CodeClientError, "refresh token user not found", errList, nil)
	}

	td = model.TokenDetails{
		UserID:        userData.ID,
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
	newAccessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	td.AccessToken, err = newAccessToken.SignedString([]byte(s.cfg.ACCESS_SECRET))
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeServerError, "failed to create access token", nil, nil)
	}

	// Create refresh token
	refreshClaims := jwt.MapClaims{
		"user_id": userData.ID,
		"exp":     td.RefreshExpiry.Unix(),
		"uuid":    td.RefreshUUID,
	}
	newRefreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	td.RefreshToken, err = newRefreshToken.SignedString([]byte(s.cfg.REFRESH_SECRET))
	if err != nil {
		return handling.NewErrorWrapper(handling.CodeServerError, "failed to create refresh token", nil, nil)
	}

	s.cache.Set(ctx, cache.GetCacheKey(s.cfg.CACHE_KEY_USER, refreshUUID), nil, constant.CacheTTLInvalidate)
	s.cache.Set(ctx, cache.GetCacheKey(s.cfg.CACHE_KEY_USER, td.RefreshUUID), td, td.RefreshExpiry.Sub(time.Time{}))

	m.AccessToken = td.AccessToken
	m.RefreshToken = td.RefreshToken

	return nil
}

func (s *authService) Logout(ctx context.Context, refreshToken string) error {
	errList := []validator.ErrorValidator{}

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(s.cfg.REFRESH_SECRET), nil
	})
	if err != nil {
		errList = append(errList, validator.ErrorValidator{
			Key:     "refresh_token",
			Message: "invalid or expired refresh token",
		})
		return handling.NewErrorWrapper(handling.CodeClientError, "invalid or expired refresh token", errList, err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["uuid"] == nil {
		errList = append(errList, validator.ErrorValidator{
			Key:     "refresh_token",
			Message: "invalid token claims",
		})
		return handling.NewErrorWrapper(handling.CodeClientError, "invalid token claims", errList, err)
	}

	refreshUUID := claims["uuid"].(string)
	s.cache.Set(ctx, cache.GetCacheKey(s.cfg.CACHE_KEY_USER, refreshUUID), nil, constant.CacheTTLInvalidate)

	return nil
}
