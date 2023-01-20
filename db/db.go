package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Init() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connectionString := os.Getenv("MONGODB_CONNSTRING")

	if connectionString == "" {
		log.Fatal("MONGODB_CONNSTRING is empty")
	}

	options := options.Client().ApplyURI(connectionString)

	return mongo.Connect(ctx, options)
}
