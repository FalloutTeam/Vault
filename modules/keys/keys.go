package keys

import (
	_handler "vault/modules/keys/handler"
	"vault/modules/keys/repository"
	_service "vault/modules/keys/service"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

func InitModule(router *chi.Mux, db *sqlx.DB) {
	repo := repository.NewRepository(db)

	service := _service.NewService(repo)

	handler := _handler.NewHandler(service)

	router.Route("/secret", func(secret chi.Router) {
		secret.Post("/{path}", handler.AddSecret)
		secret.Delete("/{path}", handler.DeleteSecret)
		secret.Patch("/{path}", handler.PatchSecret)
		secret.Get("/{path}", handler.GetSecret)
	})
}
