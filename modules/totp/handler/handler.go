package handler

import (
	"net/http"
	"vault/modules/totp/service"
)

type Totp interface {
	CreateTotp(w http.ResponseWriter, r *http.Request)
	//DeleteTotp(w http.ResponseWriter, r *http.Request)
}
type Handler struct {
	Totp
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		Totp: NewTotpHandler(s),
	}
}
