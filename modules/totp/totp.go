package totp

import (
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

func InitModule(router *chi.Mux, db *sqlx.DB) {
	//repo := repository.NewRepository(db)
	//
	//service := _service.NewService(repo)
	//
	//handler := _handler.NewHandler(service)

	router.Route("/totp", func(totp chi.Router) {

	})
}
