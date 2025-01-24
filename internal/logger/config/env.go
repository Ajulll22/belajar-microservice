package config

import (
	"os"
	"strconv"

	"github.com/Ajulll22/belajar-microservice/pkg/constant"
)

type Config struct {
	ELASTIC_PROTOCOL      string
	ELASTIC_HOST          string
	ELASTIC_PORT          string
	ELASTIC_USER          string
	ELASTIC_PASS          string
	RABBIT_HOST           string
	RABBIT_PORT           string
	RABBIT_USER           string
	RABBIT_PASS           string
	RABBIT_MAX_RETRY      int
	APPLICATION_LOG_INDEX string
	AUDIT_LOG_INDEX       string
	PERFORM_LOG_INDEX     string
	ERROR_LOG_INDEX       string
	constant.GlobalConfig
}

func GetEnv() Config {
	RABBIT_MAX_RETRY, _ := strconv.Atoi(os.Getenv("RABBIT_MAX_RETRY"))
	return Config{
		ELASTIC_PROTOCOL:      os.Getenv("ELASTIC_PROTOCOL"),
		ELASTIC_HOST:          os.Getenv("ELASTIC_HOST"),
		ELASTIC_PORT:          os.Getenv("ELASTIC_PORT"),
		ELASTIC_USER:          os.Getenv("ELASTIC_USER"),
		ELASTIC_PASS:          os.Getenv("ELASTIC_PASS"),
		RABBIT_HOST:           os.Getenv("RABBIT_HOST"),
		RABBIT_PORT:           os.Getenv("RABBIT_PORT"),
		RABBIT_USER:           os.Getenv("RABBIT_USER"),
		RABBIT_PASS:           os.Getenv("RABBIT_PASS"),
		RABBIT_MAX_RETRY:      RABBIT_MAX_RETRY,
		APPLICATION_LOG_INDEX: os.Getenv("APPLICATION_LOG_INDEX"),
		AUDIT_LOG_INDEX:       os.Getenv("AUDIT_LOG_INDEX"),
		PERFORM_LOG_INDEX:     os.Getenv("PERFORM_LOG_INDEX"),
		ERROR_LOG_INDEX:       os.Getenv("ERROR_LOG_INDEX"),

		GlobalConfig: constant.GetGlobalConfig(),
	}
}
