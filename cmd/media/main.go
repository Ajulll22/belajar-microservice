package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Ajulll22/belajar-microservice/internal/media/config"
	"github.com/Ajulll22/belajar-microservice/internal/media/router"
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

	if cfg.APP_ENV == constant.EnvironmentProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())

	validator.RegisterCustomValidation()

	router.Register(r, db, cfg)
	port := fmt.Sprintf(":%s", cfg.APP_PORT)
	r.Run(port)
}
