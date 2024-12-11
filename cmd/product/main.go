package main

import (
	"fmt"
	"log"

	"github.com/Ajulll22/belajar-microservice/internal/product/config"
	"github.com/Ajulll22/belajar-microservice/internal/product/router"
	"github.com/Ajulll22/belajar-microservice/pkg/constant"
	"github.com/Ajulll22/belajar-microservice/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./config/product.env")
	if err != nil {
		log.Println(err)
	}
	cfg := config.GetEnv()

	redis := config.RedisClient(cfg)
	db := config.DbConnect(cfg)

	if cfg.APP_ENV == constant.EnvironmentProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())

	validator.RegisterCustomValidation()

	router.Register(r, db, redis, cfg)
	port := fmt.Sprintf(":%s", cfg.APP_PORT)
	r.Run(port)
}
