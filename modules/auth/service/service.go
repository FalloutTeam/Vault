package service

import (
	"vault/modules/auth/models"
	"vault/modules/auth/repository"
)

type Login interface {
	UserPassLogin(userLogin models.UserPassLogin) (string, error)
	AppRoleLogin(appRoleLogin models.AppRoleLogin) (string, error)
	// TODO: CreateUserPass
	// TODO: CreateAppRole
}

type Role interface {
	// TODO: CreateRole
	// TODO: GetRole
}

type User interface {
}

type Service struct {
	Login
	User
	Role
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Login: NewAuthService(repo),
		User:  NewUserService(repo),
	}
}
