package repositories

import (
	"database/sql"

	"github.com/MahdiPezeshkian/LinkShortener/internal/domain"
	"github.com/MahdiPezeshkian/LinkShortener/internal/domain/entity"
	_ "github.com/mattn/go-sqlite3"
)

type sqliteLinkRepository struct {
	db *sql.DB
}

func NewSQLiteLinkRepository(db *sql.DB) domain.LinkRepository {
	return &sqliteLinkRepository{db: db}
}

func (r *sqliteLinkRepository) Insert(link *entity.Link) error {
	_, err := r.db.Exec(`
        INSERT INTO links (id, is_deleted, is_visibled, original_url, short_url, created_at, modified_at, expiration, clicks)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		link.Id, link.Isdeleted, link.IsVisibled, link.OriginalURL, link.ShortURL, link.CreatedAt, link.ModifiedAt, link.Expiration, link.Clicks,
	)
	return err
}

func (r *sqliteLinkRepository) Update(link *entity.Link) error {
	_, err := r.db.Exec(`
        UPDATE links SET 
            is_deleted = ?, 
            is_visibled = ?, 
            original_url = ?, 
            short_url = ?, 
            modified_at = ?, 
            expiration = ?, 
            clicks = ? 
        WHERE id = ?`,
		link.Isdeleted, link.IsVisibled, link.OriginalURL, link.ShortURL, link.ModifiedAt, link.Expiration, link.Clicks, link.Id,
	)
	return err
}

// FindByID finds a link by its ID
func (r *sqliteLinkRepository) FindByID(id string) (*entity.Link, error) {
	row := r.db.QueryRow(`
        SELECT id, is_deleted, is_visibled, original_url, short_url, created_at, modified_at, expiration, clicks
        FROM links WHERE id = ?`, id)

	link := &entity.Link{}
	err := row.Scan(
		&link.Id, &link.Isdeleted, &link.IsVisibled, &link.OriginalURL, &link.ShortURL,
		&link.CreatedAt, &link.ModifiedAt, &link.Expiration, &link.Clicks,
	)
	if err != nil {
		return nil, err
	}
	return link, nil
}

// FindAll retrieves all links from the database
func (r *sqliteLinkRepository) FindAll() ([]*entity.Link, error) {
	rows, err := r.db.Query(`
        SELECT id, is_deleted, is_visibled, original_url, short_url, created_at, modified_at, expiration, clicks
        FROM links`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []*entity.Link
	for rows.Next() {
		link := &entity.Link{}
		err := rows.Scan(
			&link.Id, &link.Isdeleted, &link.IsVisibled, &link.OriginalURL, &link.ShortURL,
			&link.CreatedAt, &link.ModifiedAt, &link.Expiration, &link.Clicks,
		)
		if err != nil {
			return nil, err
		}
		links = append(links, link)
	}
	return links, nil
}

// Delete soft-deletes a link by marking it as deleted
func (r *sqliteLinkRepository) Delete(id string) error {
	_, err := r.db.Exec("UPDATE links SET is_deleted = ? WHERE id = ?", true, id)
	return err
}
