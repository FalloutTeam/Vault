package repository

import (
	"github.com/jmoiron/sqlx"
	"vault/modules/totp/models"
)

type TotpRepo struct {
	db *sqlx.DB
}

func (t TotpRepo) CreateTotp(name string, totp models.Totp) error {
	return nil
}

func NewTotpRepo(db *sqlx.DB) *TotpRepo {
	return &TotpRepo{db: db}
}
