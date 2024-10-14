package auth

import (
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"net/http"
	_handler "vault/modules/auth/handler"
	"vault/modules/auth/repository"
	_service "vault/modules/auth/service"
)

func InitModule(router *chi.Mux, db *sqlx.DB) {
	repo := repository.NewRepository(db)

	service := _service.NewService(repo)

	handler := _handler.NewHandler(service)

	// Регистрация ручек
	router.Route("/auth", func(auth chi.Router) {
		auth.Route("/login", func(login chi.Router) {
			login.Post("/userpass", handler.Auth.Login)
			login.Post("/app_role", func(w http.ResponseWriter, r *http.Request) {})
		})
		// Другие маршруты аутентификации
	})
}
