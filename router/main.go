package router

import (
	"github.com/DeathVenom54/github-deploy-inator/handlers"
	"github.com/DeathVenom54/github-deploy-inator/structs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func CreateRouter(config *structs.Config) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	router.Post(config.Endpoint, handlers.CreateWebhookHandler(config))

	return router
}
