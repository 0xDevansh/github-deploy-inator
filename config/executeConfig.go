package config

import (
	"fmt"
	"github.com/DeathVenom54/github-deploy-inator/logger"
	"github.com/DeathVenom54/github-deploy-inator/router"
	"net/http"
)

func ExecuteConfig() error {
	config, err := ReadConfig()
	if err != nil {
		return err
	}

	// start server
	r := router.CreateRouter(config)

	logger.Log(fmt.Sprintf("Listening for Github webhooks at %s", config.Port))
	err = http.ListenAndServe(config.Port, r)
	if err != nil {
		return err
	}

	return nil
}
