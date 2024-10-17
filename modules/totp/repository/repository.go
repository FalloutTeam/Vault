package repository

import (
	"github.com/jmoiron/sqlx"
	"vault/modules/totp/models"
)

type Totp interface {
	CreateTotp(name string, totp models.Totp) error
	//TODO: Delete()
	//TODO: Read()
	//TODO: List()
}
type Repository struct {
	Totp
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Totp: NewTotpRepo(db),
	}
}
