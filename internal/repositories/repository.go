package repositories

import (
	"beta/internal/entity"
	"beta/pkg/database"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type VotingRepository interface {
	InsertVoteInVoting(ctx *context.Context, vote entity.Vote, votingId string) (string, error)
	FindVoting(ctx *context.Context, votingId string) (*entity.Voting, error)
	InsertNewVoting(ctx *context.Context, vote *entity.Voting) (string, error)
	GetVotingCountVotes(ctx *context.Context, votingId string) ([]entity.VotesStateCount, error)
	GetVotesStates(ctx *context.Context, votingId string) ([]entity.VotesStatePercents, error)
}

type Repository struct {
	database *mongo.Database
}

func NewRepository(client *database.DatabaseClient) *Repository {
	return &Repository{database: client.GetClient().Database("beta")}
}
