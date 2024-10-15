package auth

import (
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_handler "vault/modules/auth/handler"
	"vault/modules/auth/repository"
	_service "vault/modules/auth/service"
)

func InitModule(router *chi.Mux, db *sqlx.DB) {
	repo := repository.NewRepository(db)

	service := _service.NewService(repo)

	handler := _handler.NewHandler(service)

	router.Route("/auth", func(auth chi.Router) {
		auth.Route("/appRole", func(appRole chi.Router) {
			appRole.Post("/login", handler.AppRoleLogin)
		})
		auth.Route("/userpass", func(userpass chi.Router) {
			userpass.Post("login", handler.UserPassLogin)
		})
		// TODO: Другие маршруты аутентификации
	})
}
