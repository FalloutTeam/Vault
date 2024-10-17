package service

import (
	"vault/modules/totp/repository"
)

type TotpService struct {
	repo *repository.Repository
}

func (t TotpService) CreateTotp(userId string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func NewTotpService(repo *repository.Repository) *TotpService {
	return &TotpService{repo: repo}
}
