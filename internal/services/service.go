package services

import "beta/internal/repositories"

type Voting interface {
}

type Service struct {
	repositories *repositories.Repository
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{repos}
}
