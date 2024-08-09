package cmd

import (
	"database/sql"
	"errors"
)

func CreateTable(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS links (
        id TEXT PRIMARY KEY,
        is_deleted BOOLEAN,
        is_visibled BOOLEAN,
        original_url TEXT,
        short_url TEXT,
        created_at TIMESTAMP,
        modified_at TIMESTAMP,
        expiration TIMESTAMP,
        clicks INTEGER
    );
    `
	_, err := db.Exec(query)
	if err != nil {
		return errors.New("failed to create links table: " + err.Error())
	}

	return nil
}
