package repositories

import (
	"beta/internal/entity"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (rep *Repository) InsertVoteInVoting(ctx *context.Context, vote entity.Vote, votingId string) (string, error) {
	filter := bson.D{{"voting_id", votingId}}
	push := bson.D{{"$push", bson.M{"votes": vote}}}
	_, err := rep.database.Collection("vote").UpdateOne(*ctx, filter, push)
	return votingId, err
}

func (rep *Repository) FindVoting(ctx *context.Context, votingId string) (*entity.Voting, error) {
	filter := bson.D{{"voting_id", votingId}}
	var ent entity.Voting
	err := rep.database.Collection("vote").FindOne(*ctx, filter).Decode(&ent)

	if err == mongo.ErrNoDocuments {
		return nil, err
	}

	return &ent, err
}

func (rep *Repository) InsertNewVoting(ctx *context.Context, vote *entity.Voting) (string, error) {
	_, err := rep.database.Collection("vote").InsertOne(*ctx, vote)
	return vote.VotingId, err
}

func (rep *Repository) GetVotingCountVotes(ctx *context.Context, votingId string) ([]entity.VotesStateCount, error) {

	function := bson.A{
		bson.D{{"$match", bson.D{{"voting_id", votingId}}}},
		bson.D{{"$unwind", "$votes"}},
		bson.D{
			{"$group",
				bson.D{
					{"_id", "$votes.option_id"},
					{"count", bson.D{{"$count", bson.D{}}}},
				},
			},
		},
	}

	result, err := rep.database.Collection("vote").Aggregate(*ctx, function)
	if err != nil {
		return nil, err
	}

	var resultArray []entity.VotesStateCount
	err = result.All(*ctx, &resultArray)

	return resultArray, err
}

func (rep *Repository) GetVotesStates(ctx *context.Context, votingId string) ([]entity.VotesStatePercents, error) {

	aggregate := bson.A{
		bson.D{{"$match", bson.D{{"voting_id", votingId}}}},
		bson.D{
			{"$project",
				bson.D{
					{"_id", 0},
					{"votesCount", bson.D{{"$size", "$votes"}}},
					{"votes", "$votes"},
					{"voting_id", "$voting_id"},
				},
			},
		},
		bson.D{{"$unwind", "$votes"}},
		bson.D{
			{"$group",
				bson.D{
					{"_id", "$votes.option_id"},
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
	}

	result, err := rep.database.Collection("vote").Aggregate(*ctx, aggregate)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	var resultArray []entity.VotesStatePercents
	err = result.All(*ctx, &resultArray)

	return resultArray, err
}
