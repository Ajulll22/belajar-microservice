package constant

import "os"

type GlobalConfig struct {
	GATEWAY_SERVICE_NAME string
	GATEWAY_SERVICE_PORT string
	PRODUCT_SERVICE_NAME string
	PRODUCT_SERVICE_PORT string
	MEDIA_SERVICE_NAME   string
	MEDIA_SERVICE_PORT   string
	USER_SERVICE_NAME    string
	USER_SERVICE_PORT    string
	APP_ENV              string
	HOST                 string
	ACCESS_SECRET        string
	REFRESH_SECRET       string
	RABBIT_HOST          string
	RABBIT_PORT          string
	RABBIT_USER          string
	RABBIT_PASS          string
	MEDIA_EXCHANGE       string
	MEDIA_QUEUE          string
	LOGGER_EXCHANGE      string
	LOGGER_QUEUE         string
}

func GetGlobalConfig() GlobalConfig {
	return GlobalConfig{
		GATEWAY_SERVICE_NAME: os.Getenv("GATEWAY_SERVICE_NAME"),
		GATEWAY_SERVICE_PORT: os.Getenv("GATEWAY_SERVICE_PORT"),
		PRODUCT_SERVICE_NAME: os.Getenv("PRODUCT_SERVICE_NAME"),
		PRODUCT_SERVICE_PORT: os.Getenv("PRODUCT_SERVICE_PORT"),
		MEDIA_SERVICE_NAME:   os.Getenv("MEDIA_SERVICE_NAME"),
		MEDIA_SERVICE_PORT:   os.Getenv("MEDIA_SERVICE_PORT"),
		USER_SERVICE_NAME:    os.Getenv("USER_SERVICE_NAME"),
		USER_SERVICE_PORT:    os.Getenv("USER_SERVICE_PORT"),
		APP_ENV:              os.Getenv("APP_ENV"),
		HOST:                 os.Getenv("HOST"),
		ACCESS_SECRET:        os.Getenv("ACCESS_SECRET"),
		REFRESH_SECRET:       os.Getenv("REFRESH_SECRET"),
		RABBIT_HOST:          os.Getenv("RABBIT_HOST"),
		RABBIT_PORT:          os.Getenv("RABBIT_PORT"),
		RABBIT_USER:          os.Getenv("RABBIT_USER"),
		RABBIT_PASS:          os.Getenv("RABBIT_PASS"),
		MEDIA_EXCHANGE:       os.Getenv("MEDIA_EXCHANGE"),
		MEDIA_QUEUE:          os.Getenv("MEDIA_QUEUE"),
		LOGGER_EXCHANGE:      os.Getenv("LOGGER_EXCHANGE"),
		LOGGER_QUEUE:         os.Getenv("LOGGER_QUEUE"),
	}
}
