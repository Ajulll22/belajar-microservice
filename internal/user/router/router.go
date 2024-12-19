package router

import (
	"github.com/Ajulll22/belajar-microservice/internal/user/config"
	"github.com/Ajulll22/belajar-microservice/pkg/cache"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(router *gin.Engine, db *gorm.DB, redis cache.Cache, cfg config.Config) {}
