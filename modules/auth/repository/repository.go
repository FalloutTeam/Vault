package repository

import "github.com/jmoiron/sqlx"

type Auth interface {
}

type User interface {
}

type Repository struct {
	Auth
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthRepo(db),
		User: NewUserRepo(db),
	}
}
