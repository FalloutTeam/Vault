package handler

import (
	"net/http"
	"vault/modules/totp/service"
)

type Totp interface {
	CreateTotp(w http.ResponseWriter, r *http.Request)
	//TODO: Delete()
	//TODO: Read()
	//TODO: List()
}
type Handler struct {
	Totp
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		Totp: NewTotpHandler(s),
	}
}
