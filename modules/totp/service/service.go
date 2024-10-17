package service

import (
	"vault/modules/totp/models"
	"vault/modules/totp/repository"
)

type Totp interface {
	CreateTotp(name string, totpParams models.Totp) (string, string, error)
	//TODO: Delete()
	//TODO: Read()
	//TODO: List()
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
