package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"vault/modules/auth/models"
)

type AuthRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (s *AuthRepo) GetUserTotpKey(id uuid.UUID) (string, error) {
	return "B2KKFE7NTMVCNXSPKBQD47BEEU2IESQ3", nil
}

func (s *AuthRepo) GetUserCreds(login string) (models.UserCredentials, error) {
	return models.UserCredentials{
		Login:        login,
		PasswordHash: "$2a$10$iQtJMTXWYK88g3uFt./jm.UPHlNd37BJLUmcAE/J.34KAM1jcz3ba",
		MfaEnabled:   true,
	}, nil
}

func (s *AuthRepo) GetAppRoleCreds(roleId string) (models.AppRoleCredentials, error) {
	return models.AppRoleCredentials{
		RoleId:     roleId,
		SecretHash: "$2a$10$iQtJMTXWYK88g3uFt./jm.UPHlNd37BJLUmcAE/J.34KAM1jcz3ba",
	}, nil
}
