package services

import (
	"beta/internal/repositories"
	"beta/internal/request"
	"beta/pkg/config"
	"beta/pkg/regulation"
	"context"
)

type VotingService interface {
	CreateVoting(ctx *context.Context, input *request.Vote)
}

type GammaStateService interface {
	ClearInfoForRequests(requestRegulator *regulation.Request)
}

type IService interface {
	VotingService
	GammaStateService
}

type Service struct {
	repositories      repositories.VotingRepository
	config            *config.Config
	requestRegulation *regulation.Request
}

func NewService(repos repositories.VotingRepository, conf *config.Config, requestRegulation *regulation.Request) *Service {
	return &Service{repos, conf, requestRegulation}
}
