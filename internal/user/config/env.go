package config

import (
	"os"

	"github.com/Ajulll22/belajar-microservice/pkg/constant"
)

type Config struct {
	APP_ENV        string
	APP_PORT       string
	DB_HOST        string
	DB_PORT        string
	DB_USER        string
	DB_PASS        string
	DB_NAME        string
	REDIS_HOST     string
	REDIS_PORT     string
	REDIS_PASS     string
	CACHE_KEY_USER string
	constant.GlobalConfig
}

func GetEnv() Config {
	return Config{
		APP_ENV:        os.Getenv("APP_ENV"),
		APP_PORT:       os.Getenv("APP_PORT"),
		DB_HOST:        os.Getenv("DB_HOST"),
		DB_PORT:        os.Getenv("DB_PORT"),
		DB_USER:        os.Getenv("DB_USER"),
		DB_PASS:        os.Getenv("DB_PASS"),
		DB_NAME:        os.Getenv("DB_NAME"),
		REDIS_HOST:     os.Getenv("REDIS_HOST"),
		REDIS_PORT:     os.Getenv("REDIS_PORT"),
		REDIS_PASS:     os.Getenv("REDIS_PASS"),
		CACHE_KEY_USER: os.Getenv("CACHE_KEY_USER"),

		GlobalConfig: constant.GetGlobalConfig(),
	}
}
