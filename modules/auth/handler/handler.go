package handler

import (
	"net/http"
	"vault/modules/auth/service"
)

type Auth interface {
	Login(w http.ResponseWriter, r *http.Request)
}
type Handler struct {
	Auth
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		Auth: NewAuthHandler(s),
	}
}
