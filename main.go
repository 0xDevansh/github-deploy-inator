package main

import (
	"github.com/DeathVenom54/github-deploy-inator/logger"
	"github.com/DeathVenom54/github-deploy-inator/router"
	"net/http"
)

func main() {
	r := router.CreateRouter()

	err := http.ListenAndServe(":4567", r)
	if err != nil {
		logger.Error(err)
	}

	logger.Log("Listening...")
}
