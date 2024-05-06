package client

import (
	"100gombs/logger"
)

const (
	address = "localhost:50051"
)

func Run() {
	if err := run(); err != nil {
		logger.ErrorLogger.Fatalln(err)
	}
}
