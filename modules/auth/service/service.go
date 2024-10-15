package service

import (
	"vault/modules/auth/models"
	"vault/modules/auth/repository"
)

type Auth interface {
	UserPassLogin(userLogin models.UserPassLogin) (string, error)
	AppRoleLogin(appRoleLogin models.AppRoleLogin) (string, error)
}

type User interface {
}

type Service struct {
	Auth
	User
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repo),
		User: NewUserService(repo),
	}
}
