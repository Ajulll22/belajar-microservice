package main

import (
	"fmt"
	"log"

	"github.com/Ajulll22/belajar-microservice/internal/api-gateway/config"
	"github.com/Ajulll22/belajar-microservice/internal/api-gateway/router"
	"github.com/Ajulll22/belajar-microservice/pkg/constant"
	"github.com/Ajulll22/belajar-microservice/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./config/global.env")
	if err != nil {
		log.Println(err)
	}
	cfg := config.GetEnv()

	if cfg.APP_ENV == constant.EnvironmentProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())

	validator.RegisterCustomValidation()

	router.Register(r, cfg)
	port := fmt.Sprintf(":%s", cfg.GATEWAY_SERVICE_PORT)
	r.Run(port)
}
