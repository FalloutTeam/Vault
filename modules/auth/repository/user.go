package repository

import (
	"github.com/jmoiron/sqlx"
	"vault/modules/auth/service/models"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (s *UserRepo) GetUserCreds(login string) (models.UserCredentials, error) {
	return models.UserCredentials{
		Login:        login,
		PasswordHash: "$2a$10$iQtJMTXWYK88g3uFt./jm.UPHlNd37BJLUmcAE/J.34KAM1jcz3ba",
	}, nil
}
