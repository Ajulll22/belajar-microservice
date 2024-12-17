package constant

import "os"

type GlobalConfig struct {
	PRODUCT_SERVICE_PORT string
	MEDIA_SERVICE_PORT   string
	HOST                 string
}

func GetGlobalConfig() GlobalConfig {
	return GlobalConfig{
		PRODUCT_SERVICE_PORT: os.Getenv("PRODUCT_SERVICE_PORT"),
		MEDIA_SERVICE_PORT:   os.Getenv("MEDIA_SERVICE_PORT"),
		HOST:                 os.Getenv("HOST"),
	}
}