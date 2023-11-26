package repositories

import "beta/pkg/database"

type Voting interface {
}

type Repository struct {
	Voting
	client *database.DatabaseClient
}

func NewRepository() *Repository {
	return &Repository{}
}
