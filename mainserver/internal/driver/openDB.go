package driver

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

// OpenDb
// by calling this function we initialize connection with database
func OpenDb() (*sql.DB, error) {
	dsn := os.Getenv("DATABASE_DSN")

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
