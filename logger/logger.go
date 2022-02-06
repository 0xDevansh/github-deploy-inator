package logger

import (
	"io"
	"log"
	"os"
)

var (
	All *log.Logger
	Err *log.Logger
)

func Setuplogger(allFile, errFile *os.File) {
	All = log.New(io.MultiWriter(allFile, os.Stdout), "      ", log.LstdFlags) // blank space to pad error logs
	Err = log.New(io.MultiWriter(errFile, os.Stderr), "ERROR ", log.LstdFlags)
}
