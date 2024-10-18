package handler

import (
	"net/http"
	"vault/modules/keys/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) AddSecret(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) DeleteSecret(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) PatchSecret(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) GetSecret(w http.ResponseWriter, r *http.Request) {

}
