package database

import (
	"context"
	"fmt"

	"github.com/Ajulll22/belajar-microservice/pkg/security"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoConnect(ctx context.Context, DB_USER string, DB_PASS string, DB_HOST string, DB_PORT string, DB_NAME string) *mongo.Database {
	db_username := DB_USER
	db_password := DB_PASS
	db_server := DB_HOST
	db_port := DB_PORT
	db_name := DB_NAME

	clear_password := security.Decrypt(db_password, "62277ecdae08d9e813ab17a4ec2db8c58db38e398617824a2ef035c64d3da4be")
	dsn := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", db_username, clear_password, db_server, db_port, db_name)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	if err != nil {
		panic("failed to connect database")
	}

	if err := client.Ping(context.TODO(), nil); err != nil {
		panic("MongoDB client is disconnected")
	}

	return client.Database(db_name)
}
