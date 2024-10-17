package handler

import (
	"encoding/json"
	"net/http"
	"vault/modules/auth/handler/response"
	"vault/modules/auth/models"
	"vault/modules/auth/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) AppRoleLogin(w http.ResponseWriter, r *http.Request) {
	var body struct {
		RoleId    string `json:"role_id" binding:"required"`
		SecreteId string `json:"secrete_id" binding:"required"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	token, err := h.service.Login.AppRoleLogin(models.AppRoleLogin{
		RoleId:   body.RoleId,
		SecretId: body.SecreteId,
	})

	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	auth := response.AppRoleAuth{ClientToken: token}

	json.NewEncoder(w).Encode(response.AppRoleLoginResponse{AppRoleAuth: auth})

}

func (h *Handler) UserPassLogin(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"login" binding:"required"`
		Password string `json:"password" binding:"required"`
		Totp     string `json:"totp"`
	}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	token, err := h.service.Login.UserPassLogin(models.UserPassLogin{
		Login:    creds.Username,
		Password: creds.Password,
		Totp:     creds.Totp,
	})
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	auth := response.UserpassAuth{ClientToken: token}

	json.NewEncoder(w).Encode(response.UserpassLoginResponse{UserpassAuth: auth})

}

//func RegisterRoutes(router *chi.Mux, service service.AuthService) {
//	//handler := NewAuthHandler(service)
//	router.Route("/auth", func(r chi.Router) {
//		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {})
//	})
//}
