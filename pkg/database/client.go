package database

import (
	"beta/pkg/config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type DatabaseClient struct {
	client *mongo.Client
}

func Connection(config *config.DatabaseConfig, ctx *context.Context) *DatabaseClient {
	client, err := mongo.NewClient(options.Client().ApplyURI(config.Host + ":" + config.Port))

	if err != nil {
		log.Fatalf("cannot connect to database %s", err.Error())
	}

	databaseClient := DatabaseClient{client: client}

	err = client.Connect(*ctx)
	if err != nil {
		log.Fatalf("cannot connect to database %s", err.Error())
	}

	return &databaseClient
}

func (database *DatabaseClient) Disconnect(ctx *context.Context) error {
	return database.client.Disconnect(*ctx)
}

func (database *DatabaseClient) GetClient() *mongo.Client {
	return database.client
}
