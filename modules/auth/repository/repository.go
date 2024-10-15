package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"vault/modules/auth/models"
)

type Auth interface {
	GetUserCreds(login string) (models.UserCredentials, error)
	GetUserTotpKey(id uuid.UUID) (string, error)
	GetAppRoleCreds(roleId string) (models.AppRoleCredentials, error)
}

//type User interface {
//	GetUserCreds(login string) (models.UserCredentials, error)
//	GetUserTotpKey(id uuid.UUID) (string, error)
//}
//
//type AppRole interface {
//	GetAppRoleCreds(roleId string) (string, error)
//}

type Policy interface{}

type Repository struct {
	Auth
	Policy
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthRepo(db),
		//User: NewUserRepo(db),
	}
}
