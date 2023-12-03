package services

import (
	gammaClient "beta/internal/client"
	"beta/internal/entity"
	"beta/internal/request"
	"context"
	"log"
	"time"
)

func (s *Service) CreateVoting(ctx *context.Context, input *request.Vote) {
	voting, err := s.repositories.FindVoting(ctx, input.VotingID)

	var votingId string

	afterInsertVotingState, err := s.repositories.GetVotesStates(ctx, input.VotingID)

	if voting == nil {
		votingId, err = s.repositories.InsertNewVoting(ctx, s.formEntityToInsertNewVoting(input))
	} else {
		votingId, err = s.repositories.InsertVoteInVoting(
			ctx,
			entity.Vote{VoteId: input.VoteID, OptionId: input.OptionID},
			voting.VotingId)
	}

	if err != nil {
		log.Printf("insert new vote err %s", err.Error())
		return
	}

	beforeInsertVotingState, err := s.repositories.GetVotesStates(ctx, input.VotingID)

	if !s.checkNeedToSendGamma(afterInsertVotingState, beforeInsertVotingState) {
		return
	}

	s.sendToGammaService(ctx, votingId)
}

func (s *Service) checkNeedToSendGamma(
	afterInsertVotingStates []entity.VotesStatePercents,
	beforeInsertVotingStates []entity.VotesStatePercents) bool {

	//Первым делом проверяем сколько в эту секунду было уже послано запросов, если 2 или больше, то ниче не отправляем
	now := time.Now()
	count := s.requestRegulation.GetCountRequestsNow(now)

	if count >= 2 {
		s.requestRegulation.ClearInfo(now)
		return false
	}

	//Если укладываемся в лимит (2 запроса в секунду), то уже проверяем состояние
	for _, beforeInsertVotingState := range beforeInsertVotingStates {
		for _, afterInsertVotingState := range afterInsertVotingStates {

			if beforeInsertVotingState.OptionId != afterInsertVotingState.OptionId {
				continue
			}

			//Приводим к типу int, чтобы округлилось в меньшую сторону
			checkStateDifference := int(afterInsertVotingState.Percents - beforeInsertVotingState.Percents)

			if checkStateDifference >= 1 {
				return true
			}
		}
	}

	return false
}

func (s *Service) formRequestDataForSend(ctx *context.Context, votingId string) (*request.VoteGammaRequest, error) {
	result, err := s.repositories.GetVotingCountVotes(ctx, votingId)

	if err != nil {
		return nil, err
	}

	arrayResult := []request.VotingResult{}

	for res, voteEntity := range result {
		log.Print(res, voteEntity)
		voteResult := request.VotingResult{Count: voteEntity.Count, OptionId: voteEntity.OptionId}
		arrayResult = append(arrayResult, voteResult)
	}

	return request.NewGammaRequest(votingId, arrayResult), nil
}

func (s *Service) formEntityToInsertNewVoting(input *request.Vote) *entity.Voting {
	voteArray := []entity.Vote{entity.Vote{VoteId: input.VoteID, OptionId: input.OptionID}}
	voting := entity.Voting{VotingId: input.VotingID, Votes: voteArray}
	return &voting
}

func (s *Service) sendToGammaService(ctx *context.Context, votingId string) {
	requestData, err := s.formRequestDataForSend(ctx, votingId)

	if err != nil {
		log.Printf("error to form gamma request: %s", err.Error())
		return
	}

	s.requestRegulation.NewRequest(time.Now())
	client := gammaClient.NewGammaClient(s.config.GammaUrl)
	err = client.Send(requestData)

	if err != nil {
		log.Printf("error to send gamma service: %s", err.Error())
	}
}
