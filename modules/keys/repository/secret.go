package repository

import "github.com/jmoiron/sqlx"

type SecretRepo struct {
	db *sqlx.DB
}

func NewSecretRepo(db *sqlx.DB) *SecretRepo {
	return &SecretRepo{db: db}
}
