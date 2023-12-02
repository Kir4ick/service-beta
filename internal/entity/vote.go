package entity

type Vote struct {
	VoteId   string `json:"id" bson:"vote_id"`
	OptionId string `json:"optionId" bson:"option_id"`
}

type Voting struct {
	VotingId string `json:"votingId" bson:"voting_id"`
	Votes    []Vote `json:"votes" bson:"votes"`
}

type VotesStatePercents struct {
	OptionId string  `json:"optionId" bson:"_id"`
	Percents float64 `bson:"percents"`
}

type VotesStateCount struct {
	OptionId string `json:"optionId" bson:"_id"`
	Count    int    `json:"count" bson:"count"`
}
