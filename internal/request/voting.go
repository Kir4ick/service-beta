package request

type Vote struct {
	VoteID   string `json:"voteId" binding:"required"`
	VotingID string `json:"votingId" binding:"required"`
	OptionID string `json:"optionId" binding:"required"`
}
