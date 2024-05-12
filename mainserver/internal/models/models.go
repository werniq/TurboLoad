package models

import (
	"database/sql"
	"github.com/werniq/turboload/internal/driver"
	"time"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase() *Database {
	db, err := driver.OpenDb()
	if err != nil {
		panic(err)
	}

	return &Database{
		DB: db,
	}
}

/*
	CREATE TABLE files_statistics(
		id bigserial not null primary key,
		filename varchar(255) not null,
		downloads_count bigint,
		size bigint,
		created_at timestamp default current_timestamp
	);

	CREATE TABLE recent_statistics(
	  	id bigserial not null primary key,
		filename varchar(255) not null,
		download_time int,
		created_at timestamp default current_timestamp
	);
*/

// AfterResponseUpdate updates statistics in database after download
func (m *Database) AfterResponseUpdate(filename string, duration int64) error {
	stmt := `UPDATE file_info SET 
                            downloads_count = downloads_count + 1 
                        WHERE filename = $1`

	_, err := m.DB.Exec(stmt, filename)
	if err != nil {
		return err
	}

	stmt = `INSERT INTO statistics(
            	filename, 
                download_duration, 
                created_at) 
		VALUES($1, $2, $3)`

	_, err = m.DB.Exec(stmt, filename, duration, time.Now())
	if err != nil {
		return err
	}

	return nil
}
