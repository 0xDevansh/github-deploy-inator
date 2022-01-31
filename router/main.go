package router

import (
	"github.com/DeathVenom54/github-deploy-inator/handlers"
	"github.com/DeathVenom54/github-deploy-inator/structs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func CreateRouter(listeners []structs.Listener) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	for _, listener := range listeners {
		router.Post(listener.Endpoint, handlers.CreateWebhookHandler(listener))
	}

	return router
}
