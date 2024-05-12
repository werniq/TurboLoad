package models

import "time"

// Statistics
type Statistics struct {
	// sort files by downloads_count
	FilesData []FilesData `json:"files_data"`

	TotalDownloads      int `json:"total_downloads"`
	AverageDownloadTime int `json:"average_download_time"`
}

type FilesData struct {
	Id             int       `json:"id"`
	Filename       string    `json:"filename"`
	DownloadsCount int       `json:"downloads_count"`
	Size           int       `json:"size"`
	CreatedAt      time.Time `json:"created_at"`
}

// GetAllFilesInfo returns file's statistics (*FilesData)
func (m *Database) GetAllFilesInfo() ([]*FilesData, error) {
	stmt := `SELECT id, filename, downloads_count, size, created_at FROM file_info ORDER BY downloads_count DESC;`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	var filesData []*FilesData

	for rows.Next() {
		f := &FilesData{}

		if err = rows.Scan(
			&f.Id,
			&f.Filename,
			&f.DownloadsCount,
			&f.Size,
			&f.CreatedAt,
		); err != nil {
			return nil, err
		}

		filesData = append(filesData, f)
	}

	return filesData, nil
}

// GetAllFilenames returns all file names
func (m *Database) GetAllFilenames() ([]string, error) {
	stmt := `SELECT filename FROM file_info`

	res := []string{}

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var s string

		err = rows.Scan(&s)
		if err != nil {
			return nil, err
		}

		res = append(res, s)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return res, nil
}

// InsertFile used to insert new file into database after it was created
func (m *Database) InsertFile(filename string) error {
	stmt := `INSERT INTO file_info(filename) VALUES($1)`

	_, err := m.DB.Exec(stmt, filename)
	if err != nil {
		return err
	}

	return nil
}
