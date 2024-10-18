package repository

import "github.com/jmoiron/sqlx"

type Secret interface {
}

type Repository struct {
	Secret
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Secret: NewSecretRepo(db),
	}
}
