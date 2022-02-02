package router

import (
	"fmt"
	"github.com/DeathVenom54/github-deploy-inator/handlers"
	"github.com/DeathVenom54/github-deploy-inator/logger"
	"github.com/DeathVenom54/github-deploy-inator/structs"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func CreateRouter(config *structs.Config) *chi.Mux {
	router := chi.NewRouter()

	router.Use(recovery)

	router.Post(config.Endpoint, handlers.CreateWebhookHandler(config))

	return router
}

func recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				logger.Error(fmt.Errorf("error: %v", err), false)

				w.Header().Set("Content-Type", "text/plain")
				w.WriteHeader(http.StatusInternalServerError)
				_, err := w.Write([]byte("500 Internal server error"))
				if err != nil {
					logger.Error(err, false)
				}
			}

		}()

		next.ServeHTTP(w, r)

	})
}
