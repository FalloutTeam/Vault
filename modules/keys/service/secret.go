package service

import (
	"vault/modules/keys/repository"
)

type SecretService struct {
	repo *repository.Repository
}

func NewSecretService(repo *repository.Repository) *SecretService {
	return &SecretService{repo: repo}
}
