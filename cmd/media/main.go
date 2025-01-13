package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Ajulll22/belajar-microservice/internal/media/config"
	"github.com/Ajulll22/belajar-microservice/internal/media/router"
	"github.com/Ajulll22/belajar-microservice/pkg/broker"
	"github.com/Ajulll22/belajar-microservice/pkg/constant"
	"github.com/Ajulll22/belajar-microservice/pkg/database"
	"github.com/Ajulll22/belajar-microservice/pkg/validator"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./config/global.env", "./config/media.env")
	if err != nil {
		log.Println(err)
	}
	cfg := config.GetEnv()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db := database.MongoConnect(ctx, cfg.DB_USER, cfg.DB_PASS, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	defer db.Client().Disconnect(ctx)

	rmq := broker.RabbitMQConnect(cfg.RABBIT_HOST, cfg.RABBIT_USER, cfg.RABBIT_PASS, cfg.RABBIT_PORT)
	defer rmq.Close()

	if cfg.APP_ENV == constant.EnvironmentProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())

	validator.RegisterCustomValidation()

	router.Register(r, db, cfg, rmq)
	router.RegisterConsumer(db, cfg, rmq)

	port := fmt.Sprintf(":%s", cfg.MEDIA_SERVICE_PORT)
	r.Run(port)
}
