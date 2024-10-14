package handler

import "vault/modules/auth/service"

type Auth interface {
}
type Handler struct {
	Auth
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		Auth: NewAuthHandler(s),
	}
}
