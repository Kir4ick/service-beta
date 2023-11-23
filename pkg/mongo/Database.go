package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Connection struct {
	client *mongo.Client
}

func (database *Connection) NewDatabase(host string, port string) *Connection {
	client, err := mongo.NewClient(options.Client().ApplyURI(host + port))

	if err != nil {
		log.Fatalf("cannot connect to mongo %s", err.Error())
	}

	database.client = client
	return database
}

func (database *Connection) Connect(ctx *context.Context) error {
	return database.client.Connect(*ctx)
}
