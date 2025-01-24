package config

import (
	"os"
	"strconv"

	"github.com/Ajulll22/belajar-microservice/pkg/constant"
)

type Config struct {
	DB_HOST          string
	DB_PORT          string
	DB_USER          string
	DB_PASS          string
	DB_NAME          string
	RABBIT_HOST      string
	RABBIT_PORT      string
	RABBIT_USER      string
	RABBIT_PASS      string
	RABBIT_MAX_RETRY int
	constant.GlobalConfig
}

func GetEnv() Config {
	RABBIT_MAX_RETRY, _ := strconv.Atoi(os.Getenv("RABBIT_MAX_RETRY"))
	return Config{
		DB_HOST:          os.Getenv("DB_HOST"),
		DB_PORT:          os.Getenv("DB_PORT"),
		DB_USER:          os.Getenv("DB_USER"),
		DB_PASS:          os.Getenv("DB_PASS"),
		DB_NAME:          os.Getenv("DB_NAME"),
		RABBIT_HOST:      os.Getenv("RABBIT_HOST"),
		RABBIT_PORT:      os.Getenv("RABBIT_PORT"),
		RABBIT_USER:      os.Getenv("RABBIT_USER"),
		RABBIT_PASS:      os.Getenv("RABBIT_PASS"),
		RABBIT_MAX_RETRY: RABBIT_MAX_RETRY,

		GlobalConfig: constant.GetGlobalConfig(),
	}
}
