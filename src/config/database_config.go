package config

import (
	"context"
	"log"

	"github.com/kaungmyathan22/golang-product-crud-with-auth/src/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDatabase() *mongo.Client {
	uri := AppConfigInstance.DATABASE_URI
	ctx := context.TODO()
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	logger.Info("Pinged. You successfully connected to MongoDB!")
	logger.Info("Connected to database.")

	return client
}
