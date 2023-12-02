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

type Service struct {
	repositories      *repositories.Repository
	config            *config.Config
	requestRegulation *regulation.Request
}

func NewService(repos *repositories.Repository, conf *config.Config, requestRegulation *regulation.Request) *Service {
	return &Service{repos, conf, requestRegulation}
}
