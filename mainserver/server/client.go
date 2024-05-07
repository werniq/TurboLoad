package server

import "github.com/werniq/turboload/logger"

const (
	address = "localhost:50051"
)

func Run() {

	if err := run(); err != nil {
		logger.ErrorLogger.Fatalln(err)
	}
}
