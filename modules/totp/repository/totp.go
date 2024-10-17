package repository

import (
	"github.com/jmoiron/sqlx"
	"vault/modules/totp/models"
)

type TotpRepo struct {
	db *sqlx.DB
}

func (t TotpRepo) CreateTotp(name string, totp models.Totp) error {
	// TODO:  сохранение totp-ключа, для доступа - название (name).
	// TODO: Eсли ключ уже существует - переписать.
	return nil
}

func NewTotpRepo(db *sqlx.DB) *TotpRepo {
	return &TotpRepo{db: db}
}
