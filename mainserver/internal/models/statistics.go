package models

// Statistics
type Statistics struct {
	// sort files by downloads_count
	FilesData []FilesData `json:"files_data"`

	TotalDownloads      int `json:"total_downloads"`
	AverageDownloadTime int `json:"average_download_time"`
}

type FilesData struct {
	AvgDownloadTime int `json:"download_time"`
	DownloadsCount  int `json:"downloads_count"`
	Size            int `json:"size"`
	CreatedAt       int `json:"created_at"`
}

// GetFileInfo returns file's statistics (*FilesData)
func (m *Database) GetFileInfo() ([]*FilesData, error) {
	stmt := `SELECT * FROM file_info ORDER BY total_downloads DESC;`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	var filesData []*FilesData

	for rows.Next() {
		f := &FilesData{}

		if err = rows.Scan(&f.AvgDownloadTime,
			f.DownloadsCount,
			f.Size,
			f.CreatedAt,
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
