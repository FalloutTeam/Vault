package handler

import (
	"net/http"
	"vault/modules/totp/service"
)

type TotpHandler struct {
	service *service.Service
}

func (t TotpHandler) CreateTotp(w http.ResponseWriter, r *http.Request) {
	//TODO implement me
	panic("implement me")
}

func NewTotpHandler(service *service.Service) *TotpHandler {
	return &TotpHandler{
		service: service,
	}
}
