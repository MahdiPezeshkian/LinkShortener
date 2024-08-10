package repositories

import (
	"database/sql"
	"fmt"

	domain "github.com/MahdiPezeshkian/LinkShortener/internal/domain/Link"
	"github.com/MahdiPezeshkian/LinkShortener/pkg"
	_ "github.com/mattn/go-sqlite3"
)

type sqliteLinkRepository struct {
	db *sql.DB
}

func SQLiteLinkRepository(db *sql.DB) domain.LinkRepository {
	return &sqliteLinkRepository{db: db}
}

func (r *sqliteLinkRepository) Insert(link *domain.Link) error {
	_, err := r.db.Exec(`
        INSERT INTO links (id, is_deleted, is_visibled, original_url, short_url, created_at, modified_at, expiration, clicks)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		link.Id, link.Isdeleted, link.IsVisibled, link.OriginalURL, link.ShortURL, link.CreatedAt, link.ModifiedAt, link.Expiration, link.Clicks,
	)
	return err
}

func (r *sqliteLinkRepository) Update(link *domain.Link) error {
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

func (r *sqliteLinkRepository) FindByID(id string) (*domain.Link, error) {
	row := r.db.QueryRow(`
        SELECT id, is_deleted, is_visibled, original_url, short_url, created_at, modified_at, expiration, clicks
        FROM links WHERE id = ?`, id)

	link := &domain.Link{}
	err := row.Scan(
		&link.Id, &link.Isdeleted, &link.IsVisibled, &link.OriginalURL, &link.ShortURL,
		&link.CreatedAt, &link.ModifiedAt, &link.Expiration, &link.Clicks,
	)
	if err != nil {
		return nil, err
	}
	return link, nil
}

func (r *sqliteLinkRepository) FindOneByCondition(condition string, args ...interface{}) (*domain.Link, error) {
	query := `
        SELECT id, is_deleted, is_visibled, original_url, short_url, created_at, modified_at, expiration, clicks
        FROM links WHERE ` + condition + ` LIMIT 1`

	row := r.db.QueryRow(query, args...)

	link := &domain.Link{}
	err := row.Scan(
		&link.Id, &link.Isdeleted, &link.IsVisibled, &link.OriginalURL, &link.ShortURL,
		&link.CreatedAt, &link.ModifiedAt, &link.Expiration, &link.Clicks,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return link, nil
}

func (r *sqliteLinkRepository) FindManyByCondition(condition string, args ...interface{}) ([]*domain.Link, error) {
	query := `
        SELECT id, is_deleted, is_visibled, original_url, short_url, created_at, modified_at, expiration, clicks
        FROM links WHERE ` + condition

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []*domain.Link
	for rows.Next() {
		link := &domain.Link{}
		err := rows.Scan(
			&link.Id, &link.Isdeleted, &link.IsVisibled, &link.OriginalURL, &link.ShortURL,
			&link.CreatedAt, &link.ModifiedAt, &link.Expiration, &link.Clicks,
		)
		if err != nil {
			return nil, err
		}
		links = append(links, link)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return links, nil
}


func (r *sqliteLinkRepository) FindAll() ([]*domain.Link, error) {
	rows, err := r.db.Query(`
        SELECT id, is_deleted, is_visibled, original_url, short_url, created_at, modified_at, expiration, clicks
        FROM links`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var links []*domain.Link
	for rows.Next() {
		link := &domain.Link{}
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

func (r *sqliteLinkRepository) GetPaged(pagination *pkg.PaginationRequest) ([]*domain.Link, int, error) {
	order := "ASC"
	if pagination.SortOrder == "desc" {
		order = "DESC"
	}

	offset := (pagination.PageNumber - 1) * pagination.PageSize

	query := fmt.Sprintf(`
        SELECT id, is_deleted, is_visibled, original_url, short_url, created_at, modified_at, expiration, clicks
        FROM links
        WHERE is_deleted = ?
        ORDER BY created_at %s
        LIMIT ? OFFSET ?`, order)

	rows, err := r.db.Query(query, false, pagination.PageSize, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var links []*domain.Link
	for rows.Next() {
		link := &domain.Link{}
		err := rows.Scan(
			&link.Id, &link.Isdeleted, &link.IsVisibled, &link.OriginalURL, &link.ShortURL,
			&link.CreatedAt, &link.ModifiedAt, &link.Expiration, &link.Clicks,
		)
		if err != nil {
			return nil, 0, err
		}
		links = append(links, link)
	}

	var totalCount int
	countQuery := "SELECT COUNT(*) FROM links WHERE is_deleted = ?"
	err = r.db.QueryRow(countQuery, false).Scan(&totalCount)
	if err != nil {
		return nil, 0, err
	}

	return links, totalCount, nil
}

func (r *sqliteLinkRepository) Delete(id string) error {
	_, err := r.db.Exec("UPDATE links SET is_deleted = ? WHERE id = ?", true, id)
	return err
}

func (r *sqliteLinkRepository) HardDelete(id string) error {
	_, err := r.db.Exec("DELETE FROM links WHERE id = ?", id)
	return err
}
