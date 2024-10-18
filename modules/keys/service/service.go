package service

import "vault/modules/keys/repository"

type Secret interface {
}

type Service struct {
	Secret
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Secret: NewSecretService(repo),
	}
}
