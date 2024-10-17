package service

import (
	"vault/modules/totp/repository"
)

type Totp interface {
	CreateTotp(userId string) (string, error)
}

type User interface {
}

type Service struct {
	Totp
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Totp: NewTotpService(repo),
	}
}
