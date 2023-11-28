package repositories

import (
	"beta/internal/request"
	"beta/pkg/database"
	"go.mongodb.org/mongo-driver/mongo"
)

type Voting interface {
	InsertVote(input request.Vote) (string, error)
}

type Repository struct {
	Voting
	database *mongo.Database
}

func NewRepository(client *database.DatabaseClient) *Repository {
	return &Repository{database: client.GetClient().Database("beta")}
}
