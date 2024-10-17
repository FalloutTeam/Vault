package repository

import "github.com/jmoiron/sqlx"

type Totp interface {
	CreateTotp(userId string) error
}
type Repository struct {
	Totp
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Totp: NewTotpRepo(db),
	}
}
