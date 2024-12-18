package main

import (
	"fmt"
	"log"

	"github.com/Ajulll22/belajar-microservice/internal/product/config"
	"github.com/Ajulll22/belajar-microservice/internal/product/router"
	"github.com/Ajulll22/belajar-microservice/pkg/cache"
	"github.com/Ajulll22/belajar-microservice/pkg/constant"
	"github.com/Ajulll22/belajar-microservice/pkg/database"
	"github.com/Ajulll22/belajar-microservice/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./config/global.env", "./config/product.env")
	if err != nil {
		log.Println(err)
	}
	cfg := config.GetEnv()

	redis := cache.RedisClient(cfg.REDIS_HOST, cfg.REDIS_PORT, cfg.REDIS_PASS)
	db := database.SQLConnect(cfg.DB_USER, cfg.DB_PASS, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)

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
