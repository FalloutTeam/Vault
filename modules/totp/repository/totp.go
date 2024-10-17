package repository

import "github.com/jmoiron/sqlx"

type TotpRepo struct {
	db *sqlx.DB
}

func (t TotpRepo) CreateTotp(userId string) error {
	//TODO implement me
	panic("implement me")
}

func NewTotpRepo(db *sqlx.DB) *TotpRepo {
	return &TotpRepo{db: db}
}
