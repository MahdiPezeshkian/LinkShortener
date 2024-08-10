package domain

import (
	"time"

	"github.com/MahdiPezeshkian/LinkShortener/pkg"
	"github.com/google/uuid"
)

type Link struct {
	Id          string    `json:"id"`
	Isdeleted   bool      `json:"is_deleted"`
	IsVisibled  bool      `json:"is_visibled"`
	OriginalURL string    `json:"original_url"`
	ShortURL    string    `json:"short_url"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
	Expiration  time.Time `json:"expiration"`
	Clicks      int       `json:"clicks"`
}

func NewLink(originalURL string, expiration time.Time) *Link {
	return &Link{
		Id:          uuid.NewString(),
		Isdeleted:   false,
		IsVisibled:  true,
		OriginalURL: originalURL,
		ShortURL:    pkg.RandomString(4, 8),
		CreatedAt:   time.Now(),
		Expiration:  expiration,
		Clicks:      0,
	}
}

func (l *Link) Click() {
	l.Clicks++ 
}
