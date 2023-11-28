package services

import (
	"beta/internal/repositories"
	"beta/pkg/config"
	"context"
)

type Voting interface {
}

type Service struct {
	repositories *repositories.Repository
	config       *config.Config
	ctx          *context.Context
}

func NewService(repos *repositories.Repository, conf *config.Config, ctx *context.Context) *Service {
	return &Service{repos, conf, ctx}
}
