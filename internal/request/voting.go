package request

type Vote struct {
	VoteID   string `json:"voteId" binding:"required,uuid"`
	VotingID string `json:"votingId" binding:"required,uuid"`
	OptionID string `json:"optionId" binding:"required,uuid"`
}

type VotingResult struct {
	OptionId string `json:"optionId"`
	Count    int    `json:"count"`
}

type VoteGammaRequest struct {
	VotingId string         `json:"votingId"`
	Results  []VotingResult `json:"results"`
}

func NewGammaRequest(votingId string, result []VotingResult) *VoteGammaRequest {
	return &VoteGammaRequest{VotingId: votingId, Results: result}
}
