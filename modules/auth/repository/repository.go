package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"vault/modules/auth/models"
)

type User interface {
	GetUserCreds(login string) (models.UserCredentials, error)
	GetUserTotpKey(id uuid.UUID) (string, error)
	GetAppRoleCreds(roleId string) (models.AppRoleCredentials, error)
	// TODO: CreateUserPass
	// TODO: CreateAppRole
}

type Role interface {
	// TODO: CreateRole
	// TODO: GetRole
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
	User
	Policy
	Role
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserRepo(db),
		//User: NewUserRepo(db),
	}
}
