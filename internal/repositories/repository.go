package repositories

import (
	"beta/internal/entity"
	"beta/pkg/config"
	"beta/pkg/database"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type IVotingRepository interface {
	InsertVoteInVoting(ctx *context.Context, vote entity.Vote, votingId string) (string, error)
	FindVoting(ctx *context.Context, votingId string) (*entity.Voting, error)
	InsertNewVoting(ctx *context.Context, vote *entity.Voting) (string, error)
	GetVotingCountVotes(ctx *context.Context, votingId string) ([]entity.VotesStateCount, error)
	GetVotesStates(ctx *context.Context, votingId string) ([]entity.VotesStatePercents, error)
}

type Repository struct {
	database *mongo.Database
}

type VotingRepository struct {
	*Repository
	collection *mongo.Collection
}

func NewRepository(client *database.DatabaseClient, config config.DatabaseConfig) *Repository {
	return &Repository{database: client.GetClient().Database(config.Name)}
}

func NewVotingRepository(repository *Repository) *VotingRepository {
	return &VotingRepository{repository, repository.database.Collection("vote")}
}
