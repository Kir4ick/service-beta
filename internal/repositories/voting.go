package repositories

import (
	"beta/internal/entity"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (repository *Repository) InsertVoteInVoting(ctx *context.Context, vote entity.Vote, votingId string) (string, error) {
	filter := bson.D{{"votingid", votingId}}
	push := bson.D{{"$push", bson.M{"votes": vote}}}
	_, err := repository.database.Collection("vote").UpdateOne(*ctx, filter, push)
	return votingId, err
}

func (repository *Repository) FindVoting(ctx *context.Context, votingId string) (*entity.Voting, error) {
	filter := bson.D{{"votingid", votingId}}
	var ent entity.Voting
	err := repository.database.Collection("vote").FindOne(*ctx, filter).Decode(&ent)

	if err == mongo.ErrNoDocuments {
		return nil, err
	}

	return &ent, err
}

func (repository *Repository) InsertNewVoting(ctx *context.Context, vote *entity.Voting) (string, error) {
	_, err := repository.database.Collection("vote").InsertOne(*ctx, vote)
	return vote.VotingId, err
}
