package server

import (
	"github.com/werniq/turboload/internal/models"
	"github.com/werniq/turboload/logger"
)

var (
	currentConcurrentRequest = 0
	database                 *models.Database
)

func Run() {
	if err := run(); err != nil {
		logger.ErrorLogger.Fatalln(err)
	}
}
