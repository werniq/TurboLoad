package server

import "github.com/werniq/turboload/logger"

var (
    currentConcurrentRequest = 0
)

func Run() {
	if err := run(); err != nil {
		logger.ErrorLogger.Fatalln(err)
	}
}
