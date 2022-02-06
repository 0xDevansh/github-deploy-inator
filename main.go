package main

import (
	"github.com/DeathVenom54/github-deploy-inator/config"
	"github.com/DeathVenom54/github-deploy-inator/logger"
	"os"
)

func main() {
	// setting up logger
	allFile, err := os.OpenFile("all.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	errFile, err := os.OpenFile("error.log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	// close file
	defer func(allFile, errFile *os.File) {
		err := allFile.Close()
		if err != nil {
			panic(err)
		}
		err = errFile.Close()
		if err != nil {
			panic(err)
		}
	}(allFile, errFile)

	logger.Setuplogger(allFile, errFile)

	// finally, run app
	err = config.ExecuteConfig()
	if err != nil {
		logger.Err.Fatalln(err)
	}
}
