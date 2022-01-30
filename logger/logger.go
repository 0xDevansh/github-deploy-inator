package logger

import (
	"log"
	"os"
)

func writeToFile(file, text string) error {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("Error while closing file %s: %s\n", file, err)
		}
	}(f)

	log.SetOutput(f)
	log.Println(text)

	return nil
}

func Log(message string) {
	err := writeToFile("all.log", message)
	if err != nil {
		log.Fatalf("Error while writing log to normal.log: %s", err)
	}
}

func Error(error string) {
	err := writeToFile("error.log", error)
	if err != nil {
		log.Fatalf("Error while writing log to error.log: %s", err)
	}
}
