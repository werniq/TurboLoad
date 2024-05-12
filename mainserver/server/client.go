package server

import (
	"github.com/joho/godotenv"
	"github.com/werniq/turboload/internal/models"
	"github.com/werniq/turboload/logger"
	"os"
	"path/filepath"
)

var (
	currentConcurrentRequest = 0
	database                 *models.Database
)

func Run() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	p := filepath.Dir(dir)

	if err := godotenv.Load(p + "\\.env"); err != nil {
		panic(err)
	}

	database = models.NewDatabase()

	if err := run(); err != nil {
		logger.ErrorLogger.Fatalln(err)
	}
}
