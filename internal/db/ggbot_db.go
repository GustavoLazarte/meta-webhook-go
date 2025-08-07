package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db *mongo.Database
)

func ConnectDatabase(host string) (*mongo.Database, error) {
	dbName := os.Getenv("MONGO_HOST_DATABASE_NAME")
	clientOptions := options.Client().ApplyURI(host)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	err = client.Database(dbName).Drop(context.TODO())
	if err != nil {
		return nil, err
	}
	db = client.Database(dbName)
	return db, nil
}

func GetDB() *mongo.Database {
	return db
}
