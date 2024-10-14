package handler

import (
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

//func RegisterRoutes(router *chi.Mux, service service.AuthService) {
//	//handler := NewAuthHandler(service)
//	router.Route("/auth", func(r chi.Router) {
//		r.Post("/login", func(w http.ResponseWriter, r *http.Request) {})
//	})
//}
