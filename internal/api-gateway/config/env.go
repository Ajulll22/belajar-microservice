package config

import (
	"github.com/Ajulll22/belajar-microservice/pkg/constant"
)

type Config struct {
	constant.GlobalConfig
}

func GetEnv() Config {
	return Config{
		GlobalConfig: constant.GetGlobalConfig(),
	}
}
