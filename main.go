package main

import (
	"github.com/DeathVenom54/github-deploy-inator/config"
	"github.com/DeathVenom54/github-deploy-inator/logger"
)

func main() {
	err := config.ExecuteConfig()
	if err != nil {
		logger.Error(err, true)
	}
}
