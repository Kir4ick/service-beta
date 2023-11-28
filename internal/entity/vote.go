package entity

type Vote struct {
	Id       string `json:"id"`
	OptionId string `json:"optionId"`
}

type Voting struct {
	VotingId string `json:"votingId"`
	Votes    []Vote `json:"votes"`
}
