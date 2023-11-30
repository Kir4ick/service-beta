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

func (repository *Repository) GetVotingCountVotes(ctx *context.Context, votingId string) ([]entity.Vote, error) {

	function := bson.A{
		bson.D{{"$match", bson.D{{"votingid", votingId}}}},
		bson.D{{"$unwind", "$votes"}},
		bson.D{
			{"$group",
				bson.D{
					{"_id", "$votes.optionid"},
					{"count", bson.D{{"$count", bson.D{}}}},
				},
			},
		},
	}

	result, err := repository.database.Collection("vote").Aggregate(*ctx, function)
	if err != nil {
		return nil, err
	}

	var resultArray []entity.Vote
	err = result.All(*ctx, &resultArray)

	return resultArray, err
}

func (repository *Repository) GetVotesCondition(ctx *context.Context, votingId string, optionId string) {

	function := bson.A{
		bson.D{{"$match", bson.D{{"votingid", votingId}}}},
		bson.D{
			{"$project",
				bson.D{
					{"_id", 0},
					{"votesCount", bson.D{{"$size", "$votes"}}},
					{"votes", "$votes"},
					{"votingid", "$votingid"},
				},
			},
		},
		bson.D{{"$unwind", "$votes"}},
		bson.D{
			{"$group",
				bson.D{
					{"_id", "$votes.optionid"},
					{"count", bson.D{{"$count", bson.D{}}}},
					{"votesCount", bson.D{{"$push", "$votesCount"}}},
				},
			},
		},
		bson.D{
			{"$project",
				bson.D{
					{"_id", "$_id"},
					{"count", "$count"},
					{"votesCount", bson.D{{"$first", "$votesCount"}}},
				},
			},
		},
		bson.D{
			{"$project",
				bson.D{
					{"_id", "$_id"},
					{"percents",
						bson.D{
							{"$multiply",
								bson.A{
									bson.D{
										{"$divide",
											bson.A{
												"$count",
												"$votesCount",
											},
										},
									},
									100,
								},
							},
						},
					},
				},
			},
		},
		bson.D{{"$match", bson.D{{"_id", optionId}}}},
	}

	result, err := repository.database.Collection("vote").Aggregate(*ctx, function)
}
