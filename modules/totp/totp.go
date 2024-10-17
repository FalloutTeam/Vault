package totp

import (
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	_handler "vault/modules/totp/handler"
	"vault/modules/totp/repository"
	_service "vault/modules/totp/service"
)

func InitModule(router *chi.Mux, db *sqlx.DB) {
	repo := repository.NewRepository(db)

	service := _service.NewService(repo)

	handler := _handler.NewHandler(service)

	router.Route("/totp", func(totp chi.Router) {
		totp.Route("/keys", func(keys chi.Router) {
			keys.Post("/{name}", handler.CreateTotp)
		})
	})
}
