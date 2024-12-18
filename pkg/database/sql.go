package database

import (
	"fmt"

	"github.com/Ajulll22/belajar-microservice/pkg/security"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SQLConnect(DB_USER string, DB_PASS string, DB_HOST string, DB_PORT string, DB_NAME string) *gorm.DB {
	db_username := DB_USER
	db_password := DB_PASS
	db_server := DB_HOST
	db_port := DB_PORT
	db_name := DB_NAME

	clear_password := security.Decrypt(db_password, "62277ecdae08d9e813ab17a4ec2db8c58db38e398617824a2ef035c64d3da4be")
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", db_username, clear_password, db_server, db_port, db_name)

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Error),
		PrepareStmt: true,
	})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}
