package services

import (
	gammaClient "beta/internal/client"
	"beta/internal/entity"
	"beta/internal/request"
	"log"
)

func (s *Service) CreateVoting(input *request.Vote) {
	voting, err := s.repositories.FindVoting(s.ctx, input.VotingID)
	var votingId string

	if voting == nil {
		votingId, err = s.repositories.InsertNewVoting(s.ctx, s.formEntityToInsertNewVoting(input))
	} else {
		votingId, err = s.repositories.InsertVoteInVoting(s.ctx, entity.Vote{Id: input.VoteID, OptionId: input.OptionID}, voting.VotingId)
	}

	if err != nil {
		log.Printf("insert new vote err %s", err.Error())
	}

	if s.checkNeedToRequestGamma(votingId) {
		client := gammaClient.NewGammaClient(s.config.GammaUrl)
		err = client.Send(s.formRequestDataForSend(votingId))

		if err != nil {
			log.Printf("error to send gamma service: %s", err.Error())
		}
	}
}

func (s *Service) checkNeedToRequestGamma(votingId string) bool {
	return true
}

func (s *Service) formRequestDataForSend(voteId string) *request.VoteGammaRequest {
	votingCount := request.VotingResult{Count: 20, OptionId: "s;aflmsaflm"}
	arrayResult := []request.VotingResult{votingCount}
	return request.NewGammaRequest("sdsadsdasda", arrayResult)
}

func (s *Service) formEntityToInsertNewVoting(input *request.Vote) *entity.Voting {
	voteArray := []entity.Vote{entity.Vote{Id: input.VoteID, OptionId: input.OptionID}}
	voting := entity.Voting{VotingId: input.VotingID, Votes: voteArray}
	return &voting
}
