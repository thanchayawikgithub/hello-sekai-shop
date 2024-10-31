package database

import (
	"context"
	"log"
	"time"

	"github.com/thanchayawikgithub/hello-sekai-shop/config"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func Conn(config *config.DB) *mongo.Client {
	client, err := mongo.Connect(options.Client().ApplyURI(config.URI))
	if err != nil {
		log.Fatal("Error: Failed to connect database: ", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("Error: Failed to ping database: ", err)
	}

	return client
}
