package config

import (
	"os"
)

type Config struct {
	APP_ENV            string
	APP_PORT           string
	DB_HOST            string
	DB_PORT            string
	DB_USER            string
	DB_PASS            string
	DB_NAME            string
	REDIS_HOST         string
	REDIS_PORT         string
	REDIS_PASS         string
	CACHE_KEY_PRODUCT  string
	CACHE_KEY_CATEGORY string
}

func GetEnv() Config {
	return Config{
		APP_ENV:            os.Getenv("APP_ENV"),
		APP_PORT:           os.Getenv("APP_PORT"),
		DB_HOST:            os.Getenv("DB_HOST"),
		DB_PORT:            os.Getenv("DB_PORT"),
		DB_USER:            os.Getenv("DB_USER"),
		DB_PASS:            os.Getenv("DB_PASS"),
		DB_NAME:            os.Getenv("DB_NAME"),
		REDIS_HOST:         os.Getenv("REDIS_HOST"),
		REDIS_PORT:         os.Getenv("REDIS_PORT"),
		REDIS_PASS:         os.Getenv("REDIS_PASS"),
		CACHE_KEY_PRODUCT:  os.Getenv("CACHE_KEY_PRODUCT"),
		CACHE_KEY_CATEGORY: os.Getenv("CACHE_KEY_CATEGORY"),
	}
}
