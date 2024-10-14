package repository

import (
	"github.com/jmoiron/sqlx"
	"vault/modules/auth/service/models"
)

type Auth interface {
	UserPassLogin(login string, passwordHash string) error
}

type User interface {
	GetUserCreds(login string) (models.UserCredentials, error)
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
