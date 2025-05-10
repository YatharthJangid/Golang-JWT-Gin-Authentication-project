package database

import (
	"context"
	"errors"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	clientErr  error
	clientOnce sync.Once
)

func Connect() (*mongo.Client, error) {
	clientOnce.Do(func() {
		_ = godotenv.Load(".env")

		mongoURI := os.Getenv("MONGODB_URL")
		if mongoURI == "" {
			clientErr = errors.New("MONGODB_URL is required")
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, clientErr = mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	})

	return client, clientErr
}

func OpenCollection(collectionName string) (*mongo.Collection, error) {
	client, err := Connect()
	if err != nil {
		return nil, err
	}

	return client.Database(databaseName()).Collection(collectionName), nil
}

func databaseName() string {
	if name := os.Getenv("DB_NAME"); name != "" {
		return name
	}

	return "go_jwt_demo"
}
