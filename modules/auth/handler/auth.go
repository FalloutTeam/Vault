package handler

import (
	"encoding/json"
	"net/http"
	"vault/modules/auth/handler/response"
	"vault/modules/auth/models"
	"vault/modules/auth/service"
)

type AuthHandler struct {
	service *service.Service
}

func NewAuthHandler(service *service.Service) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"login" binding:"required"`
		Password string `json:"password" binding:"required"`
		Totp     string `json:"totp"`
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	token, err := h.service.Auth.Login(models.UserLogin{
		Login:    creds.Username,
		Password: creds.Password,
		Totp:     creds.Totp,
	})
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	auth := response.Auth{ClientToken: token}

	json.NewEncoder(w).Encode(response.LoginResponse{Auth: auth})

}

//func RegisterRoutes(router *chi.Mux, service service.AuthService) {
//	//handler := NewAuthHandler(service)
//	router.Route("/auth", func(r chi.Router) {
//		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {})
//	})
//}
