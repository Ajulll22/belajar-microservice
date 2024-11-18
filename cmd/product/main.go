package main

import (
	"fmt"
	"log"

	"github.com/Ajulll22/belajar-microservice/config"
	"github.com/Ajulll22/belajar-microservice/internal/product/router"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}
	cfg := config.GetEnv()

	redis := config.RedisClient(cfg)
	db := config.DbConnect(cfg)

	r := gin.New()
	r.Use(gin.Recovery())

	router.Register(r, db, redis, cfg)
	port := fmt.Sprintf(":%s", cfg.APP_PORT)
	r.Run(port)
}
