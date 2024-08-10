package cmd

import (
	"database/sql"
	"errors"
	"time"

	domain "github.com/MahdiPezeshkian/LinkShortener/internal/domain/Link"
	"github.com/google/uuid"
)

func SeedData(db *sql.DB) error {
	insertQuery := `
		INSERT INTO links (id, is_deleted, is_visibled, original_url, short_url, created_at, modified_at, expiration, clicks)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?);`

	checkQuery := `
		SELECT COUNT(*) FROM links WHERE original_url = ? OR short_url = ?;`

	links := []domain.Link{
		{
			Id:          uuid.NewString(),
			Isdeleted:   false,
			IsVisibled:  true,
			OriginalURL: "https://www.google.com",
			ShortURL:    "abc123",
			CreatedAt:   time.Now(),
			ModifiedAt:  time.Now(),
			Expiration:  time.Now().AddDate(0, 1, 0),
			Clicks:      10,
		},
		{
			Id:          uuid.NewString(),
			Isdeleted:   false,
			IsVisibled:  true,
			OriginalURL: "https://www.yahoo.com",
			ShortURL:    "xyz789",
			CreatedAt:   time.Now(),
			ModifiedAt:  time.Now(),
			Expiration:  time.Now().AddDate(0, 1, 0),
			Clicks:      5,
		},
	}

	for _, link := range links {
		var count int
		err := db.QueryRow(checkQuery, link.OriginalURL, link.ShortURL).Scan(&count)
		if err != nil {
			return errors.New("failed to check existing data: " + err.Error())
		}

		if count == 0 {
			_, err := db.Exec(insertQuery, link.Id, link.Isdeleted, link.IsVisibled, link.OriginalURL, link.ShortURL, link.CreatedAt, link.ModifiedAt, link.Expiration, link.Clicks)
			if err != nil {
				return errors.New("failed to seed data: " + err.Error())
			}
		}
	}

	return nil
}
