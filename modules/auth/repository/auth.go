package repository

import "github.com/jmoiron/sqlx"

type AuthRepo struct {
	db *sqlx.DB
}

func (a *AuthRepo) UserPassLogin(login string, passwordHash string) error {
	return nil
}

func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{db: db}
}
