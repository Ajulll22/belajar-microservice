package config

import (
	"fmt"

	"github.com/Ajulll22/belajar-microservice/pkg/security"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DbConnect(cfg Config) *gorm.DB {
	db_username := cfg.DB_USER
	db_password := cfg.DB_PASS
	db_server := cfg.DB_HOST
	db_port := cfg.DB_PORT
	db_name := cfg.DB_NAME

	clear_password := security.Decrypt(db_password, "62277ecdae08d9e813ab17a4ec2db8c58db38e398617824a2ef035c64d3da4be")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_username, clear_password, db_server, db_port, db_name)
	// dsn = fmt.Sprintf("sqlserver://%s:123456@localhost:%s?database=%s", db_username, db_port, db_name)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Error),
		PrepareStmt: true,
	})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
