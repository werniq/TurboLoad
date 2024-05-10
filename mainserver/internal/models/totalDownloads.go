package models

func (m *Database) UpdateTotalDownloads() error {
	stmt := `UPDATE total_downloads SET total_downloads = total_downloads + 1`

	_, err := m.DB.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}

// GetTotalDownloads
func (m *Database) GetTotalDownloads() (int, error) {
	stmt := `SELECT total_downloads FROM total_downloads`

	var id int

	row := m.DB.QueryRow(stmt)

	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
