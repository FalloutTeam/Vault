package service

import (
	"vault/modules/auth/repository"
)

type Auth interface {
	Login(login string, password string) (string, error)
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
